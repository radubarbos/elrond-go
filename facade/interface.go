package facade

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/vmcommon"
	"github.com/ElrondNetwork/elrond-go/data/api"
	"github.com/ElrondNetwork/elrond-go/data/esdt"
	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/data/transaction"
	"github.com/ElrondNetwork/elrond-go/debug"
	"github.com/ElrondNetwork/elrond-go/heartbeat/data"
	"github.com/ElrondNetwork/elrond-go/node/external"
	"github.com/ElrondNetwork/elrond-go/process"
)

// NodeHandler contains all functions that a node should contain.
type NodeHandler interface {
	// StartConsensus will start the consesus service for the current node
	StartConsensus() error

	// GetBalance returns the balance for a specific address
	GetBalance(address string) (*big.Int, error)

	// GetUsername returns the username for a specific address
	GetUsername(address string) (string, error)

	// GetValueForKey returns the value of a key from a given account
	GetValueForKey(address string, key string) (string, error)

	// GetKeyValuePairs returns the key-value pairs under a given address
	GetKeyValuePairs(address string) (map[string]string, error)

	// GetAllIssuedESDTs returns all the issued esdt tokens from esdt system smart contract
	GetAllIssuedESDTs(tokenType string) ([]string, error)

	// GetESDTData returns the esdt data from a given account, given key and given nonce
	GetESDTData(address, tokenID string, nonce uint64) (*esdt.ESDigitalToken, error)

	// GetAllESDTTokens returns the value of a key from a given account
	GetAllESDTTokens(address string) (map[string]*esdt.ESDigitalToken, error)

	// CreateTransaction will return a transaction from all needed fields
	CreateTransaction(nonce uint64, value string, receiver string, receiverUsername []byte, sender string, senderUsername []byte, gasPrice uint64,
		gasLimit uint64, data []byte, signatureHex string, chainID string, version uint32, options uint32) (*transaction.Transaction, []byte, error)

	// ValidateTransaction will validate a transaction
	ValidateTransaction(tx *transaction.Transaction) error
	ValidateTransactionForSimulation(tx *transaction.Transaction, checkSignature bool) error

	// SendBulkTransactions will send a bulk of transactions on the 'send transactions pipe' channel
	SendBulkTransactions(txs []*transaction.Transaction) (uint64, error)

	// GetTransaction will return a transaction based on the hash
	GetTransaction(hash string, withResults bool) (*transaction.ApiTransactionResult, error)

	// GetAccount returns an accountResponse containing information
	//  about the account correlated with provided address
	GetAccount(address string) (state.UserAccountHandler, error)

	// GetCode returns the code for the given account
	GetCode(account state.UserAccountHandler) []byte

	// GetHeartbeats returns the heartbeat status for each public key defined in genesis.json
	GetHeartbeats() []data.PubKeyHeartbeat

	// IsInterfaceNil returns true if there is no value under the interface
	IsInterfaceNil() bool

	// ValidatorStatisticsApi return the statistics for all the validators
	ValidatorStatisticsApi() (map[string]*state.ValidatorApiResponse, error)
	DirectTrigger(epoch uint32, withEarlyEndOfEpoch bool) error
	IsSelfTrigger() bool

	EncodeAddressPubkey(pk []byte) (string, error)
	DecodeAddressPubkey(pk string) ([]byte, error)

	GetQueryHandler(name string) (debug.QueryHandler, error)
	GetPeerInfo(pid string) ([]core.QueryP2PPeerInfo, error)

	GetBlockByHash(hash string, withTxs bool) (*api.Block, error)
	GetBlockByNonce(nonce uint64, withTxs bool) (*api.Block, error)
}

// TransactionSimulatorProcessor defines the actions which a transaction simulator processor has to implement
type TransactionSimulatorProcessor interface {
	ProcessTx(tx *transaction.Transaction) (*transaction.SimulationResults, error)
	IsInterfaceNil() bool
}

// ApiResolver defines a structure capable of resolving REST API requests
type ApiResolver interface {
	ExecuteSCQuery(query *process.SCQuery) (*vmcommon.VMOutput, error)
	ComputeTransactionGasLimit(tx *transaction.Transaction) (*transaction.CostResponse, error)
	StatusMetrics() external.StatusMetricsHandler
	GetTotalStakedValue() (*api.StakeValues, error)
	GetDirectStakedList() ([]*api.DirectStakedValue, error)
	GetDelegatorsList() ([]*api.Delegator, error)
	IsInterfaceNil() bool
}

// HardforkTrigger defines the structure used to trigger hardforks
type HardforkTrigger interface {
	Trigger(epoch uint32, withEarlyEndOfEpoch bool) error
	IsSelfTrigger() bool
	IsInterfaceNil() bool
}
