package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ammtypes "github.com/elys-network/elys/x/amm/types"
)

type PerpetualHooks interface {
	// AfterPerpetualPositionOpen is called after OpenLong or OpenShort position.
	AfterPerpetualPositionOpen(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error

	// AfterPerpetualPositionModified is called after a position gets modified.
	AfterPerpetualPositionModified(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error

	// AfterPerpetualPositionClosed is called after a position gets closed.
	AfterPerpetualPositionClosed(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error

	// AfterPoolCreated is called after CreatePool
	AfterAmmPoolCreated(ctx sdk.Context, ammPool ammtypes.Pool, sender sdk.AccAddress) error

	// AfterJoinPool is called after JoinPool, JoinSwapExternAmountIn, and JoinSwapShareAmountOut
	AfterAmmJoinPool(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error

	// AfterExitPool is called after ExitPool, ExitSwapShareAmountIn, and ExitSwapExternAmountOut
	AfterAmmExitPool(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error

	// AfterSwap is called after SwapExactAmountIn and SwapExactAmountOut
	AfterAmmSwap(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error
}

var _ PerpetualHooks = MultiPerpetualHooks{}

// combine multiple perpetual hooks, all hook functions are run in array sequence.
type MultiPerpetualHooks []PerpetualHooks

// Creates hooks for the Amm Module.
func NewMultiPerpetualHooks(hooks ...PerpetualHooks) MultiPerpetualHooks {
	return hooks
}

func (h MultiPerpetualHooks) AfterPerpetualPositionOpen(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error {
	for i := range h {
		err := h[i].AfterPerpetualPositionOpen(ctx, ammPool, perpetualPool, sender)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h MultiPerpetualHooks) AfterPerpetualPositionModified(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error {
	for i := range h {
		err := h[i].AfterPerpetualPositionModified(ctx, ammPool, perpetualPool, sender)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h MultiPerpetualHooks) AfterPerpetualPositionClosed(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error {
	for i := range h {
		err := h[i].AfterPerpetualPositionClosed(ctx, ammPool, perpetualPool, sender)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h MultiPerpetualHooks) AfterAmmPoolCreated(ctx sdk.Context, ammPool ammtypes.Pool, sender sdk.AccAddress) error {
	for i := range h {
		err := h[i].AfterAmmPoolCreated(ctx, ammPool, sender)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h MultiPerpetualHooks) AfterAmmJoinPool(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error {
	for i := range h {
		err := h[i].AfterAmmJoinPool(ctx, ammPool, perpetualPool, sender)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h MultiPerpetualHooks) AfterAmmExitPool(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error {
	for i := range h {
		err := h[i].AfterAmmExitPool(ctx, ammPool, perpetualPool, sender)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h MultiPerpetualHooks) AfterAmmSwap(ctx sdk.Context, ammPool ammtypes.Pool, perpetualPool Pool, sender sdk.AccAddress) error {
	for i := range h {
		err := h[i].AfterAmmSwap(ctx, ammPool, perpetualPool, sender)
		if err != nil {
			return err
		}
	}
	return nil
}
