package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	assetkeeper "github.com/mokitanetwork/aether/x/asset/keeper"
	auctionKeeper "github.com/mokitanetwork/aether/x/auction/keeper"
	collectorKeeper "github.com/mokitanetwork/aether/x/collector/keeper"
	esmKeeper "github.com/mokitanetwork/aether/x/esm/keeper"
	lendKeeper "github.com/mokitanetwork/aether/x/lend/keeper"
	liquidationKeeper "github.com/mokitanetwork/aether/x/liquidation/keeper"
	liquidityKeeper "github.com/mokitanetwork/aether/x/liquidity/keeper"
	lockerkeeper "github.com/mokitanetwork/aether/x/locker/keeper"
	rewardsKeeper "github.com/mokitanetwork/aether/x/rewards/keeper"
	tokenMintkeeper "github.com/mokitanetwork/aether/x/tokenmint/keeper"
	vaultKeeper "github.com/mokitanetwork/aether/x/vault/keeper"
)

func RegisterCustomPlugins(
	locker *lockerkeeper.Keeper,
	tokenMint *tokenMintkeeper.Keeper,
	asset *assetkeeper.Keeper,
	rewards *rewardsKeeper.Keeper,
	collector *collectorKeeper.Keeper,
	liquidation *liquidationKeeper.Keeper,
	auction *auctionKeeper.Keeper,
	esm *esmKeeper.Keeper,
	vault *vaultKeeper.Keeper,
	lend *lendKeeper.Keeper,
	liquidity *liquidityKeeper.Keeper,
) []wasmkeeper.Option {
	aetherQueryPlugin := NewQueryPlugin(asset, locker, tokenMint, rewards, collector, liquidation, esm, vault, lend, liquidity)

	appDataQueryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(aetherQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(*locker, *rewards, *asset, *collector, *liquidation, *auction, *tokenMint, *esm, *vault),
	)

	return []wasm.Option{
		appDataQueryPluginOpt,
		messengerDecoratorOpt,
	}
}
