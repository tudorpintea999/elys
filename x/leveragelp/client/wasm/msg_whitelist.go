package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/elys-network/elys/x/leveragelp/keeper"
	types "github.com/elys-network/elys/x/leveragelp/types"
)

func (m *Messenger) msgWhitelist(ctx sdk.Context, contractAddr sdk.AccAddress, msg *types.MsgWhitelist) ([]sdk.Event, [][]byte, error) {
	if msg == nil {
		return nil, nil, wasmvmtypes.InvalidRequest{Err: "Whitelist null msg"}
	}

	msgServer := keeper.NewMsgServerImpl(*m.keeper)

	msgWhitelist := types.NewMsgWhitelist(
		msg.Authority,
		msg.WhitelistedAddress,
	)

	if err := msgWhitelist.ValidateBasic(); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed validating msgWhitelist")
	}

	res, err := msgServer.Whitelist(
		sdk.WrapSDKContext(ctx),
		msgWhitelist,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "Whitelist msg")
	}

	responseBytes, err := json.Marshal(*res)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to serialize Whitelist response")
	}

	resp := [][]byte{responseBytes}

	return nil, resp, nil
}
