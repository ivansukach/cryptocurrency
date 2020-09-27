package octa

import (
	"fmt"
	"github.com/ivansukach/cryptocurrency/x/octa/keeper"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ivansukach/cryptocurrency/x/octa/types"
)

// NewHandler creates an sdk.Handler for all the octa type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgMakeTransferOfFunds:
			return handleMsgMakeTransferOfFunds(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreateScavenge creates a new scavenge and moves the reward into escrow
func handleMsgMakeTransferOfFunds(ctx sdk.Context, keeper keeper.Keeper, msg MsgMakeTransferOfFunds) (*sdk.Result, error) {
	var transfer = types.TransferOfFunds{
		Id:       msg.Id,
		Sender:   msg.Sender,
		Receiver: msg.Receiver,
		Amount:   msg.Amount,
		Time:     msg.Time,
	}
	//moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	//log.Println("SENDER: ", transfer.Sender)
	//log.Println("RECEIVER: ", transfer.Receiver)
	//log.Println("AMOUNT: ", transfer.Amount)
	//log.Println("ModuleAccount: ", moduleAcct)
	//sdkError := keeper.CoinKeeper.SendCoins(ctx, transfer.Sender, moduleAcct, transfer.Amount)
	//if sdkError != nil {
	//	return nil, sdkError
	//}
	//sdkError = keeper.CoinKeeper.SendCoins(ctx, moduleAcct, transfer.Receiver, transfer.Amount)
	//if sdkError != nil {
	//	return nil, sdkError
	//}
	sdkError := keeper.CoinKeeper.SendCoins(ctx, transfer.Sender, transfer.Receiver, transfer.Amount)
	if sdkError != nil {
		log.Println("Error in handler during send coins: ", sdkError)
		return nil, sdkError
	}
	log.Println("TRANSFER COMPLETED SUCCESSFULLY")
	keeper.SetTransfer(ctx, transfer)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeMakeTransferOfFunds),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(types.AttributeReceiver, msg.Receiver.String()),
			sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
