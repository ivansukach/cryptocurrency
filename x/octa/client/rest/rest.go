package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers octa-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {

	registerQueryRoutes(cliCtx, r, storeName)
	registerTxRoutes(cliCtx, r)
}
