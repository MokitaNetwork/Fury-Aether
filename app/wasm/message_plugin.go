package wasm

import (
	"encoding/json"

	esmkeeper "github.com/mokitanetwork/aether/x/esm/keeper"
	vaultkeeper "github.com/mokitanetwork/aether/x/vault/keeper"

	auctionkeeper "github.com/mokitanetwork/aether/x/auction/keeper"
	liquidationkeeper "github.com/mokitanetwork/aether/x/liquidation/keeper"
	tokenmintkeeper "github.com/mokitanetwork/aether/x/tokenmint/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mokitanetwork/aether/app/wasm/bindings"
	assetkeeper "github.com/mokitanetwork/aether/x/asset/keeper"
	collectorkeeper "github.com/mokitanetwork/aether/x/collector/keeper"
	lockerkeeper "github.com/mokitanetwork/aether/x/locker/keeper"
	lockertypes "github.com/mokitanetwork/aether/x/locker/types"
	rewardskeeper "github.com/mokitanetwork/aether/x/rewards/keeper"
	rewardstypes "github.com/mokitanetwork/aether/x/rewards/types"
)

func CustomMessageDecorator(lockerKeeper lockerkeeper.Keeper, rewardsKeeper rewardskeeper.Keeper,
	assetKeeper assetkeeper.Keeper, collectorKeeper collectorkeeper.Keeper, liquidationKeeper liquidationkeeper.Keeper,
	auctionKeeper auctionkeeper.Keeper, tokenMintKeeper tokenmintkeeper.Keeper, esmKeeper esmkeeper.Keeper, vaultKeeper vaultkeeper.Keeper,
) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:           old,
			lockerKeeper:      lockerKeeper,
			rewardsKeeper:     rewardsKeeper,
			assetKeeper:       assetKeeper,
			collectorKeeper:   collectorKeeper,
			liquidationKeeper: liquidationKeeper,
			auctionKeeper:     auctionKeeper,
			tokenMintKeeper:   tokenMintKeeper,
			esmKeeper:         esmKeeper,
			vaultKeeper:       vaultKeeper,
		}
	}
}

type CustomMessenger struct {
	wrapped           wasmkeeper.Messenger
	lockerKeeper      lockerkeeper.Keeper
	rewardsKeeper     rewardskeeper.Keeper
	assetKeeper       assetkeeper.Keeper
	collectorKeeper   collectorkeeper.Keeper
	liquidationKeeper liquidationkeeper.Keeper
	auctionKeeper     auctionkeeper.Keeper
	tokenMintKeeper   tokenmintkeeper.Keeper
	esmKeeper         esmkeeper.Keeper
	vaultKeeper       vaultkeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path where this is really minting / swapping ...
		// leave everything else for the wrapped version
		var aetherMsg bindings.ComdexMessages
		if err := json.Unmarshal(msg.Custom, &aetherMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "aether msg error")
		}
		if aetherMsg.MsgWhiteListAssetLocker != nil {
			return m.whitelistAssetLocker(ctx, contractAddr, aetherMsg.MsgWhiteListAssetLocker)
		}
		if aetherMsg.MsgWhitelistAppIDLockerRewards != nil {
			return m.whitelistAppIDLockerRewards(ctx, contractAddr, aetherMsg.MsgWhitelistAppIDLockerRewards)
		}
		if aetherMsg.MsgWhitelistAppIDVaultInterest != nil {
			return m.whitelistAppIDVaultInterest(ctx, contractAddr, aetherMsg.MsgWhitelistAppIDVaultInterest)
		}
		if aetherMsg.MsgAddExtendedPairsVault != nil {
			return m.AddExtendedPairsVault(ctx, contractAddr, aetherMsg.MsgAddExtendedPairsVault)
		}
		if aetherMsg.MsgSetCollectorLookupTable != nil {
			return m.SetCollectorLookupTable(ctx, contractAddr, aetherMsg.MsgSetCollectorLookupTable)
		}
		if aetherMsg.MsgSetAuctionMappingForApp != nil {
			return m.SetAuctionMappingForApp(ctx, contractAddr, aetherMsg.MsgSetAuctionMappingForApp)
		}
		if aetherMsg.MsgUpdatePairsVault != nil {
			return m.UpdatePairsVault(ctx, contractAddr, aetherMsg.MsgUpdatePairsVault)
		}
		if aetherMsg.MsgUpdateCollectorLookupTable != nil {
			return m.UpdateCollectorLookupTable(ctx, contractAddr, aetherMsg.MsgUpdateCollectorLookupTable)
		}
		if aetherMsg.MsgRemoveWhitelistAssetLocker != nil {
			return m.RemoveWhitelistAssetLocker(ctx, contractAddr, aetherMsg.MsgRemoveWhitelistAssetLocker)
		}
		if aetherMsg.MsgRemoveWhitelistAppIDVaultInterest != nil {
			return m.RemoveWhitelistAppIDVaultInterest(ctx, contractAddr, aetherMsg.MsgRemoveWhitelistAppIDVaultInterest)
		}
		if aetherMsg.MsgWhitelistAppIDLiquidation != nil {
			return m.WhitelistAppIDLiquidation(ctx, contractAddr, aetherMsg.MsgWhitelistAppIDLiquidation)
		}
		if aetherMsg.MsgRemoveWhitelistAppIDLiquidation != nil {
			return m.RemoveWhitelistAppIDLiquidation(ctx, contractAddr, aetherMsg.MsgRemoveWhitelistAppIDLiquidation)
		}
		if aetherMsg.MsgAddAuctionParams != nil {
			return m.AddAuctionParams(ctx, contractAddr, aetherMsg.MsgAddAuctionParams)
		}
		if aetherMsg.MsgBurnGovTokensForApp != nil {
			return m.BurnGovTokensForApp(ctx, contractAddr, aetherMsg.MsgBurnGovTokensForApp)
		}
		if aetherMsg.MsgAddESMTriggerParams != nil {
			return m.AddESMTriggerParams(ctx, contractAddr, aetherMsg.MsgAddESMTriggerParams)
		}
		if aetherMsg.MsgEmissionRewards != nil {
			return m.ExecuteAddEmissionRewards(ctx, contractAddr, aetherMsg.MsgEmissionRewards)
		}
		if aetherMsg.MsgFoundationEmission != nil {
			return m.ExecuteFoundationEmission(ctx, contractAddr, aetherMsg.MsgFoundationEmission)
		}
		if aetherMsg.MsgRebaseMint != nil {
			return m.ExecuteMsgRebaseMint(ctx, contractAddr, aetherMsg.MsgRebaseMint)
		}
		if aetherMsg.MsgGetSurplusFund != nil {
			return m.ExecuteMsgGetSurplusFund(ctx, contractAddr, aetherMsg.MsgGetSurplusFund)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

func (m *CustomMessenger) whitelistAssetLocker(ctx sdk.Context, contractAddr sdk.AccAddress, whiteListAsset *bindings.MsgWhiteListAssetLocker) ([]sdk.Event, [][]byte, error) {
	err := WhiteListAsset(m.lockerKeeper, ctx, contractAddr.String(), whiteListAsset)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "white list asset")
	}
	return nil, nil, nil
}

func WhiteListAsset(lockerKeeper lockerkeeper.Keeper, ctx sdk.Context, contractAddr string,
	whiteListAsset *bindings.MsgWhiteListAssetLocker,
) error {
	msg := lockertypes.MsgAddWhiteListedAssetRequest{
		From:    contractAddr,
		AppId:   whiteListAsset.AppID,
		AssetId: whiteListAsset.AssetID,
	}
	_, err := lockerKeeper.AddWhiteListedAsset(ctx, &msg)
	if err != nil {
		return err
	}

	return nil
}

func (m *CustomMessenger) whitelistAppIDLockerRewards(ctx sdk.Context, contractAddr sdk.AccAddress, whiteListAsset *bindings.MsgWhitelistAppIDLockerRewards) ([]sdk.Event, [][]byte, error) {
	err := WhitelistAppIDLockerRewards(m.rewardsKeeper, ctx, contractAddr.String(), whiteListAsset)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "white list appId locker rewards")
	}
	return nil, nil, nil
}

func WhitelistAppIDLockerRewards(rewardsKeeper rewardskeeper.Keeper, ctx sdk.Context, contractAddr string,
	whiteListAsset *bindings.MsgWhitelistAppIDLockerRewards,
) error {
	msg := rewardstypes.WhitelistAsset{
		From:         contractAddr,
		AppMappingId: whiteListAsset.AppID,
		AssetId:      whiteListAsset.AssetID,
	}
	_, err := rewardsKeeper.Whitelist(ctx, &msg)
	if err != nil {
		return err
	}

	return nil
}

func (m *CustomMessenger) whitelistAppIDVaultInterest(ctx sdk.Context, contractAddr sdk.AccAddress, whiteListAsset *bindings.MsgWhitelistAppIDVaultInterest) ([]sdk.Event, [][]byte, error) {
	err := WhitelistAppIDVaultInterest(m.rewardsKeeper, ctx, contractAddr.String(), whiteListAsset)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "white list appId vault Interest")
	}
	return nil, nil, nil
}

func WhitelistAppIDVaultInterest(rewardsKeeper rewardskeeper.Keeper, ctx sdk.Context, contractAddr string,
	whiteListAsset *bindings.MsgWhitelistAppIDVaultInterest,
) error {
	msg := rewardstypes.WhitelistAppIdVault{
		From:         contractAddr,
		AppMappingId: whiteListAsset.AppID,
	}
	_, err := rewardsKeeper.WhitelistAppVault(ctx, &msg)
	if err != nil {
		return err
	}

	return nil
}

func (m *CustomMessenger) AddExtendedPairsVault(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgAddExtendedPairsVault) ([]sdk.Event, [][]byte, error) {
	err := MsgAddExtendedPairsVault(m.assetKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "AddExtendedPairsVault error")
	}
	return nil, nil, nil
}

func MsgAddExtendedPairsVault(assetKeeper assetkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	pairVaultBinding *bindings.MsgAddExtendedPairsVault,
) error {
	err := assetKeeper.WasmAddExtendedPairsVaultRecords(ctx, pairVaultBinding)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) SetCollectorLookupTable(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgSetCollectorLookupTable) ([]sdk.Event, [][]byte, error) {
	err := MsgSetCollectorLookupTable(m.collectorKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "SetCollectorLookupTable error")
	}
	return nil, nil, nil
}

func MsgSetCollectorLookupTable(collectorKeeper collectorkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	collectorBindings *bindings.MsgSetCollectorLookupTable,
) error {
	err := collectorKeeper.WasmSetCollectorLookupTable(ctx, collectorBindings)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) SetAuctionMappingForApp(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgSetAuctionMappingForApp) ([]sdk.Event, [][]byte, error) {
	err := MsgSetAuctionMappingForApp(m.collectorKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "SetAuctionMappingForApp error")
	}
	return nil, nil, nil
}

func MsgSetAuctionMappingForApp(collectorKeeper collectorkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	auctionMappingBinding *bindings.MsgSetAuctionMappingForApp,
) error {
	err := collectorKeeper.WasmSetAuctionMappingForApp(ctx, auctionMappingBinding)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) UpdatePairsVault(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgUpdatePairsVault) ([]sdk.Event, [][]byte, error) {
	err := MsgUpdatePairsVault(m.assetKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "UpdatePairsVault error")
	}
	return nil, nil, nil
}

func MsgUpdatePairsVault(assetKeeper assetkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	updatePairVault *bindings.MsgUpdatePairsVault,
) error {
	err := assetKeeper.WasmUpdatePairsVault(ctx, updatePairVault)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) UpdateCollectorLookupTable(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgUpdateCollectorLookupTable) ([]sdk.Event, [][]byte, error) {
	err := MsgUpdateCollectorLookupTable(m.collectorKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "UpdateCollectorLookupTable error")
	}
	return nil, nil, nil
}

func MsgUpdateCollectorLookupTable(collectorKeeper collectorkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	updateColBinding *bindings.MsgUpdateCollectorLookupTable,
) error {
	err := collectorKeeper.WasmUpdateCollectorLookupTable(ctx, updateColBinding)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) RemoveWhitelistAssetLocker(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgRemoveWhitelistAssetLocker) ([]sdk.Event, [][]byte, error) {
	err := MsgRemoveWhitelistAssetLocker(m.rewardsKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "RemoveWhitelistAssetRewards error")
	}
	return nil, nil, nil
}

func MsgRemoveWhitelistAssetLocker(rewardsKeeper rewardskeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	a *bindings.MsgRemoveWhitelistAssetLocker,
) error {
	err := rewardsKeeper.WasmRemoveWhitelistAssetLocker(ctx, a.AppID, a.AssetID)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) RemoveWhitelistAppIDVaultInterest(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgRemoveWhitelistAppIDVaultInterest) ([]sdk.Event, [][]byte, error) {
	err := MsgRemoveWhitelistAppIDVaultInterest(m.rewardsKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "RemoveWhitelistAppIdVaultInterest error")
	}
	return nil, nil, nil
}

func MsgRemoveWhitelistAppIDVaultInterest(rewardsKeeper rewardskeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	a *bindings.MsgRemoveWhitelistAppIDVaultInterest,
) error {
	err := rewardsKeeper.WasmRemoveWhitelistAppIDVaultInterest(ctx, a.AppMappingID)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) WhitelistAppIDLiquidation(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgWhitelistAppIDLiquidation) ([]sdk.Event, [][]byte, error) {
	err := MsgWhitelistAppIDLiquidation(m.liquidationKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "WhitelistAppIdLiquidation error")
	}
	return nil, nil, nil
}

func MsgWhitelistAppIDLiquidation(liquidationKeeper liquidationkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	a *bindings.MsgWhitelistAppIDLiquidation,
) error {
	err := liquidationKeeper.WasmWhitelistAppIDLiquidation(ctx, a.AppID)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) RemoveWhitelistAppIDLiquidation(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgRemoveWhitelistAppIDLiquidation) ([]sdk.Event, [][]byte, error) {
	err := MsgRemoveWhitelistAppIDLiquidation(m.liquidationKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "RemoveWhitelistAppIdLiquidation error")
	}
	return nil, nil, nil
}

func MsgRemoveWhitelistAppIDLiquidation(liquidationKeeper liquidationkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	a *bindings.MsgRemoveWhitelistAppIDLiquidation,
) error {
	err := liquidationKeeper.WasmRemoveWhitelistAppIDLiquidation(ctx, a.AppID)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) AddAuctionParams(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgAddAuctionParams) ([]sdk.Event, [][]byte, error) {
	err := MsgAddAuctionParams(m.auctionKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "AddAuctionParams error")
	}
	return nil, nil, nil
}

func MsgAddAuctionParams(auctionKeeper auctionkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	auctionParamsBinding *bindings.MsgAddAuctionParams,
) error {
	err := auctionKeeper.AddAuctionParams(ctx, auctionParamsBinding)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) BurnGovTokensForApp(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgBurnGovTokensForApp) ([]sdk.Event, [][]byte, error) {
	err := MsgBurnGovTokensForApp(m.tokenMintKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "BurnGovTokensForApp error")
	}
	return nil, nil, nil
}

func MsgBurnGovTokensForApp(tokenMintKeeper tokenmintkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	a *bindings.MsgBurnGovTokensForApp,
) error {
	err := tokenMintKeeper.BurnGovTokensForApp(ctx, a.AppID, a.From, a.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) AddESMTriggerParams(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgAddESMTriggerParams) ([]sdk.Event, [][]byte, error) {
	err := MsgAddESMTriggerParams(m.esmKeeper, ctx, contractAddr, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "BurnGovTokensForApp error")
	}
	return nil, nil, nil
}

func MsgAddESMTriggerParams(esmKeeper esmkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress,
	a *bindings.MsgAddESMTriggerParams,
) error {
	err := esmKeeper.AddESMTriggerParamsForApp(ctx, a)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) ExecuteAddEmissionRewards(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgEmissionRewards) ([]sdk.Event, [][]byte, error) {
	err := MsgAddEmissionRewards(m.vaultKeeper, ctx, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "Emission rewards error")
	}
	return nil, nil, nil
}

func MsgAddEmissionRewards(vaultKeeper vaultkeeper.Keeper, ctx sdk.Context,
	a *bindings.MsgEmissionRewards,
) error {
	err := vaultKeeper.WasmMsgAddEmissionRewards(ctx, a.AppID, a.Amount, a.ExtendedPair, a.VotingRatio)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) ExecuteFoundationEmission(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgFoundationEmission) ([]sdk.Event, [][]byte, error) {
	err := MsgFoundationEmission(m.tokenMintKeeper, ctx, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "Foundation Emission rewards error")
	}
	return nil, nil, nil
}

func MsgFoundationEmission(tokenmintKeeper tokenmintkeeper.Keeper, ctx sdk.Context,
	a *bindings.MsgFoundationEmission,
) error {
	err := tokenmintKeeper.WasmMsgFoundationEmission(ctx, a.AppID, a.Amount, a.FoundationAddress)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) ExecuteMsgRebaseMint(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgRebaseMint) ([]sdk.Event, [][]byte, error) {
	err := MsgRebaseMint(m.tokenMintKeeper, ctx, a)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "Foundation Emission rewards error")
	}
	return nil, nil, nil
}

func MsgRebaseMint(tokenmintKeeper tokenmintkeeper.Keeper, ctx sdk.Context,
	a *bindings.MsgRebaseMint,
) error {
	err := tokenmintKeeper.WasmMsgRebaseMint(ctx, a.AppID, a.Amount, a.ContractAddr)
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomMessenger) ExecuteMsgGetSurplusFund(ctx sdk.Context, contractAddr sdk.AccAddress, a *bindings.MsgGetSurplusFund) ([]sdk.Event, [][]byte, error) {
	err := MsgGetSurplusFund(m.collectorKeeper, ctx, a, contractAddr)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "Execute surplus fund rewards error")
	}
	return nil, nil, nil
}

func MsgGetSurplusFund(collectorKeeper collectorkeeper.Keeper, ctx sdk.Context,
	a *bindings.MsgGetSurplusFund, contractAddr sdk.AccAddress,
) error {
	err := collectorKeeper.WasmMsgGetSurplusFund(ctx, a.AppID, a.AssetID, contractAddr, a.Amount)
	if err != nil {
		return err
	}
	return nil
}
