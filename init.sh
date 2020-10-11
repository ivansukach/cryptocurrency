#!/bin/bash
rm -r ~/.cryptocurrencyCLI
rm -r ~/.cryptocurrencyD

cryptocurrencyD init nodeIvanAmazon --chain-id octa

cryptocurrencyCLI config keyring-backend test

cryptocurrencyCLI keys add admin
cryptocurrencyCLI keys add genesis
#cryptocurrencyD start
#cryptocurrencyCLI rest-server --chain-id octa --trust-node
cryptocurrencyD add-genesis-account $(cryptocurrencyCLI keys show genesis -a) 7999999octa,100000000stake
cryptocurrencyD add-genesis-account $(cryptocurrencyCLI keys show admin -a) 1octa

cryptocurrencyCLI config chain-id octa
cryptocurrencyCLI config output json
cryptocurrencyCLI config indent true
cryptocurrencyCLI config trust-node true

cryptocurrencyD gentx --name genesis --keyring-backend test
cryptocurrencyD collect-gentxs