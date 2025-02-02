package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/elys-network/elys/x/commitment/types"
)

// UpdateVestingInfo add/update specific vesting info by denom on Params
func (k msgServer) UpdateVestingInfo(goCtx context.Context, msg *types.MsgUpdateVestingInfo) (*types.MsgUpdateVestingInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	params := k.GetParams(ctx)
	vestingInfo, index := k.GetVestingInfo(ctx, msg.BaseDenom)
	if vestingInfo == nil {
		vestingInfo = &types.VestingInfo{
			BaseDenom:       msg.BaseDenom,
			VestingDenom:    msg.VestingDenom,
			EpochIdentifier: msg.EpochIdentifier,
			NumEpochs:       msg.NumEpochs,
			VestNowFactor:   sdk.NewInt(msg.VestNowFactor),
			NumMaxVestings:  msg.NumMaxVestings,
		}
		params.VestingInfos = append(params.VestingInfos, vestingInfo)
	} else {
		params.VestingInfos[index].BaseDenom = msg.BaseDenom
		params.VestingInfos[index].VestingDenom = msg.VestingDenom
		params.VestingInfos[index].EpochIdentifier = msg.EpochIdentifier
		params.VestingInfos[index].NumEpochs = msg.NumEpochs
		params.VestingInfos[index].VestNowFactor = sdk.NewInt(msg.VestNowFactor)
		params.VestingInfos[index].NumMaxVestings = msg.NumMaxVestings
	}

	// store params
	k.SetParams(ctx, params)
	return &types.MsgUpdateVestingInfoResponse{}, nil
}
