package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/elys-network/elys/x/tokenomics/keeper"
	types "github.com/elys-network/elys/x/tokenomics/types"
)

func (m *Messenger) msgUpdateTimeBasedInflation(ctx sdk.Context, contractAddr sdk.AccAddress, msg *types.MsgUpdateTimeBasedInflation) ([]sdk.Event, [][]byte, error) {
	if msg == nil {
		return nil, nil, wasmvmtypes.InvalidRequest{Err: "UpdateTimeBasedInflation null msg"}
	}

	msgServer := keeper.NewMsgServerImpl(*m.keeper)

	msgUpdateTimeBasedInflation := types.NewMsgUpdateTimeBasedInflation(
		msg.Authority,
		msg.StartBlockHeight,
		msg.EndBlockHeight,
		msg.Description,
		msg.Inflation,
	)

	if err := msgUpdateTimeBasedInflation.ValidateBasic(); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed validating msgUpdateTimeBasedInflation")
	}

	res, err := msgServer.UpdateTimeBasedInflation(
		sdk.WrapSDKContext(ctx),
		msgUpdateTimeBasedInflation,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "UpdateTimeBasedInflation msg")
	}

	responseBytes, err := json.Marshal(*res)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to serialize UpdateTimeBasedInflation response")
	}

	resp := [][]byte{responseBytes}

	return nil, resp, nil
}
