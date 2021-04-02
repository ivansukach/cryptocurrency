module github.com/ivansukach/cryptocurrency

go 1.15

require (
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3 // indirect
	github.com/google/gofuzz v1.1.1-0.20200604201612-c04b05f3adfa // indirect
	github.com/gorilla/mux v1.8.0
	github.com/ivansukach/modified-cosmos-sdk v0.0.0-20210402081635-b63d7a634459
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/sys v0.0.0-20210317225723-c4fcb01b228e // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
