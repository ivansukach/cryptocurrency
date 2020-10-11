package keeper

import (
	"fmt"
	"github.com/ivansukach/cryptocurrency/x/octa/types"
	"github.com/ivansukach/modified-cosmos-sdk/codec"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	"github.com/ivansukach/modified-cosmos-sdk/x/bank"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the octa store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a octa keeper
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetTransfer(ctx sdk.Context, keyWithPrefix string) (types.TransferOfFunds, error) {
	store := ctx.KVStore(k.storeKey)
	var item types.TransferOfFunds
	//byteKey := []byte(types.TransferPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get([]byte(keyWithPrefix)), &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (k Keeper) SetTransfer(ctx sdk.Context, value types.TransferOfFunds) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
	//trxHash := k.cdc.MustMarshalBinaryLengthPrefixed(&types.TransferOfFundsWithTime{
	//	Sender:   value.Sender,
	//	Receiver: value.Receiver,
	//	Amount:   value.Amount,
	//	Time:     time.Now().UTC(),
	//})
	key := []byte(types.TransferPrefix + value.Id)
	store.Set(key, bz)
}

func (k Keeper) GetTransfersIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.TransferPrefix))
}
