package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/elys-network/elys/x/leveragelp/keeper"
	types "github.com/elys-network/elys/x/leveragelp/types"
)

func (m *Messenger) msgUpdatePools(ctx sdk.Context, contractAddr sdk.AccAddress, msg *types.MsgUpdatePools) ([]sdk.Event, [][]byte, error) {
	if msg == nil {
		return nil, nil, wasmvmtypes.InvalidRequest{Err: "UpdatePools null msg"}
	}

	msgServer := keeper.NewMsgServerImpl(*m.keeper)

	msgUpdatePools := types.NewMsgUpdatePools(
		msg.Authority,
		msg.Pools,
	)

	if err := msgUpdatePools.ValidateBasic(); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed validating msgUpdatePools")
	}

	res, err := msgServer.UpdatePools(
		sdk.WrapSDKContext(ctx),
		msgUpdatePools,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "UpdatePools msg")
	}

	responseBytes, err := json.Marshal(*res)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to serialize UpdatePools response")
	}

	resp := [][]byte{responseBytes}

	return nil, resp, nil
}
