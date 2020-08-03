package tendermint

import (
	"github.com/irisnet/service-sdk-go/codec"
	sdk "github.com/irisnet/service-sdk-go/types"
)

var (
	_ Tm = tmClient{}
)

type tmClient struct {
	cdc *codec.Codec
	sdk.BaseClient
}

func NewClient(cdc *codec.Codec, baseClient sdk.BaseClient) Tm {
	return tmClient{
		cdc:        cdc,
		BaseClient: baseClient,
	}
}

func (tm tmClient) Block(height int64) (QueryBlockResponse, error) {
	block, err := tm.BaseClient.Block(&height)
	if err != nil {
		return QueryBlockResponse{}, err
	}

	blockResult, err := tm.BaseClient.BlockResults(&height)
	if err != nil {
		return QueryBlockResponse{}, err
	}

	return QueryBlockResponse{
		Block:       sdk.ParseBlock(tm.cdc, block.Block),
		BlockResult: sdk.ParseBlockResult(blockResult),
	}, nil
}

func (tm tmClient) Commit(height int64) (QueryCommitResponse, error) {
	commit, err := tm.BaseClient.Commit(&height)
	if err != nil {
		return QueryCommitResponse{}, err
	}
	return QueryCommitResponse{
		Commit: commit,
	}, nil
}

func (tm tmClient) Tx(hash string) (QueryTxResponse, error) {
	tx, err := tm.QueryTx(hash)
	if err != nil {
		return QueryTxResponse{}, err
	}
	return QueryTxResponse{
		ResultQueryTx: tx,
	}, nil
}

func (tm tmClient) Txs(builder *sdk.EventQueryBuilder, page, size int) (QueryTxsResponse, error) {
	txs, err := tm.QueryTxs(builder, page, size)
	if err != nil {
		return QueryTxsResponse{}, err
	}
	return QueryTxsResponse{
		ResultSearchTxs: txs,
	}, nil
}

func (tm tmClient) Status() (QueryStatusResponse, error) {
	status, err := tm.BaseClient.Status()
	if err != nil {
		return QueryStatusResponse{}, err
	}
	return QueryStatusResponse{
		Status: status,
	}, nil
}

func (tm tmClient) NetInfo() (QueryNetInfoResponse, error) {
	info, err := tm.BaseClient.NetInfo()
	if err != nil {
		return QueryNetInfoResponse{}, err
	}
	return QueryNetInfoResponse{
		Info: info,
	}, nil
}
