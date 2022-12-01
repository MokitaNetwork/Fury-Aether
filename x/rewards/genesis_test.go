package rewards_test

import (
	"testing"

	"github.com/mokitanetwork/aether/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/mokitanetwork/aether/x/rewards"
	"github.com/mokitanetwork/aether/x/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	aetherApp := app.Setup(false)
	ctx := aetherApp.BaseApp.NewContext(false, tmproto.Header{})

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	rewards.InitGenesis(ctx, aetherApp.Rewardskeeper, &genesisState)
	got := rewards.ExportGenesis(ctx, aetherApp.Rewardskeeper)
	require.NotNil(t, got)
}
