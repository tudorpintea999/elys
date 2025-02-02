package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/elys-network/elys/x/commitment/keeper"
	types "github.com/elys-network/elys/x/commitment/types"
)

func (m *Messenger) msgUncommitTokens(ctx sdk.Context, contractAddr sdk.AccAddress, msg *types.MsgUncommitTokens) ([]sdk.Event, [][]byte, error) {
	if msg == nil {
		return nil, nil, wasmvmtypes.InvalidRequest{Err: "UncommitTokens null msg"}
	}

	msgServer := keeper.NewMsgServerImpl(*m.keeper)

	msgUncommitTokens := types.NewMsgUncommitTokens(
		msg.Creator,
		msg.Amount,
		msg.Denom,
	)

	if err := msgUncommitTokens.ValidateBasic(); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed validating msgUncommitTokens")
	}

	res, err := msgServer.UncommitTokens(
		sdk.WrapSDKContext(ctx),
		msgUncommitTokens,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "UncommitTokens msg")
	}

	responseBytes, err := json.Marshal(*res)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to serialize UncommitTokens response")
	}

	resp := [][]byte{responseBytes}

	return nil, resp, nil
}
