package keeper_test

import (
	"testing"

	"github.com/mokitanetwork/aether/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/mokitanetwork/aether/x/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	aetherApp := app.Setup(false)
	ctx := aetherApp.BaseApp.NewContext(false, tmproto.Header{})

	params := types.DefaultParams()

	aetherApp.Rewardskeeper.SetParams(ctx, params)

	require.EqualValues(t, params, aetherApp.Rewardskeeper.GetParams(ctx))
}
