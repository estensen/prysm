package apimiddleware

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/prysmaticlabs/prysm/beacon-chain/rpc/eth/v1/events"
	"github.com/prysmaticlabs/prysm/shared/gateway"
	"github.com/prysmaticlabs/prysm/shared/grpcutils"
	"github.com/r3labs/sse"
)

type sszConfig struct {
	sszPath      string
	fileName     string
	responseJSON sszResponseJSON
}

func handleGetBeaconStateSSZ(m *gateway.APIProxyMiddleware, endpoint gateway.Endpoint, w http.ResponseWriter, req *http.Request) (handled bool) {
	config := sszConfig{
		sszPath:      "/eth/v1/debug/beacon/states/{state_id}/ssz",
		fileName:     "beacon_state.ssz",
		responseJSON: &beaconStateSSZResponseJSON{},
	}
	return handleGetSSZ(m, endpoint, w, req, config)
}

func handleGetBeaconBlockSSZ(m *gateway.APIProxyMiddleware, endpoint gateway.Endpoint, w http.ResponseWriter, req *http.Request) (handled bool) {
	config := sszConfig{
		sszPath:      "/eth/v1/beacon/blocks/{block_id}/ssz",
		fileName:     "beacon_block.ssz",
		responseJSON: &blockSSZResponseJSON{},
	}
	return handleGetSSZ(m, endpoint, w, req, config)
}

func handleGetSSZ(
	m *gateway.APIProxyMiddleware,
	endpoint gateway.Endpoint,
	w http.ResponseWriter,
	req *http.Request,
	config sszConfig,
) (handled bool) {
	if !sszRequested(req) {
		return false
	}

	if errJSON := prepareSSZRequestForProxying(m, endpoint, req, config.sszPath); errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
		return true
	}
	grpcResponse, errJSON := gateway.ProxyRequest(req)
	if errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
		return true
	}
	grpcResponseBody, errJSON := gateway.ReadGrpcResponseBody(grpcResponse.Body)
	if errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
		return true
	}
	if errJSON := gateway.DeserializeGrpcResponseBodyIntoErrorJSON(endpoint.Err, grpcResponseBody); errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
		return true
	}
	if endpoint.Err.Msg() != "" {
		gateway.HandleGrpcResponseError(endpoint.Err, grpcResponse, w)
		return true
	}
	if errJSON := gateway.DeserializeGrpcResponseBodyIntoContainer(grpcResponseBody, config.responseJSON); errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
		return true
	}
	responseSsz, errJSON := serializeMiddlewareResponseIntoSSZ(config.responseJSON.SSZData())
	if errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
		return true
	}
	if errJSON := writeSSZResponseHeaderAndBody(grpcResponse, w, responseSsz, config.fileName); errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
		return true
	}
	if errJSON := gateway.Cleanup(grpcResponse.Body); errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
		return true
	}

	return true
}

func sszRequested(req *http.Request) bool {
	accept, ok := req.Header["Accept"]
	if !ok {
		return false
	}
	for _, v := range accept {
		if v == "application/octet-stream" {
			return true
		}
	}
	return false
}

func prepareSSZRequestForProxying(m *gateway.APIProxyMiddleware, endpoint gateway.Endpoint, req *http.Request, sszPath string) gateway.ErrorJSON {
	req.URL.Scheme = "http"
	req.URL.Host = m.GatewayAddress
	req.RequestURI = ""
	req.URL.Path = sszPath
	return gateway.HandleURLParameters(endpoint.Path, req, []string{})
}

func serializeMiddlewareResponseIntoSSZ(data string) (sszResponse []byte, errJSON gateway.ErrorJSON) {
	// Serialize the SSZ part of the deserialized value.
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, gateway.InternalServerErrorWithMessage(err, "could not decode response body into base64")
	}
	return b, nil
}

func writeSSZResponseHeaderAndBody(grpcResp *http.Response, w http.ResponseWriter, responseSsz []byte, fileName string) gateway.ErrorJSON {
	var statusCodeHeader string
	for h, vs := range grpcResp.Header {
		// We don't want to expose any gRPC metadata in the HTTP response, so we skip forwarding metadata headers.
		if strings.HasPrefix(h, "Grpc-Metadata") {
			if h == "Grpc-Metadata-"+grpcutils.HTTPCodeMetadataKey {
				statusCodeHeader = vs[0]
			}
		} else {
			for _, v := range vs {
				w.Header().Set(h, v)
			}
		}
	}
	if statusCodeHeader != "" {
		code, err := strconv.Atoi(statusCodeHeader)
		if err != nil {
			return gateway.InternalServerErrorWithMessage(err, "could not parse status code")
		}
		w.WriteHeader(code)
	} else {
		w.WriteHeader(grpcResp.StatusCode)
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(responseSsz)))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	if _, err := io.Copy(w, ioutil.NopCloser(bytes.NewReader(responseSsz))); err != nil {
		return gateway.InternalServerErrorWithMessage(err, "could not write response message")
	}
	return nil
}

func handleEvents(m *gateway.APIProxyMiddleware, _ gateway.Endpoint, w http.ResponseWriter, req *http.Request) (handled bool) {
	sseClient := sse.NewClient("http://" + m.GatewayAddress + req.URL.RequestURI())
	eventChan := make(chan *sse.Event)

	// We use grpc-gateway as the server side of events, not the sse library.
	// Because of this subscribing to streams doesn't work as intended, resulting in each event being handled by all subscriptions.
	// To handle events properly, we subscribe just once using a placeholder value ('events') and handle all topics inside this subscription.
	if err := sseClient.SubscribeChan("events", eventChan); err != nil {
		gateway.WriteError(w, gateway.InternalServerError(err), nil)
		sseClient.Unsubscribe(eventChan)
		return
	}

	errJSON := receiveEvents(eventChan, w, req)
	if errJSON != nil {
		gateway.WriteError(w, errJSON, nil)
	}

	sseClient.Unsubscribe(eventChan)
	return true
}

func receiveEvents(eventChan <-chan *sse.Event, w http.ResponseWriter, req *http.Request) gateway.ErrorJSON {
	for {
		select {
		case msg := <-eventChan:
			var data interface{}

			// The message's event comes to us with trailing whitespace.  Remove it here for
			// ease of future procesing.
			msg.Event = bytes.TrimSpace(msg.Event)

			switch string(msg.Event) {
			case events.HeadTopic:
				data = &eventHeadJSON{}
			case events.BlockTopic:
				data = &receivedBlockDataJSON{}
			case events.AttestationTopic:
				data = &attestationJSON{}

				// Data received in the event does not fit the expected event stream output.
				// We extract the underlying attestation from event data
				// and assign the attestation back to event data for further processing.
				eventData := &aggregatedAttReceivedDataJSON{}
				if err := json.Unmarshal(msg.Data, eventData); err != nil {
					return gateway.InternalServerError(err)
				}
				attData, err := json.Marshal(eventData.Aggregate)
				if err != nil {
					return gateway.InternalServerError(err)
				}
				msg.Data = attData
			case events.VoluntaryExitTopic:
				data = &signedVoluntaryExitJSON{}
			case events.FinalizedCheckpointTopic:
				data = &eventFinalizedCheckpointJSON{}
			case events.ChainReorgTopic:
				data = &eventChainReorgJSON{}
			case "error":
				data = &eventErrorJSON{}
			default:
				return &gateway.DefaultErrorJSON{
					Message: fmt.Sprintf("Event type '%s' not supported", string(msg.Event)),
					Code:    http.StatusInternalServerError,
				}
			}

			if errJSON := writeEvent(msg, w, data); errJSON != nil {
				return errJSON
			}
			if errJSON := flushEvent(w); errJSON != nil {
				return errJSON
			}
		case <-req.Context().Done():
			return nil
		}
	}
}

func writeEvent(msg *sse.Event, w http.ResponseWriter, data interface{}) gateway.ErrorJSON {
	if err := json.Unmarshal(msg.Data, data); err != nil {
		return gateway.InternalServerError(err)
	}
	if errJSON := gateway.ProcessMiddlewareResponseFields(data); errJSON != nil {
		return errJSON
	}
	dataJSON, errJSON := gateway.SerializeMiddlewareResponseIntoJSON(data)
	if errJSON != nil {
		return errJSON
	}

	w.Header().Set("Content-Type", "text/event-stream")

	if _, err := w.Write([]byte("event: ")); err != nil {
		return gateway.InternalServerError(err)
	}
	if _, err := w.Write(msg.Event); err != nil {
		return gateway.InternalServerError(err)
	}
	if _, err := w.Write([]byte("\ndata: ")); err != nil {
		return gateway.InternalServerError(err)
	}
	if _, err := w.Write(dataJSON); err != nil {
		return gateway.InternalServerError(err)
	}
	if _, err := w.Write([]byte("\n\n")); err != nil {
		return gateway.InternalServerError(err)
	}

	return nil
}

func flushEvent(w http.ResponseWriter) gateway.ErrorJSON {
	flusher, ok := w.(http.Flusher)
	if !ok {
		return &gateway.DefaultErrorJSON{Message: fmt.Sprintf("Flush not supported in %T", w), Code: http.StatusInternalServerError}
	}
	flusher.Flush()
	return nil
}
