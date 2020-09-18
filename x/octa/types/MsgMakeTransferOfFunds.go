package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"log"
	"time"
)

var _ sdk.Msg = &MsgMakeTransferOfFunds{}

// Msg<Action> - struct for unjailing jailed validator
type MsgMakeTransferOfFunds struct {
	Sender   sdk.AccAddress `json:"senderAddress" yaml:"senderAddress"` // address of the validator operator
	Receiver sdk.AccAddress `json:"receiverAddress" yaml:"receiverAddress"`
	Amount   sdk.Coins      `json:"amount" yaml:"amount"`
	Time     string         `json:"time" yaml:"time"`
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsgMakeTransferOfFunds(sender sdk.AccAddress, receiver sdk.AccAddress, amount sdk.Coins) MsgMakeTransferOfFunds {
	log.Println("MSG SENDER: ", sender)
	log.Println("MSG RECEIVER: ", receiver)
	return MsgMakeTransferOfFunds{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Time:     time.Now().UTC().String(),
	}
}

const MakeTransferOfFundsConst = "MakeTransferOfFunds"

// nolint
func (msg MsgMakeTransferOfFunds) Route() string { return RouterKey }
func (msg MsgMakeTransferOfFunds) Type() string  { return MakeTransferOfFundsConst }
func (msg MsgMakeTransferOfFunds) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgMakeTransferOfFunds) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgMakeTransferOfFunds) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing receiver address")
	}
	if msg.Amount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing amount of coins")
	}
	return nil
}
