package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/elys-network/elys/x/oracle/keeper"
	types "github.com/elys-network/elys/x/oracle/types"
)

func (m *Messenger) msgRequestBandPrice(ctx sdk.Context, contractAddr sdk.AccAddress, msg *types.MsgRequestBandPrice) ([]sdk.Event, [][]byte, error) {
	if msg == nil {
		return nil, nil, wasmvmtypes.InvalidRequest{Err: "RequestBandPrice null msg"}
	}

	msgServer := keeper.NewMsgServerImpl(*m.keeper)

	msgRequestBandPrice := types.NewMsgRequestBandPrice(
		msg.Creator,
		types.OracleScriptID(msg.OracleScriptID),
		msg.SourceChannel,
		msg.Calldata,
		msg.AskCount,
		msg.MinCount,
		msg.FeeLimit,
		msg.PrepareGas,
		msg.ExecuteGas,
	)

	if err := msgRequestBandPrice.ValidateBasic(); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed validating msgRequestBandPrice")
	}

	res, err := msgServer.RequestBandPrice(
		sdk.WrapSDKContext(ctx),
		msgRequestBandPrice,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "RequestBandPrice msg")
	}

	responseBytes, err := json.Marshal(*res)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to serialize RequestBandPrice response")
	}

	resp := [][]byte{responseBytes}

	return nil, resp, nil
}
