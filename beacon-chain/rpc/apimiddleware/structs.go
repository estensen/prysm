package apimiddleware

import "github.com/prysmaticlabs/prysm/shared/gateway"

// genesisResponseJSON is used in /beacon/genesis API endpoint.
type genesisResponseJSON struct {
	Data *genesisResponseGenesisJSON `json:"data"`
}

// genesisResponseGenesisJSON is used in /beacon/genesis API endpoint.
type genesisResponseGenesisJSON struct {
	GenesisTime           string `json:"genesis_time" time:"true"`
	GenesisValidatorsRoot string `json:"genesis_validators_root" hex:"true"`
	GenesisForkVersion    string `json:"genesis_fork_version" hex:"true"`
}

// stateRootResponseJSON is used in /beacon/states/{state_id}/root API endpoint.
type stateRootResponseJSON struct {
	Data *stateRootResponseStateRootJSON `json:"data"`
}

// stateRootResponseStateRootJSON is used in /beacon/states/{state_id}/root API endpoint.
type stateRootResponseStateRootJSON struct {
	StateRoot string `json:"root" hex:"true"`
}

// stateForkResponseJSON is used in /beacon/states/{state_id}/fork API endpoint.
type stateForkResponseJSON struct {
	Data *forkJson `json:"data"`
}

// stateFinalityCheckpointResponseJSON is used in /beacon/states/{state_id}/finality_checkpoints API endpoint.
type stateFinalityCheckpointResponseJSON struct {
	Data *stateFinalityCheckpointResponseStateFinalityCheckpointJSON `json:"data"`
}

// stateFinalityCheckpointResponseStateFinalityCheckpointJSON is used in /beacon/states/{state_id}/finality_checkpoints API endpoint.
type stateFinalityCheckpointResponseStateFinalityCheckpointJSON struct {
	PreviousJustified *checkpointJSON `json:"previous_justified"`
	CurrentJustified  *checkpointJSON `json:"current_justified"`
	Finalized         *checkpointJSON `json:"finalized"`
}

// stateValidatorResponseJSON is used in /beacon/states/{state_id}/validators API endpoint.
type stateValidatorsResponseJSON struct {
	Data []*validatorContainerJSON `json:"data"`
}

// stateValidatorResponseJSON is used in /beacon/states/{state_id}/validators/{validator_id} API endpoint.
type stateValidatorResponseJSON struct {
	Data *validatorContainerJSON `json:"data"`
}

// validatorBalancesResponseJSON is used in /beacon/states/{state_id}/validator_balances API endpoint.
type validatorBalancesResponseJSON struct {
	Data []*validatorBalanceJSON `json:"data"`
}

// stateCommitteesResponseJSON is used in /beacon/states/{state_id}/committees API endpoint.
type stateCommitteesResponseJSON struct {
	Data []*committeeJSON `json:"data"`
}

// blockHeadersResponseJSON is used in /beacon/headers API endpoint.
type blockHeadersResponseJSON struct {
	Data []*blockHeaderContainerJSON `json:"data"`
}

// blockHeaderResponseJSON is used in /beacon/headers/{block_id} API endpoint.
type blockHeaderResponseJSON struct {
	Data *blockHeaderContainerJSON `json:"data"`
}

// blockResponseJSON is used in /beacon/blocks/{block_id} API endpoint.
type blockResponseJSON struct {
	Data *beaconBlockContainerJSON `json:"data"`
}

// blockRootResponseJSON is used in /beacon/blocks/{block_id}/root API endpoint.
type blockRootResponseJSON struct {
	Data *blockRootContainerJSON `json:"data"`
}

// blockAttestationsResponseJSON is used in /beacon/blocks/{block_id}/attestations API endpoint.
type blockAttestationsResponseJSON struct {
	Data []*attestationJSON `json:"data"`
}

// attestationsPoolResponseJSON is used in /beacon/pool/attestations GET API endpoint.
type attestationsPoolResponseJSON struct {
	Data []*attestationJSON `json:"data"`
}

// submitAttestationRequestJSON is used in /beacon/pool/attestations POST API endpoint.
type submitAttestationRequestJSON struct {
	Data []*attestationJSON `json:"data"`
}

// attesterSlashingsPoolResponseJSON is used in /beacon/pool/attester_slashings API endpoint.
type attesterSlashingsPoolResponseJSON struct {
	Data []*attesterSlashingJSON `json:"data"`
}

// proposerSlashingsPoolResponseJSON is used in /beacon/pool/proposer_slashings API endpoint.
type proposerSlashingsPoolResponseJSON struct {
	Data []*proposerSlashingJSON `json:"data"`
}

// voluntaryExitsPoolResponseJSON is used in /beacon/pool/voluntary_exits API endpoint.
type voluntaryExitsPoolResponseJSON struct {
	Data []*signedVoluntaryExitJSON `json:"data"`
}

// identityResponseJSON is used in /node/identity API endpoint.
type identityResponseJSON struct {
	Data *identityJSON `json:"data"`
}

// peersResponseJSON is used in /node/peers API endpoint.
type peersResponseJSON struct {
	Data []*peerJSON `json:"data"`
}

// peerResponseJSON is used in /node/peers/{peer_id} API endpoint.
type peerResponseJSON struct {
	Data *peerJSON `json:"data"`
}

// peerCountResponseJSON is used in /node/peer_count API endpoint.
type peerCountResponseJSON struct {
	Data peerCountResponse_PeerCountJson `json:"data"`
}

// peerCountResponse_PeerCountJson is used in /node/peer_count API endpoint.
type peerCountResponse_PeerCountJson struct {
	Disconnected  string `json:"disconnected"`
	Connecting    string `json:"connecting"`
	Connected     string `json:"connected"`
	Disconnecting string `json:"disconnecting"`
}

// versionResponseJSON is used in /node/version API endpoint.
type versionResponseJSON struct {
	Data *versionJSON `json:"data"`
}

// syncingResponseJSON is used in /node/syncing API endpoint.
type syncingResponseJSON struct {
	Data *syncInfoJson `json:"data"`
}

// beaconStateResponseJSON is used in /debug/beacon/states/{state_id} API endpoint.
type beaconStateResponseJSON struct {
	Data *beaconStateJson `json:"data"`
}

// forkChoiceHeadsResponseJSON is used in /debug/beacon/heads API endpoint.
type forkChoiceHeadsResponseJSON struct {
	Data []*forkChoiceHeadJson `json:"data"`
}

// forkScheduleResponseJSON is used in /config/fork_schedule API endpoint.
type forkScheduleResponseJSON struct {
	Data []*forkJson `json:"data"`
}

// depositContractResponseJSON is used in /config/deposit_contract API endpoint.
type depositContractResponseJSON struct {
	Data *depositContractJson `json:"data"`
}

// specResponseJSON is used in /config/spec API endpoint.
type specResponseJSON struct {
	Data interface{} `json:"data"`
}

// attesterDutiesRequestJSON is used in /validator/duties/attester/{epoch} API endpoint.
type attesterDutiesRequestJSON struct {
	Index []string `json:"index"`
}

// attesterDutiesResponseJSON is used in /validator/duties/attester/{epoch} API endpoint.
type attesterDutiesResponseJSON struct {
	DependentRoot string              `json:"dependent_root" hex:"true"`
	Data          []*attesterDutyJson `json:"data"`
}

// proposerDutiesResponseJSON is used in /validator/duties/proposer/{epoch} API endpoint.
type proposerDutiesResponseJSON struct {
	DependentRoot string              `json:"dependent_root" hex:"true"`
	Data          []*proposerDutyJson `json:"data"`
}

// produceBlockResponseJSON is used in /validator/blocks/{slot} API endpoint.
type produceBlockResponseJSON struct {
	Data *beaconBlockJSON `json:"data"`
}

// produceAttestationDataResponseJSON is used in /validator/attestation_data API endpoint.
type produceAttestationDataResponseJSON struct {
	Data *attestationDataJSON `json:"data"`
}

// aggregateAttestationResponseJSON is used in /validator/aggregate_attestation API endpoint.
type aggregateAttestationResponseJSON struct {
	Data *attestationJSON `json:"data"`
}

// submitBeaconCommitteeSubscriptionsRequestJSON is used in /validator/beacon_committee_subscriptions
type submitBeaconCommitteeSubscriptionsRequestJSON struct {
	Data []*beaconCommitteeSubscribeJSON `json:"data"`
}

// beaconCommitteeSubscribeJSON is used in /validator/beacon_committee_subscriptions
type beaconCommitteeSubscribeJSON struct {
	ValidatorIndex   string `json:"validator_index"`
	CommitteeIndex   string `json:"committee_index"`
	CommitteesAtSlot string `json:"committees_at_slot"`
	Slot             string `json:"slot"`
	IsAggregator     bool   `json:"is_aggregator"`
}

// submitAggregateAndProofsRequestJSON is used in /validator/aggregate_and_proofs API endpoint.
type submitAggregateAndProofsRequestJSON struct {
	Data []*signedAggregateAttestationAndProofJson `json:"data"`
}

//----------------
// Reusable types.
//----------------

type checkpointJSON struct {
	Epoch string `json:"epoch"`
	Root  string `json:"root" hex:"true"`
}

type blockRootContainerJSON struct {
	Root string `json:"root" hex:"true"`
}

type beaconBlockContainerJSON struct {
	Message   *beaconBlockJSON `json:"message"`
	Signature string           `json:"signature" hex:"true"`
}

type beaconBlockJSON struct {
	Slot          string               `json:"slot"`
	ProposerIndex string               `json:"proposer_index"`
	ParentRoot    string               `json:"parent_root" hex:"true"`
	StateRoot     string               `json:"state_root" hex:"true"`
	Body          *beaconBlockBodyJSON `json:"body"`
}

type beaconBlockBodyJSON struct {
	RandaoReveal      string                     `json:"randao_reveal" hex:"true"`
	Eth1Data          *eth1DataJSON              `json:"eth1_data"`
	Graffiti          string                     `json:"graffiti" hex:"true"`
	ProposerSlashings []*proposerSlashingJSON    `json:"proposer_slashings"`
	AttesterSlashings []*attesterSlashingJSON    `json:"attester_slashings"`
	Attestations      []*attestationJSON         `json:"attestations"`
	Deposits          []*depositJSON             `json:"deposits"`
	VoluntaryExits    []*signedVoluntaryExitJSON `json:"voluntary_exits"`
}

type blockHeaderContainerJSON struct {
	Root      string                          `json:"root" hex:"true"`
	Canonical bool                            `json:"canonical"`
	Header    *beaconBlockHeaderContainerJSON `json:"header"`
}

type beaconBlockHeaderContainerJSON struct {
	Message   *beaconBlockHeaderJSON `json:"message"`
	Signature string                 `json:"signature" hex:"true"`
}

type signedBeaconBlockHeaderJSON struct {
	Header    *beaconBlockHeaderJSON `json:"message"`
	Signature string                 `json:"signature" hex:"true"`
}

type beaconBlockHeaderJSON struct {
	Slot          string `json:"slot"`
	ProposerIndex string `json:"proposer_index"`
	ParentRoot    string `json:"parent_root" hex:"true"`
	StateRoot     string `json:"state_root" hex:"true"`
	BodyRoot      string `json:"body_root" hex:"true"`
}

type eth1DataJSON struct {
	DepositRoot  string `json:"deposit_root" hex:"true"`
	DepositCount string `json:"deposit_count"`
	BlockHash    string `json:"block_hash" hex:"true"`
}

type proposerSlashingJSON struct {
	Header_1 *signedBeaconBlockHeaderJSON `json:"signed_header_1"`
	Header_2 *signedBeaconBlockHeaderJSON `json:"signed_header_2"`
}

type attesterSlashingJSON struct {
	Attestation_1 *indexedAttestationJSON `json:"attestation_1"`
	Attestation_2 *indexedAttestationJSON `json:"attestation_2"`
}

type indexedAttestationJSON struct {
	AttestingIndices []string             `json:"attesting_indices"`
	Data             *attestationDataJSON `json:"data"`
	Signature        string               `json:"signature" hex:"true"`
}

type attestationJSON struct {
	AggregationBits string               `json:"aggregation_bits" hex:"true"`
	Data            *attestationDataJSON `json:"data"`
	Signature       string               `json:"signature" hex:"true"`
}

type attestationDataJSON struct {
	Slot            string          `json:"slot"`
	CommitteeIndex  string          `json:"index"`
	BeaconBlockRoot string          `json:"beacon_block_root" hex:"true"`
	Source          *checkpointJSON `json:"source"`
	Target          *checkpointJSON `json:"target"`
}

type depositJSON struct {
	Proof []string          `json:"proof" hex:"true"`
	Data  *deposit_DataJson `json:"data"`
}

type deposit_DataJson struct {
	PublicKey             string `json:"pubkey" hex:"true"`
	WithdrawalCredentials string `json:"withdrawal_credentials" hex:"true"`
	Amount                string `json:"amount"`
	Signature             string `json:"signature" hex:"true"`
}

type signedVoluntaryExitJSON struct {
	Exit      *voluntaryExitJSON `json:"message"`
	Signature string             `json:"signature" hex:"true"`
}

type voluntaryExitJSON struct {
	Epoch          string `json:"epoch"`
	ValidatorIndex string `json:"validator_index"`
}

type identityJSON struct {
	PeerID             string        `json:"peer_id"`
	Enr                string        `json:"enr"`
	P2PAddresses       []string      `json:"p2p_addresses"`
	DiscoveryAddresses []string      `json:"discovery_addresses"`
	Metadata           *metadataJSON `json:"metadata"`
}

type metadataJSON struct {
	SeqNumber string `json:"seq_number"`
	Attnets   string `json:"attnets" hex:"true"`
}

type peerJSON struct {
	PeerID    string `json:"peer_id"`
	Enr       string `json:"enr"`
	Address   string `json:"last_seen_p2p_address"`
	State     string `json:"state" enum:"true"`
	Direction string `json:"direction" enum:"true"`
}

type versionJSON struct {
	Version string `json:"version"`
}

type beaconStateJson struct {
	GenesisTime                 string                    `json:"genesis_time"`
	GenesisValidatorsRoot       string                    `json:"genesis_validators_root" hex:"true"`
	Slot                        string                    `json:"slot"`
	Fork                        *forkJson                 `json:"fork"`
	LatestBlockHeader           *beaconBlockHeaderJSON    `json:"latest_block_header"`
	BlockRoots                  []string                  `json:"block_roots" hex:"true"`
	StateRoots                  []string                  `json:"state_roots" hex:"true"`
	HistoricalRoots             []string                  `json:"historical_roots" hex:"true"`
	Eth1Data                    *eth1DataJSON             `json:"eth1_data"`
	Eth1DataVotes               []*eth1DataJSON           `json:"eth1_data_votes"`
	Eth1DepositIndex            string                    `json:"eth1_deposit_index"`
	Validators                  []*validatorJson          `json:"validators"`
	Balances                    []string                  `json:"balances"`
	RandaoMixes                 []string                  `json:"randao_mixes" hex:"true"`
	Slashings                   []string                  `json:"slashings"`
	PreviousEpochAttestations   []*pendingAttestationJson `json:"previous_epoch_attestations"`
	CurrentEpochAttestations    []*pendingAttestationJson `json:"current_epoch_attestations"`
	JustificationBits           string                    `json:"justification_bits" hex:"true"`
	PreviousJustifiedCheckpoint *checkpointJSON           `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint  *checkpointJSON           `json:"current_justified_checkpoint"`
	FinalizedCheckpoint         *checkpointJSON           `json:"finalized_checkpoint"`
}

type forkJson struct {
	PreviousVersion string `json:"previous_version" hex:"true"`
	CurrentVersion  string `json:"current_version" hex:"true"`
	Epoch           string `json:"epoch"`
}

type validatorContainerJSON struct {
	Index     string         `json:"index"`
	Balance   string         `json:"balance"`
	Status    string         `json:"status" enum:"true"`
	Validator *validatorJson `json:"validator"`
}

type validatorJson struct {
	PublicKey                  string `json:"pubkey" hex:"true"`
	WithdrawalCredentials      string `json:"withdrawal_credentials" hex:"true"`
	EffectiveBalance           string `json:"effective_balance"`
	Slashed                    bool   `json:"slashed"`
	ActivationEligibilityEpoch string `json:"activation_eligibility_epoch"`
	ActivationEpoch            string `json:"activation_epoch"`
	ExitEpoch                  string `json:"exit_epoch"`
	WithdrawableEpoch          string `json:"withdrawable_epoch"`
}

type validatorBalanceJSON struct {
	Index   string `json:"index"`
	Balance string `json:"balance"`
}

type committeeJSON struct {
	Index      string   `json:"index"`
	Slot       string   `json:"slot"`
	Validators []string `json:"validators"`
}

type pendingAttestationJson struct {
	AggregationBits string               `json:"aggregation_bits" hex:"true"`
	Data            *attestationDataJSON `json:"data"`
	InclusionDelay  string               `json:"inclusion_delay"`
	ProposerIndex   string               `json:"proposer_index"`
}

type forkChoiceHeadJson struct {
	Root string `json:"root" hex:"true"`
	Slot string `json:"slot"`
}

type depositContractJson struct {
	ChainId string `json:"chain_id"`
	Address string `json:"address"`
}

type syncInfoJson struct {
	HeadSlot     string `json:"head_slot"`
	SyncDistance string `json:"sync_distance"`
	IsSyncing    bool   `json:"is_syncing"`
}

type attesterDutyJson struct {
	Pubkey                  string `json:"pubkey" hex:"true"`
	ValidatorIndex          string `json:"validator_index"`
	CommitteeIndex          string `json:"committee_index"`
	CommitteeLength         string `json:"committee_length"`
	CommitteesAtSlot        string `json:"committees_at_slot"`
	ValidatorCommitteeIndex string `json:"validator_committee_index"`
	Slot                    string `json:"slot"`
}

type proposerDutyJson struct {
	Pubkey         string `json:"pubkey" hex:"true"`
	ValidatorIndex string `json:"validator_index"`
	Slot           string `json:"slot"`
}

type signedAggregateAttestationAndProofJson struct {
	Message   *aggregateAttestationAndProofJson `json:"message"`
	Signature string                            `json:"signature" hex:"true"`
}

type aggregateAttestationAndProofJson struct {
	AggregatorIndex string           `json:"aggregator_index"`
	Aggregate       *attestationJSON `json:"aggregate"`
	SelectionProof  string           `json:"selection_proof" hex:"true"`
}

//----------------
// SSZ
// ---------------

// sszResponseJson is a common abstraction over all SSZ responses.
type sszResponseJson interface {
	SSZData() string
}

// blockSSZResponseJson is used in /beacon/blocks/{block_id} API endpoint.
type blockSSZResponseJson struct {
	Data string `json:"data"`
}

func (ssz *blockSSZResponseJson) SSZData() string {
	return ssz.Data
}

// beaconStateSSZResponseJson is used in /debug/beacon/states/{state_id} API endpoint.
type beaconStateSSZResponseJson struct {
	Data string `json:"data"`
}

func (ssz *beaconStateSSZResponseJson) SSZData() string {
	return ssz.Data
}

// TODO: Documentation
// ---------------
// Events.
// ---------------

type eventHeadJson struct {
	Slot                      string `json:"slot"`
	Block                     string `json:"block" hex:"true"`
	State                     string `json:"state" hex:"true"`
	EpochTransition           bool   `json:"epoch_transition"`
	PreviousDutyDependentRoot string `json:"previous_duty_dependent_root" hex:"true"`
	CurrentDutyDependentRoot  string `json:"current_duty_dependent_root" hex:"true"`
}

type receivedBlockDataJson struct {
	Slot  string `json:"slot"`
	Block string `json:"block" hex:"true"`
}

type aggregatedAttReceivedDataJson struct {
	Aggregate *attestationJSON `json:"aggregate"`
}

type eventFinalizedCheckpointJson struct {
	Block string `json:"block" hex:"true"`
	State string `json:"state" hex:"true"`
	Epoch string `json:"epoch"`
}

type eventChainReorgJson struct {
	Slot         string `json:"slot"`
	Depth        string `json:"depth"`
	OldHeadBlock string `json:"old_head_block" hex:"true"`
	NewHeadBlock string `json:"old_head_state" hex:"true"`
	OldHeadState string `json:"new_head_block" hex:"true"`
	NewHeadState string `json:"new_head_state" hex:"true"`
	Epoch        string `json:"epoch"`
}

// ---------------
// Error handling.
// ---------------

// submitAttestationsErrorJson is a JSON representation of the error returned when submitting attestations.
type submitAttestationsErrorJson struct {
	gateway.DefaultErrorJSON
	Failures []*singleAttestationVerificationFailureJson `json:"failures"`
}

// singleAttestationVerificationFailureJson is a JSON representation of a failure when verifying a single submitted attestation.
type singleAttestationVerificationFailureJson struct {
	Index   int    `json:"index"`
	Message string `json:"message"`
}

type eventErrorJson struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
