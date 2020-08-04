package service

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	sdk "github.com/irisnet/service-sdk-go/types"
)

// SplitRequestID splits the given contextID to contextID, batchCounter, requestHeight, batchRequestIndex
func splitRequestID(reqID string) (sdk.HexBytes, uint64, int64, int16, error) {
	requestID, err := hex.DecodeString(reqID)
	if err != nil {
		return nil, 0, 0, 0, errors.New("invalid request id")
	}

	if len(requestID) != requestIDLen {
		return nil, 0, 0, 0, errors.New("invalid request id")
	}

	reqCtxID := requestID[0:40]
	batchCounter := binary.BigEndian.Uint64(requestID[40:48])
	requestHeight := int64(binary.BigEndian.Uint64(requestID[48:56]))
	batchRequestIndex := int16(binary.BigEndian.Uint16(requestID[56:]))
	return reqCtxID, batchCounter, requestHeight, batchRequestIndex, nil
}
