package main

import (
	"github.com/ivansukach/cryptocurrency/cmd/octadaemon/cmd"
	"os"

	"github.com/ivansukach/modified-cosmos-sdk/server"
	svrcmd "github.com/ivansukach/modified-cosmos-sdk/server/cmd"

	app "github.com/ivansukach/cryptocurrency/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
