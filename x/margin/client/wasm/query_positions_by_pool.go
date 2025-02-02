package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/margin/types"
)

func (oq *Querier) queryPositionsByPool(ctx sdk.Context, query *types.PositionsByPoolRequest) ([]byte, error) {
	res, err := oq.keeper.GetPositionsByPool(ctx, query)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to get positions by pool")
	}

	responseBytes, err := json.Marshal(res)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to serialize positions by pool response")
	}
	return responseBytes, nil
}
