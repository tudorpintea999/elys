package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p *Pool) CalcExitPoolCoinsFromShares(ctx sdk.Context, oracleKeeper OracleKeeper, exitingShares math.Int, tokenOutDenom string) (exitedCoins sdk.Coins, err error) {
	return CalcExitPool(ctx, oracleKeeper, *p, exitingShares, tokenOutDenom)
}