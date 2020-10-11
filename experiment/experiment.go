package main

import (
	"encoding/hex"
	"github.com/ivansukach/cryptocurrency/app"
	authtypes "github.com/ivansukach/modified-cosmos-sdk/x/auth/types"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("hello")
	cdc := app.MakeCodec()
	var stdTx authtypes.StdTx
	txBytes, err := hex.DecodeString("8402282816A90A8B01A1D5AA9C0A2434626336326630312D343638662D343864622D386166372D316363316263313665386132121415BBC8CB6BD62A53670DF1EA8E6CBC4D3A46EC521A1418C242E0683962C67A1789D073FFAE19D9F44533220A0A046F637461120233392A27323032302D31302D31312030333A34393A35382E363031353631303933202B3030303020555443120410C09A0C1A6A0A26EB5AE98721021424FA5D1EFC093BBB2FE4186C4C76FCAFC6A4AF78D2C77BE95E1569402DF4351240FA619A36A23E9B0FD8F07C9DC7BF9BB17176993CFBCB6361E64EFAB13E47AA9369DE8173E246DE54CA05794EC715AA1B88C411BC8187FC4229D728613E52429A")
	if err != nil {
		logrus.Error("ERROR:", err)
	}
	err = cdc.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx)
	logrus.Printf("%+v\n", stdTx)
}
