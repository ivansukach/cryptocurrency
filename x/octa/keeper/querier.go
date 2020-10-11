package keeper

import (
	"github.com/ivansukach/cryptocurrency/x/octa/types"
	"github.com/ivansukach/modified-cosmos-sdk/codec"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	sdkerrors "github.com/ivansukach/modified-cosmos-sdk/types/errors"
	"github.com/sirupsen/logrus"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a new querier for octa clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListTransfers:
			return listTransfers(ctx, k)
		case types.QueryGetTransferOfFunds:
			return getTransfer(ctx, path[1:], k)
			// TODO: Put the modules query routes
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown octa query endpoint")
		}
	}
}

// RemovePrefixFromHash removes the prefix from the key
func RemovePrefixFromHash(key []byte, prefix []byte) (hash []byte) {
	hash = key[len(prefix):]
	return hash
}

func listTransfers(ctx sdk.Context, k Keeper) ([]byte, error) {
	var transferList types.QueryResTransfers

	iterator := k.GetTransfersIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		logrus.Info("TRANSFER ID WITH PREFIX: ", key)
		id := RemovePrefixFromHash(iterator.Key(), []byte(types.TransferPrefix))
		logrus.Info("TRANSFER ID: ", string(id))
		var transfer types.TransferOfFunds
		transfer, err := k.GetTransfer(ctx, key)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		//k.cdc.MustUnmarshalBinaryLengthPrefixed(transferHash, &transfer)
		logrus.Info("TRANSFER: ", transfer)
		transferList = append(transferList, key)
	}
	logrus.Info("TRANSFER IDs LIST: ", transferList)
	res, err := codec.MarshalJSONIndent(k.cdc, transferList)
	if err != nil {
		return res, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getTransfer(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	hash := path[0]
	transfer, err := k.GetTransfer(ctx, hash)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, transfer)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
