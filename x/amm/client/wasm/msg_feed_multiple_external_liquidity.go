package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ammkeeper "github.com/elys-network/elys/x/amm/keeper"
	ammtypes "github.com/elys-network/elys/x/amm/types"
)

func (m *Messenger) msgFeedMultipleExternalLiquidity(ctx sdk.Context, contractAddr sdk.AccAddress, msg *ammtypes.MsgFeedMultipleExternalLiquidity) ([]sdk.Event, [][]byte, error) {
	if msg == nil {
		return nil, nil, wasmvmtypes.InvalidRequest{Err: "FeedMultipleExternalLiquidity null msg"}
	}

	msgServer := ammkeeper.NewMsgServerImpl(*m.keeper)

	msgFeedMultipleExternalLiquidity := ammtypes.NewMsgFeedMultipleExternalLiquidity(msg.Sender)

	if err := msgFeedMultipleExternalLiquidity.ValidateBasic(); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed validating msgFeedMultipleExternalLiquidity")
	}

	res, err := msgServer.FeedMultipleExternalLiquidity(
		sdk.WrapSDKContext(ctx),
		msgFeedMultipleExternalLiquidity,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "FeedMultipleExternalLiquidity msg")
	}

	responseBytes, err := json.Marshal(*res)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to serialize FeedMultipleExternalLiquidity response")
	}

	resp := [][]byte{responseBytes}

	return nil, resp, nil
}
