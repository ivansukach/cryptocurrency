package types

import (
	"fmt"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	"strings"
)

type TransferOfFunds struct {
	Id       string         `json:"id" yaml:"id"`
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Amount   sdk.Coins      `json:"amount" yaml:"amount"`
	Time     string         `json:"time" yaml:"time"`
}

func (t TransferOfFunds) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Id: %s
	Sender: %s
	Receiver: %s
	Amount: %s
	Time: %s`,
		t.Id,
		t.Sender,
		t.Receiver,
		t.Amount,
		t.Time,
	))
}
