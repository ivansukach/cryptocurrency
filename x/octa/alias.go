package octa

import (
	"github.com/ivansukach/cryptocurrency/x/octa/keeper"
	"github.com/ivansukach/cryptocurrency/x/octa/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc

	NewMsgMakeTransferOfFunds = types.NewMsgMakeTransferOfFunds
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgMakeTransferOfFunds = types.MsgMakeTransferOfFunds
)
