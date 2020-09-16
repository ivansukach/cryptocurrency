package cli

import (
	"bufio"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ivansukach/cryptocurrency/x/octa/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	octaTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	octaTxCmd.AddCommand(flags.PostCommands(
		GetCmdMakeTransferOfFunds(cdc),
	)...)

	return octaTxCmd
}

// Example:
//
// GetCmd<Action> is the CLI command for doing <Action>
func GetCmdMakeTransferOfFunds(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "makeTransfer [amount] [receiver]",
		Short: "It makes new coin transfer to the receiver",
		Args:  cobra.ExactArgs(2), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			amount, err := sdk.ParseCoins(args[0])
			if err != nil {
				log.Println("Error GetCmdMakeTransfer")
				return err

			}
			msg := types.NewMsgMakeTransferOfFunds(cliCtx.GetFromAddress(), []byte(args[1]), amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
