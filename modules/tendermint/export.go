package tendermint

import (
	"github.com/irisnet/service-sdk-go/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type Tm interface {
	Block(height int64) (QueryBlockResponse, error)
	Commit(height int64) (QueryCommitResponse, error)

	Tx(hash string) (QueryTxResponse, error)
	Txs(builder *types.EventQueryBuilder, page, size int) (QueryTxsResponse, error)

	Status() (QueryStatusResponse, error)
	NetInfo() (QueryNetInfoResponse, error)
}

type QueryBlockResponse struct {
	Block       types.Block
	BlockResult types.BlockResult
}

type QueryTxResponse struct {
	types.ResultQueryTx
}

type QueryTxsResponse struct {
	types.ResultSearchTxs
}

type QueryCommitResponse struct {
	Commit *ctypes.ResultCommit
}

type QueryStatusResponse struct {
	Status *ctypes.ResultStatus
}

type QueryNetInfoResponse struct {
	Info *ctypes.ResultNetInfo
}
