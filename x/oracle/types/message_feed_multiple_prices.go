package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFeedMultiplePrices = "feed_multiple_prices"

var _ sdk.Msg = &MsgFeedMultiplePrices{}

func NewMsgFeedMultiplePrices(creator string) *MsgFeedMultiplePrices {
	return &MsgFeedMultiplePrices{
		Creator: creator,
	}
}

func (msg *MsgFeedMultiplePrices) Route() string {
	return RouterKey
}

func (msg *MsgFeedMultiplePrices) Type() string {
	return TypeMsgFeedMultiplePrices
}

func (msg *MsgFeedMultiplePrices) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFeedMultiplePrices) GetSignBytes() []byte {
	bz := ModuleAminoCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFeedMultiplePrices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	for _, price := range msg.Prices {
		if price.Price.IsNil() {
			return errorsmod.Wrapf(ErrInvalidPrice, "price is nil")
		}

		if price.Price.IsNegative() {
			return errorsmod.Wrapf(ErrInvalidPrice, "price is negative")
		}
	}
	return nil
}
