#!/bin/bash
octadaemon init nodeHomeGenesis --chain-id octa

#octadaemon config output json
#octadaemon config indent true
#octadaemon config trust-node true
#octadaemon config chain-id namechain
#octadaemon config keyring-backend os
##octadaemon config keyring-backend test
#
#octadaemon keys add admin
#octadaemon keys add genesis
##octadaemon start
##cryptocurrencyCLI rest-server --chain-id octa --trust-node
##gdlv run ./cmd/cryptocurrencyCLI rest-server --chain-id octa --trust-node
##octadaemon q distribution rewards $(octadaemon keys show genesis -a) cosmosvaloper1x5ct04apqx27swmsklxmh4mth5xsvzfx79kdjh --chain-id octa
##octadaemon start --minimum-gas-prices 0.01uocta
octadaemon add-genesis-account $(octadaemon keys show genesis -a) 7999999000000uocta
octadaemon add-genesis-account $(octadaemon keys show admin -a) 1000000uocta
#
##cryptocurrencyCLI config chain-id octa
##cryptocurrencyCLI config output json
##cryptocurrencyCLI config indent true
##cryptocurrencyCLI config trust-node true
#
octadaemon gentx genesis 1000000000uocta --keyring-backend os --chain-id octa --min-self-delegation 888000000
octadaemon collect-gentxs