package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ivansukach/cryptocurrency/x/octa/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group octa queries under a subcommand
	octaQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	octaQueryCmd.AddCommand(
		flags.GetCommands(
			// TODO: Add query Cmds
			GetCmdListTransfers(queryRoute, cdc),
			GetCmdGetTransfer(queryRoute, cdc),
		)...,
	)

	return octaQueryCmd
}

// TODO: Add Query Commands
func GetCmdListTransfers(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all transfers",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListTransfers, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get transfers\n%s\n", err.Error())
				return nil
			}

			var out types.QueryResTransfers
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
func GetCmdGetTransfer(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [hashOfTransfer]",
		Short: "Query a transfer by hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			hash := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetTransferOfFunds, hash), nil)
			if err != nil {
				fmt.Printf("could not resolve scavenge %s \n%s\n", hash, err.Error())
				return nil
			}

			var out types.TransferOfFundsWithTime
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
