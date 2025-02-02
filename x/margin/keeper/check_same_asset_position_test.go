package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simapp "github.com/elys-network/elys/app"
	"github.com/elys-network/elys/x/margin/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"
	"github.com/stretchr/testify/assert"
)

func TestCheckSameAssets_NewPosition(t *testing.T) {
	app := simapp.InitElysTestApp(true)
	ctx := app.BaseApp.NewContext(true, tmproto.Header{})

	k := app.MarginKeeper

	mtp := types.NewMTP("creator", ptypes.BaseCurrency, ptypes.ATOM, types.Position_LONG, sdk.NewDec(5), sdk.MustNewDecFromStr(types.TakeProfitPriceDefault), 1)
	k.SetMTP(ctx, mtp)

	msg := &types.MsgOpen{
		Creator:          "creator",
		CollateralAsset:  ptypes.ATOM,
		CollateralAmount: sdk.NewInt(100),
		BorrowAsset:      ptypes.ATOM,
		Position:         types.Position_SHORT,
		Leverage:         sdk.NewDec(1),
	}

	mtp = k.CheckSamePosition(ctx, msg)

	// Expect no error
	assert.Nil(t, mtp)
}
