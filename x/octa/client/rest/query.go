package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/ivansukach/cryptocurrency/x/octa/types"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	// TODO: Define your GET REST endpoints
	r.HandleFunc("/octa/parameters", queryParamsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/list", queryListHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/api/txs", queryListTransactionsHandlerFn(cliCtx)).Methods("GET")
	r.Use(customCORSMiddleware())
	//r.Use(mux.CORSMethodMiddleware(r))
	//mux.MiddlewareFunc()
	//corsObj := handlers.AllowedOrigins([]string{"*"})
}
func customCORSMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
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

func queryListTransactionsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		limit := r.URL.Query().Get("limit")
		logrus.Info("LIMIT: ", limit)
		if !ok {
			return
		}

		//route := fmt.Sprintf("query/txs", types.QuerierRoute)
		res, height, err := cliCtx.Query("txs --events message.action=MakeTransferOfFunds")
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
