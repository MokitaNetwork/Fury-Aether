package wasm

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mokitanetwork/aether/app/wasm/bindings"
)

func CustomQuerier(queryPlugin *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var aetherQuery bindings.AetherQuery
		if err := json.Unmarshal(request, &aetherQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "app query")
		}
		if aetherQuery.AppData != nil {
			appID := aetherQuery.AppData.AppID
			MinGovDeposit, GovTimeInSeconds, assetID, _ := queryPlugin.GetAppInfo(ctx, appID)
			res := bindings.AppDataResponse{
				MinGovDeposit:    MinGovDeposit.String(),
				GovTimeInSeconds: GovTimeInSeconds,
				AssetID:          assetID,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "App data query response")
			}
			return bz, nil
		} else if aetherQuery.AssetData != nil {
			assetID := aetherQuery.AssetData.AssetID
			denom, _ := queryPlugin.GetAssetInfo(ctx, assetID)
			res := bindings.AssetDataResponse{
				Denom: denom,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "App data query response")
			}
			return bz, nil
		} else if aetherQuery.MintedToken != nil {
			appID := aetherQuery.MintedToken.AppID
			assetID := aetherQuery.MintedToken.AssetID
			MintedToken, _ := queryPlugin.GetTokenMint(ctx, appID, assetID)
			res := bindings.MintedTokenResponse{
				MintedTokens: MintedToken,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "tokenMint query response")
			}
			return bz, nil
		} else if aetherQuery.RemoveWhiteListAssetLocker != nil {
			appID := aetherQuery.RemoveWhiteListAssetLocker.AppID
			assetID := aetherQuery.RemoveWhiteListAssetLocker.AssetIDs

			found, errormsg := queryPlugin.GetRemoveWhitelistAppIDLockerRewardsCheck(ctx, appID, assetID)
			res := bindings.RemoveWhiteListAssetResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "RemoveWhiteListAssetLocker query response")
			}
			return bz, nil
		} else if aetherQuery.WhitelistAppIDLockerRewards != nil {
			appID := aetherQuery.WhitelistAppIDLockerRewards.AppID
			assetID := aetherQuery.WhitelistAppIDLockerRewards.AssetID

			found, errormsg := queryPlugin.GetWhitelistAppIDLockerRewardsCheck(ctx, appID, assetID)
			res := bindings.WhitelistAppIDLockerRewardsResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "WhitelistAppIdLockerRewards query response")
			}
			return bz, nil
		} else if aetherQuery.WhitelistAppIDVaultInterest != nil {
			appID := aetherQuery.WhitelistAppIDVaultInterest.AppID

			found, errormsg := queryPlugin.GetWhitelistAppIDVaultInterestCheck(ctx, appID)
			res := bindings.WhitelistAppIDLockerRewardsResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "WhitelistAppIdVaultInterest query response")
			}
			return bz, nil
		} else if aetherQuery.ExternalLockerRewards != nil {
			appID := aetherQuery.ExternalLockerRewards.AppID
			assetID := aetherQuery.ExternalLockerRewards.AssetID

			found, errormsg := queryPlugin.GetExternalLockerRewardsCheck(ctx, appID, assetID)
			res := bindings.WhitelistAppIDLockerRewardsResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "GetExternalLockerRewardsCheck query response")
			}
			return bz, nil
		} else if aetherQuery.ExternalVaultRewards != nil {
			appID := aetherQuery.ExternalVaultRewards.AppID
			assetID := aetherQuery.ExternalVaultRewards.AssetID

			found, errormsg := queryPlugin.GetExternalVaultRewardsCheck(ctx, appID, assetID)
			res := bindings.ExternalVaultRewardsResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "ExternalVaultRewards query response")
			}
			return bz, nil
		} else if aetherQuery.CollectorLookupTableQuery != nil {
			appID := aetherQuery.CollectorLookupTableQuery.AppID
			collectorAssetID := aetherQuery.CollectorLookupTableQuery.CollectorAssetID
			secondaryAssetID := aetherQuery.CollectorLookupTableQuery.SecondaryAssetID
			found, errormsg := queryPlugin.CollectorLookupTableQueryCheck(ctx, appID, collectorAssetID, secondaryAssetID)
			res := bindings.CollectorLookupTableQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "ExternalVaultRewards query response")
			}
			return bz, nil
		} else if aetherQuery.ExtendedPairsVaultRecordsQuery != nil {
			appID := aetherQuery.ExtendedPairsVaultRecordsQuery.AppID
			pairID := aetherQuery.ExtendedPairsVaultRecordsQuery.PairID
			StabilityFee := aetherQuery.ExtendedPairsVaultRecordsQuery.StabilityFee
			ClosingFee := aetherQuery.ExtendedPairsVaultRecordsQuery.ClosingFee
			DrawDownFee := aetherQuery.ExtendedPairsVaultRecordsQuery.DrawDownFee
			DebtCeiling := aetherQuery.ExtendedPairsVaultRecordsQuery.DebtCeiling
			DebtFloor := aetherQuery.ExtendedPairsVaultRecordsQuery.DebtFloor
			PairName := aetherQuery.ExtendedPairsVaultRecordsQuery.PairName

			found, errorMsg := queryPlugin.ExtendedPairsVaultRecordsQueryCheck(ctx, appID, pairID, StabilityFee, ClosingFee, DrawDownFee, DebtCeiling, DebtFloor, PairName)
			res := bindings.ExtendedPairsVaultRecordsQueryResponse{
				Found: found,
				Err:   errorMsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "ExternalVaultRewards query response")
			}
			return bz, nil
		} else if aetherQuery.AuctionMappingForAppQuery != nil {
			appID := aetherQuery.AuctionMappingForAppQuery.AppID
			found, errormsg := queryPlugin.AuctionMappingForAppQueryCheck(ctx, appID)
			res := bindings.AuctionMappingForAppQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "AuctionMappingForAppQuery query response")
			}
			return bz, nil
		} else if aetherQuery.WhiteListedAssetQuery != nil {
			appID := aetherQuery.WhiteListedAssetQuery.AppID
			assetID := aetherQuery.WhiteListedAssetQuery.AssetID
			found, errormsg := queryPlugin.WhiteListedAssetQueryCheck(ctx, appID, assetID)
			res := bindings.WhiteListedAssetQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "WhiteListedAssetQueryCheck query response")
			}
			return bz, nil
		} else if aetherQuery.UpdatePairsVaultQuery != nil {
			appID := aetherQuery.UpdatePairsVaultQuery.AppID
			extPairID := aetherQuery.UpdatePairsVaultQuery.ExtPairID
			found, errormsg := queryPlugin.UpdatePairsVaultQueryCheck(ctx, appID, extPairID)
			res := bindings.UpdatePairsVaultQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "UpdatePairsVaultQuery query response")
			}
			return bz, nil
		} else if aetherQuery.UpdateCollectorLookupTableQuery != nil {
			appID := aetherQuery.UpdateCollectorLookupTableQuery.AppID
			assetID := aetherQuery.UpdateCollectorLookupTableQuery.AssetID
			found, errormsg := queryPlugin.UpdateCollectorLookupTableQueryCheck(ctx, appID, assetID)
			res := bindings.UpdateCollectorLookupTableQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "UpdatePairsVaultQuery query response")
			}
			return bz, nil
		} else if aetherQuery.RemoveWhitelistAppIDVaultInterestQuery != nil {
			appID := aetherQuery.RemoveWhitelistAppIDVaultInterestQuery.AppID
			found, errormsg := queryPlugin.WasmRemoveWhitelistAppIDVaultInterestQueryCheck(ctx, appID)
			res := bindings.RemoveWhitelistAppIDVaultInterestQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "RemoveWhitelistAppIdVaultInterestQuery query response")
			}
			return bz, nil
		} else if aetherQuery.RemoveWhitelistAssetLockerQuery != nil {
			appID := aetherQuery.RemoveWhitelistAssetLockerQuery.AppID
			assetID := aetherQuery.RemoveWhitelistAssetLockerQuery.AssetID

			found, errormsg := queryPlugin.WasmRemoveWhitelistAssetLockerQueryCheck(ctx, appID, assetID)
			res := bindings.RemoveWhitelistAssetLockerQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "RemoveWhitelistAssetLockerQuery query response")
			}
			return bz, nil
		} else if aetherQuery.WhitelistAppIDLiquidationQuery != nil {
			AppID := aetherQuery.WhitelistAppIDLiquidationQuery.AppID

			found, errormsg := queryPlugin.WasmWhitelistAppIDLiquidationQueryCheck(ctx, AppID)
			res := bindings.WhitelistAppIDLiquidationQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "WhitelistAppIDLiquidationQuery query response")
			}
			return bz, nil
		} else if aetherQuery.RemoveWhitelistAppIDLiquidationQuery != nil {
			AppID := aetherQuery.RemoveWhitelistAppIDLiquidationQuery.AppID

			found, errormsg := queryPlugin.WasmRemoveWhitelistAppIDLiquidationQueryCheck(ctx, AppID)
			res := bindings.RemoveWhitelistAppIDLiquidationQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "RemoveWhitelistAppIDLiquidationQuery query response")
			}
			return bz, nil
		} else if aetherQuery.AddESMTriggerParamsForAppQuery != nil {
			AppID := aetherQuery.AddESMTriggerParamsForAppQuery.AppID

			found, errormsg := queryPlugin.WasmAddESMTriggerParamsQueryCheck(ctx, AppID)
			res := bindings.AddESMTriggerParamsForAppResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "AddESMTriggerParamsForAppResponse query response")
			}
			return bz, nil
		} else if aetherQuery.ExtendedPairByApp != nil {
			AppID := aetherQuery.ExtendedPairByApp.AppID

			extendedPair, _ := queryPlugin.WasmExtendedPairByApp(ctx, AppID)
			res := bindings.ExtendedPairByAppResponse{
				ExtendedPair: extendedPair,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "ExtendedPairByAppResponse query response")
			}
			return bz, nil
		} else if aetherQuery.CheckSurplusReward != nil {
			AppID := aetherQuery.CheckSurplusReward.AppID
			AssetID := aetherQuery.CheckSurplusReward.AssetID
			amount := queryPlugin.WasmCheckSurplusReward(ctx, AppID, AssetID)
			res := bindings.CheckSurplusRewardResponse{
				Amount: amount,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "CheckSurplusRewardResponse query response")
			}
			return bz, nil
		} else if aetherQuery.CheckWhitelistedAsset != nil {
			Denom := aetherQuery.CheckWhitelistedAsset.Denom

			found := queryPlugin.WasmCheckWhitelistedAsset(ctx, Denom)
			res := bindings.CheckWhitelistedAssetResponse{
				Found: found,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "CheckWhitelistedAssetResponse query response")
			}
			return bz, nil
		} else if aetherQuery.CheckVaultCreated != nil {
			Address := aetherQuery.CheckVaultCreated.Address
			AppID := aetherQuery.CheckVaultCreated.AppID
			found := queryPlugin.WasmCheckVaultCreated(ctx, Address, AppID)
			res := bindings.VaultCreatedResponse{
				IsCompleted: found,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "VaultCreatedResponse query response")
			}
			return bz, nil
		} else if aetherQuery.CheckBorrowed != nil {
			AssetID := aetherQuery.CheckBorrowed.AssetID
			Address := aetherQuery.CheckBorrowed.Address
			found := queryPlugin.WasmCheckBorrowed(ctx, AssetID, Address)
			res := bindings.BorrowedResponse{
				IsCompleted: found,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "BorrowedResponse query response")
			}
			return bz, nil
		} else if aetherQuery.CheckLiquidityProvided != nil {
			AppID := aetherQuery.CheckLiquidityProvided.AppID
			PoolID := aetherQuery.CheckLiquidityProvided.PoolID
			Address := aetherQuery.CheckLiquidityProvided.Address
			found := queryPlugin.WasmCheckLiquidityProvided(ctx, AppID, PoolID, Address)
			res := bindings.LiquidityProvidedResponse{
				IsCompleted: found,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "LiquidityProvidedResponse query response")
			}
			return bz, nil
		}
		return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown App Data query variant"}
	}
}
