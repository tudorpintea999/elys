package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/margin/types"
)

func (k Keeper) CheckSamePosition(ctx sdk.Context, msg *types.MsgOpen) *types.MTP {
	mtps := k.GetAllMTPs(ctx)
	for _, mtp := range mtps {
		if mtp.Address == msg.Creator && mtp.Position == msg.Position {
			return &mtp
		}
	}

	return nil
}
