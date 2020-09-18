package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type TransferOfFunds struct {
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Amount   sdk.Coins      `json:"amount" yaml:"amount"`
	Time     string         `json:"time" yaml:"time"`
}

func (t TransferOfFunds) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Sender: %s
	Receiver: %s
	Amount: %s
	Time: %s`,
		t.Sender,
		t.Receiver,
		t.Amount,
		t.Time,
	))
}
