package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/elys-network/elys/x/oracle/keeper"
	types "github.com/elys-network/elys/x/oracle/types"
)

func (m *Messenger) msgFeedPrice(ctx sdk.Context, contractAddr sdk.AccAddress, msg *types.MsgFeedPrice) ([]sdk.Event, [][]byte, error) {
	if msg == nil {
		return nil, nil, wasmvmtypes.InvalidRequest{Err: "FeedPrice null msg"}
	}

	msgServer := keeper.NewMsgServerImpl(*m.keeper)

	msgFeedPrice := types.NewMsgFeedPrice(
		msg.Provider,
		msg.Asset,
		msg.Price,
		msg.Source,
	)

	if err := msgFeedPrice.ValidateBasic(); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed validating msgFeedPrice")
	}

	res, err := msgServer.FeedPrice(
		sdk.WrapSDKContext(ctx),
		msgFeedPrice,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "FeedPrice msg")
	}

	responseBytes, err := json.Marshal(*res)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to serialize FeedPrice response")
	}

	resp := [][]byte{responseBytes}

	return nil, resp, nil
}
