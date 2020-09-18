package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/go-amino"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgMakeTransferOfFunds{}, "octa/MakeTransferOfFunds", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = amino.NewCodec()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
