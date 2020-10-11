package cli

import (
	"bufio"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/ivansukach/cryptocurrency/x/octa/types"
	"github.com/ivansukach/modified-cosmos-sdk/client"
	"github.com/ivansukach/modified-cosmos-sdk/client/context"
	"github.com/ivansukach/modified-cosmos-sdk/client/flags"
	"github.com/ivansukach/modified-cosmos-sdk/codec"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	"github.com/ivansukach/modified-cosmos-sdk/x/auth"
	"github.com/ivansukach/modified-cosmos-sdk/x/auth/client/utils"
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
			addr, err := sdk.AccAddressFromBech32(args[1])
			//log.Println("RECEIVER: ", addr)
			//log.Println("SENDER: ", cliCtx.GetFromAddress())
			if err != nil {
				log.Println("Error: ", err)
				return err

			}
			msg := types.NewMsgMakeTransferOfFunds(cliCtx.GetFromAddress(), addr, amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
