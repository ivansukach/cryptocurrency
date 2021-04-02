package cmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/ivansukach/modified-cosmos-sdk/baseapp"
	"github.com/ivansukach/modified-cosmos-sdk/client"
	"github.com/ivansukach/modified-cosmos-sdk/client/debug"
	"github.com/ivansukach/modified-cosmos-sdk/client/flags"
	"github.com/ivansukach/modified-cosmos-sdk/client/keys"
	"github.com/ivansukach/modified-cosmos-sdk/client/rpc"
	"github.com/ivansukach/modified-cosmos-sdk/codec"
	"github.com/ivansukach/modified-cosmos-sdk/server"
	servertypes "github.com/ivansukach/modified-cosmos-sdk/server/types"
	"github.com/ivansukach/modified-cosmos-sdk/snapshots"
	"github.com/ivansukach/modified-cosmos-sdk/store"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	authclient "github.com/ivansukach/modified-cosmos-sdk/x/auth/client"
	authcmd "github.com/ivansukach/modified-cosmos-sdk/x/auth/client/cli"
	"github.com/ivansukach/modified-cosmos-sdk/x/auth/types"
	vestingcli "github.com/ivansukach/modified-cosmos-sdk/x/auth/vesting/client/cli"
	banktypes "github.com/ivansukach/modified-cosmos-sdk/x/bank/types"
	"github.com/ivansukach/modified-cosmos-sdk/x/crisis"
	genutilcli "github.com/ivansukach/modified-cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	octa "github.com/ivansukach/cryptocurrency/app"
	"github.com/ivansukach/cryptocurrency/app/params"
)

// NewRootCmd creates a new root command for simd. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := octa.MakeEncodingConfig()
	initClientCtx := client.Context{}.
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(octa.DefaultNodeHome)

	rootCmd := &cobra.Command{
		Use:   "octadaemon",
		Short: "OCTA APP",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			return server.InterceptConfigsPreRunHandler(cmd)
		},
	}

	initRootCmd(rootCmd, encodingConfig)

	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	authclient.Codec = encodingConfig.Marshaler

	rootCmd.AddCommand(
		genutilcli.InitCmd(octa.ModuleBasics, octa.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, octa.DefaultNodeHome),
		genutilcli.GenTxCmd(octa.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, octa.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(octa.ModuleBasics),
		AddGenesisAccountCmd(octa.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		debug.Cmd(),
	)

	server.AddCommands(rootCmd, octa.DefaultNodeHome, newApp, createSimappAndExport, addModuleInitFlags)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		keys.Commands(octa.DefaultNodeHome),
	)
}
func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	octa.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		flags.LineBreak,
		vestingcli.GetTxCmd(),
	)

	octa.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// newApp is an AppCreator
func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := sdk.NewLevelDB("metadata", snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	return octa.NewOctaApp(
		logger, db, traceStore, true, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		octa.MakeEncodingConfig(), // Ideally, we would reuse the one created by NewRootCmd.
		appOpts,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshotStore(snapshotStore),
		baseapp.SetSnapshotInterval(cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval))),
		baseapp.SetSnapshotKeepRecent(cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent))),
	)
}

func createSimappAndExport(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailAllowedAddrs []string,
	appOpts servertypes.AppOptions) (servertypes.ExportedApp, error) {

	encCfg := octa.MakeEncodingConfig() // Ideally, we would reuse the one created by NewRootCmd.
	encCfg.Marshaler = codec.NewProtoCodec(encCfg.InterfaceRegistry)
	var octaApp *octa.OctaApp
	if height != -1 {
		octaApp = octa.NewOctaApp(logger, db, traceStore, false, map[int64]bool{}, "", cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), encCfg, appOpts)

		if err := octaApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		octaApp = octa.NewOctaApp(logger, db, traceStore, true, map[int64]bool{}, "", cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), encCfg, appOpts)
	}

	return octaApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
