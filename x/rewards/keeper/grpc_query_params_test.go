package keeper_test

import (
	"testing"

	"github.com/mokitanetwork/aether/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/mokitanetwork/aether/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	aetherApp := app.Setup(false)
	ctx := aetherApp.BaseApp.NewContext(false, tmproto.Header{})

	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	aetherApp.Rewardskeeper.SetParams(ctx, params)

	response, err := aetherApp.Rewardskeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
