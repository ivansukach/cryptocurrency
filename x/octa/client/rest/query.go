package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/ivansukach/cryptocurrency/x/octa/types"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	// TODO: Define your GET REST endpoints
	r.HandleFunc("/octa/parameters", queryParamsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/list", queryListHandlerFn(cliCtx)).Methods("GET")
}

func queryParamsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/parameters", types.QuerierRoute)

		res, height, err := cliCtx.QueryWithData(route, nil)
		log.Println("RES:", res)
		log.Println("HEIGHT:", height)
		log.Println("ERROR:", err)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryListHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/list", types.QuerierRoute)
		res, height, err := cliCtx.QueryWithData(route, nil)
		log.Println("RES:", res)
		log.Println("HEIGHT:", height)
		log.Println("ERROR:", err)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
