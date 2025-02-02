package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/elys-network/elys/x/oracle/keeper"
	types "github.com/elys-network/elys/x/oracle/types"
)

func (m *Messenger) msgFeedMultiplePrices(ctx sdk.Context, contractAddr sdk.AccAddress, msg *types.MsgFeedMultiplePrices) ([]sdk.Event, [][]byte, error) {
	if msg == nil {
		return nil, nil, wasmvmtypes.InvalidRequest{Err: "FeedMultiplePrices null msg"}
	}

	msgServer := keeper.NewMsgServerImpl(*m.keeper)

	msgFeedMultiplePrices := types.NewMsgFeedMultiplePrices(
		msg.Creator,
	)

	if err := msgFeedMultiplePrices.ValidateBasic(); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed validating msgFeedMultiplePrices")
	}

	res, err := msgServer.FeedMultiplePrices(
		sdk.WrapSDKContext(ctx),
		msgFeedMultiplePrices,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "FeedMultiplePrices msg")
	}

	responseBytes, err := json.Marshal(*res)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to serialize FeedMultiplePrices response")
	}

	resp := [][]byte{responseBytes}

	return nil, resp, nil
}
