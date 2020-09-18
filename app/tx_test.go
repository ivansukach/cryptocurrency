package app

import (
	"encoding/base64"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTendermintTxDecode(t *testing.T) {
	cdc := MakeCodec() // This needs to have every single module codec registered!!!
	txStr := "0A55A1D5AA9C0A14764ADB7322C202717A00695D77454C1C7F4CDF56122D636F736D6F73317038656539703564616C796E356A3573647370347733356739327570307A33716474653037351A0A0A046F63746112023330120410C09A0C1A6A0A26EB5AE98721022FB458D3DAA68EDC3F4CBB3ED03AD27722215A556118684B2F46ADCE40D55B881240B6294F5D82202BC1B9B748D0D426762F76A2910303E69E213D60771E6F5A3EFF3F15F91A1FBBC228F83F2141925607AA2C26FE5FDF4136E1050D0110AB37D511"

	txBz, err := base64.StdEncoding.DecodeString(txStr)
	require.NoError(t, err)

	var tx auth.StdTx
	require.NoError(t, cdc.UnmarshalBinaryLengthPrefixed(txBz, &tx))
	require.NotNil(t, tx)

	fmt.Println(tx)
}
