package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMakeTransferOfFunds{}

// Msg<Action> - struct for unjailing jailed validator
type MsgMakeTransferOfFunds struct {
	Sender   sdk.AccAddress `json:"sender_address" yaml:"sender_address"` // address of the validator operator
	Receiver sdk.AccAddress `json:"receiver_address" yaml:"receiver_address"`
	Amount   sdk.Coins      `json:"amount" yaml:"amount"`
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsgMakeTransferOfFunds(sender sdk.AccAddress, receiver sdk.AccAddress, amount sdk.Coins) MsgMakeTransferOfFunds {
	return MsgMakeTransferOfFunds{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}

const MakeTransferOfFundsConst = "MakeTransferOfFunds"

// nolint
func (msg MsgMakeTransferOfFunds) Route() string { return RouterKey }
func (msg MsgMakeTransferOfFunds) Type() string  { return MakeTransferOfFundsConst }
func (msg MsgMakeTransferOfFunds) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
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
