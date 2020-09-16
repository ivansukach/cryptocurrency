package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
	"time"
)

type TransferOfFunds struct {
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Amount   sdk.Coins      `json:"amount" yaml:"amount"`
}
type TransferOfFundsWithTime struct {
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Amount   sdk.Coins      `json:"amount" yaml:"amount"`
	Time     time.Time      `json:"time" yaml:"time"`
}

// implement fmt.Stringer
func (t TransferOfFunds) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Sender: %s
	Receiver: %s
	Amount: %s`,
		t.Sender,
		t.Receiver,
		t.Amount,
	))
}
