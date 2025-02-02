package client

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/elys-network/elys/wasmbindings/types"
	accountedpoolclientwasm "github.com/elys-network/elys/x/accountedpool/client/wasm"
	accountedpoolkeeper "github.com/elys-network/elys/x/accountedpool/keeper"
	ammclientwasm "github.com/elys-network/elys/x/amm/client/wasm"
	ammkeeper "github.com/elys-network/elys/x/amm/keeper"
	assetprofileclientwasm "github.com/elys-network/elys/x/assetprofile/client/wasm"
	assetprofilekeeper "github.com/elys-network/elys/x/assetprofile/keeper"
	burnerclientwasm "github.com/elys-network/elys/x/burner/client/wasm"
	burnerkeeper "github.com/elys-network/elys/x/burner/keeper"
	clockclientwasm "github.com/elys-network/elys/x/clock/client/wasm"
	clockkeeper "github.com/elys-network/elys/x/clock/keeper"
	commitmentclientwasm "github.com/elys-network/elys/x/commitment/client/wasm"
	commitmentkeeper "github.com/elys-network/elys/x/commitment/keeper"
	epochsclientwasm "github.com/elys-network/elys/x/epochs/client/wasm"
	epochskeeper "github.com/elys-network/elys/x/epochs/keeper"
	incentiveclientwasm "github.com/elys-network/elys/x/incentive/client/wasm"
	incentivekeeper "github.com/elys-network/elys/x/incentive/keeper"
	leveragelpclientwasm "github.com/elys-network/elys/x/leveragelp/client/wasm"
	leveragelpkeeper "github.com/elys-network/elys/x/leveragelp/keeper"
	marginclientwasm "github.com/elys-network/elys/x/margin/client/wasm"
	marginkeeper "github.com/elys-network/elys/x/margin/keeper"
	oracleclientwasm "github.com/elys-network/elys/x/oracle/client/wasm"
	oraclekeeper "github.com/elys-network/elys/x/oracle/keeper"
	parameterclientwasm "github.com/elys-network/elys/x/parameter/client/wasm"
	parameterkeeper "github.com/elys-network/elys/x/parameter/keeper"
	stablestakeclientwasm "github.com/elys-network/elys/x/stablestake/client/wasm"
	stablestakekeeper "github.com/elys-network/elys/x/stablestake/keeper"
	tokenomicsclientwasm "github.com/elys-network/elys/x/tokenomics/client/wasm"
	tokenomicskeeper "github.com/elys-network/elys/x/tokenomics/keeper"
	transferhookclientwasm "github.com/elys-network/elys/x/transferhook/client/wasm"
	transferhookkeeper "github.com/elys-network/elys/x/transferhook/keeper"
)

func RegisterCustomPlugins(
	accountedpool *accountedpoolkeeper.Keeper,
	amm *ammkeeper.Keeper,
	assetprofile *assetprofilekeeper.Keeper,
	bank *bankkeeper.BaseKeeper,
	burner *burnerkeeper.Keeper,
	clock *clockkeeper.Keeper,
	commitment *commitmentkeeper.Keeper,
	epochs *epochskeeper.Keeper,
	incentive *incentivekeeper.Keeper,
	leveragelp *leveragelpkeeper.Keeper,
	margin *marginkeeper.Keeper,
	oracle *oraclekeeper.Keeper,
	parameter *parameterkeeper.Keeper,
	stablestake *stablestakekeeper.Keeper,
	staking *stakingkeeper.Keeper,
	tokenomics *tokenomicskeeper.Keeper,
	transferhook *transferhookkeeper.Keeper,
) []wasmkeeper.Option {
	accountedpoolQuerier := accountedpoolclientwasm.NewQuerier(accountedpool)
	accountedpoolMessenger := accountedpoolclientwasm.NewMessenger(accountedpool)

	ammQuerier := ammclientwasm.NewQuerier(amm, bank, commitment)
	ammMessenger := ammclientwasm.NewMessenger(amm)

	assetprofileQuerier := assetprofileclientwasm.NewQuerier(assetprofile)
	assetprofileMessenger := assetprofileclientwasm.NewMessenger(assetprofile)

	burnerQuerier := burnerclientwasm.NewQuerier(burner)
	burnerMessenger := burnerclientwasm.NewMessenger(burner)

	clockQuerier := clockclientwasm.NewQuerier(clock)
	clockMessenger := clockclientwasm.NewMessenger(clock)

	commitmentQuerier := commitmentclientwasm.NewQuerier(commitment)
	commitmentMessenger := commitmentclientwasm.NewMessenger(commitment, staking)

	epochsQuerier := epochsclientwasm.NewQuerier(epochs)
	epochsMessenger := epochsclientwasm.NewMessenger(epochs)

	incentiveQuerier := incentiveclientwasm.NewQuerier(incentive)
	incentiveMessenger := incentiveclientwasm.NewMessenger(incentive, staking, commitment)

	leveragelpQuerier := leveragelpclientwasm.NewQuerier(leveragelp)
	leveragelpMessenger := leveragelpclientwasm.NewMessenger(leveragelp)

	marginQuerier := marginclientwasm.NewQuerier(margin)
	marginMessenger := marginclientwasm.NewMessenger(margin)

	oracleQuerier := oracleclientwasm.NewQuerier(oracle)
	oracleMessenger := oracleclientwasm.NewMessenger(oracle)

	parameterQuerier := parameterclientwasm.NewQuerier(parameter)
	parameterMessenger := parameterclientwasm.NewMessenger(parameter)

	stablestakeQuerier := stablestakeclientwasm.NewQuerier(stablestake)
	stablestakeMessenger := stablestakeclientwasm.NewMessenger(stablestake)

	tokenomicsQuerier := tokenomicsclientwasm.NewQuerier(tokenomics)
	tokenomicsMessenger := tokenomicsclientwasm.NewMessenger(tokenomics)

	transferhookQuerier := transferhookclientwasm.NewQuerier(transferhook)
	transferhookMessenger := transferhookclientwasm.NewMessenger(transferhook)

	moduleQueriers := []types.ModuleQuerier{
		accountedpoolQuerier,
		ammQuerier,
		assetprofileQuerier,
		burnerQuerier,
		clockQuerier,
		commitmentQuerier,
		epochsQuerier,
		incentiveQuerier,
		leveragelpQuerier,
		marginQuerier,
		oracleQuerier,
		parameterQuerier,
		stablestakeQuerier,
		tokenomicsQuerier,
		transferhookQuerier,
	}

	wasmQueryPlugin := types.NewQueryPlugin(
		moduleQueriers,
		accountedpool,
		amm,
		assetprofile,
		bank,
		burner,
		clock,
		commitment,
		epochs,
		incentive,
		leveragelp,
		margin,
		oracle,
		parameter,
		stablestake,
		staking,
		tokenomics,
		transferhook,
	)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: types.CustomQuerier(wasmQueryPlugin),
	})

	moduleMessengers := []types.ModuleMessenger{
		accountedpoolMessenger,
		ammMessenger,
		assetprofileMessenger,
		burnerMessenger,
		clockMessenger,
		commitmentMessenger,
		epochsMessenger,
		incentiveMessenger,
		leveragelpMessenger,
		marginMessenger,
		oracleMessenger,
		parameterMessenger,
		stablestakeMessenger,
		tokenomicsMessenger,
		transferhookMessenger,
	}

	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		types.CustomMessageDecorator(
			moduleMessengers,
			accountedpool,
			amm,
			assetprofile,
			bank,
			burner,
			clock,
			commitment,
			epochs,
			incentive,
			leveragelp,
			margin,
			oracle,
			parameter,
			stablestake,
			staking,
			tokenomics,
			transferhook,
		),
	)
	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
