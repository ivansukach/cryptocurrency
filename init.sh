#!/bin/bash
rm -r ~/.octa

octadaemon init nodeOCTAGenesis --chain-id octa

#octadaemon config keyring-backend test

octadaemon keys add admin
octadaemon keys add genesis
#cryptocurrencyD start
#cryptocurrencyCLI rest-server --chain-id octa --trust-node
#gdlv run ./cmd/cryptocurrencyCLI rest-server --chain-id octa --trust-node
octadaemon add-genesis-account $(octadaemon keys show genesis -a) 7999999000000uocta
octadaemon add-genesis-account $(octadaemon keys show admin -a) 1000000uocta

#cryptocurrencyCLI config chain-id octa
#cryptocurrencyCLI config output json
#cryptocurrencyCLI config indent true
#cryptocurrencyCLI config trust-node true

octadaemon gentx genesis 500000000uocta --keyring-backend os --chain-id octa
octadaemon collect-gentxs