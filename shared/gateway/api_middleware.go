package gateway

import (
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

// APIProxyMiddleware is a proxy between an Ethereum consensus API HTTP client and grpc-gateway.
// The purpose of the proxy is to handle HTTP requests and gRPC responses in such a way that:
//   - Ethereum consensus API requests can be handled by grpc-gateway correctly
//   - gRPC responses can be returned as spec-compliant Ethereum consensus API responses
type APIProxyMiddleware struct {
	GatewayAddress  string
	ProxyAddress    string
	EndpointCreator EndpointFactory
	router          *mux.Router
}

// EndpointFactory is responsible for creating new instances of Endpoint values.
type EndpointFactory interface {
	Create(path string) (*Endpoint, error)
	Paths() []string
	IsNil() bool
}

// Endpoint is a representation of an API HTTP endpoint that should be proxied by the middleware.
type Endpoint struct {
	Path               string          // The path of the HTTP endpoint.
	PostRequest        interface{}     // The struct corresponding to the JSON structure used in a POST request.
	PostResponse       interface{}     // The struct corresponding to the JSON structure used in a POST response.
	RequestURLLiterals []string        // Names of URL parameters that should not be base64-encoded.
	RequestQueryParams []QueryParam    // Query parameters of the request.
	GetResponse        interface{}     // The struct corresponding to the JSON structure used in a GET response.
	Err                ErrorJSON       // The struct corresponding to the error that should be returned in case of a request failure.
	Hooks              HookCollection  // A collection of functions that can be invoked at various stages of the request/response cycle.
	CustomHandlers     []CustomHandler // Functions that will be executed instead of the default request/response behavior.
}

// DefaultEndpoint returns an Endpoint with default configuration, e.g. DefaultErrorJSON for error handling.
func DefaultEndpoint() Endpoint {
	return Endpoint{
		Err: &DefaultErrorJSON{},
	}
}

// QueryParam represents a single query parameter's metadata.
type QueryParam struct {
	Name string
	Hex  bool
	Enum bool
}

// Hook is a function that can be invoked at various stages of the request/response cycle, leading to custom behavior for a specific endpoint.
type Hook = func(endpoint Endpoint, w http.ResponseWriter, req *http.Request) ErrorJSON

// CustomHandler is a function that can be invoked at the very beginning of the request,
// essentially replacing the whole default request/response logic with custom logic for a specific endpoint.
type CustomHandler = func(m *APIProxyMiddleware, endpoint Endpoint, w http.ResponseWriter, req *http.Request) (handled bool)

// HookCollection contains hooks that can be used to amend the default request/response cycle with custom logic for a specific endpoint.
type HookCollection struct {
	OnPreDeserializeRequestBodyIntoContainer  []Hook
	OnPostDeserializeRequestBodyIntoContainer []Hook
}

// fieldProcessor applies the processing function f to a value when the tag is present on the field.
type fieldProcessor struct {
	tag string
	f   func(value reflect.Value) error
}

// Run starts the proxy, registering all proxy endpoints on APIProxyMiddleware.ProxyAddress.
func (m *APIProxyMiddleware) Run() error {
	m.router = mux.NewRouter()

	for _, path := range m.EndpointCreator.Paths() {
		m.handleAPIPath(path, m.EndpointCreator)
	}

	return http.ListenAndServe(m.ProxyAddress, m.router)
}

func (m *APIProxyMiddleware) handleAPIPath(path string, endpointFactory EndpointFactory) {
	m.router.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		endpoint, err := endpointFactory.Create(path)
		if err != nil {
			errJSON := InternalServerErrorWithMessage(err, "could not create endpoint")
			WriteError(w, errJSON, nil)
		}

		for _, handler := range endpoint.CustomHandlers {
			if handler(m, *endpoint, w, req) {
				return
			}
		}

		if req.Method == "POST" {
			for _, hook := range endpoint.Hooks.OnPreDeserializeRequestBodyIntoContainer {
				if errJSON := hook(*endpoint, w, req); errJSON != nil {
					WriteError(w, errJSON, nil)
					return
				}
			}

			if errJSON := DeserializeRequestBodyIntoContainer(req.Body, endpoint.PostRequest); errJSON != nil {
				WriteError(w, errJSON, nil)
				return
			}
			for _, hook := range endpoint.Hooks.OnPostDeserializeRequestBodyIntoContainer {
				if errJSON := hook(*endpoint, w, req); errJSON != nil {
					WriteError(w, errJSON, nil)
					return
				}
			}

			if errJSON := ProcessRequestContainerFields(endpoint.PostRequest); errJSON != nil {
				WriteError(w, errJSON, nil)
				return
			}
			if errJSON := SetRequestBodyToRequestContainer(endpoint.PostRequest, req); errJSON != nil {
				WriteError(w, errJSON, nil)
				return
			}
		}

		if errJSON := m.PrepareRequestForProxying(*endpoint, req); errJSON != nil {
			WriteError(w, errJSON, nil)
			return
		}
		grpcResponse, errJSON := ProxyRequest(req)
		if errJSON != nil {
			WriteError(w, errJSON, nil)
			return
		}
		grpcResponseBody, errJSON := ReadGrpcResponseBody(grpcResponse.Body)
		if errJSON != nil {
			WriteError(w, errJSON, nil)
			return
		}

		var responseJSON []byte
		if !GrpcResponseIsEmpty(grpcResponseBody) {
			if errJSON := DeserializeGrpcResponseBodyIntoErrorJSON(endpoint.Err, grpcResponseBody); errJSON != nil {
				WriteError(w, errJSON, nil)
				return
			}
			if endpoint.Err.Msg() != "" {
				HandleGrpcResponseError(endpoint.Err, grpcResponse, w)
				return
			}
			var response interface{}
			if req.Method == "GET" {
				response = endpoint.GetResponse
			} else {
				response = endpoint.PostResponse
			}
			if errJSON := DeserializeGrpcResponseBodyIntoContainer(grpcResponseBody, response); errJSON != nil {
				WriteError(w, errJSON, nil)
				return
			}
			if errJSON := ProcessMiddlewareResponseFields(response); errJSON != nil {
				WriteError(w, errJSON, nil)
				return
			}
			var errJSON ErrorJSON
			responseJSON, errJSON = SerializeMiddlewareResponseIntoJSON(response)
			if errJSON != nil {
				WriteError(w, errJSON, nil)
				return
			}
		}

		if errJSON := WriteMiddlewareResponseHeadersAndBody(req, grpcResponse, responseJSON, w); errJSON != nil {
			WriteError(w, errJSON, nil)
			return
		}
		if errJSON := Cleanup(grpcResponse.Body); errJSON != nil {
			WriteError(w, errJSON, nil)
			return
		}
	})
}
