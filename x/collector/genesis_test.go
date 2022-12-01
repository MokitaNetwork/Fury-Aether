package collector_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	app "github.com/mokitanetwork/aether/app"
	"github.com/mokitanetwork/aether/x/collector"
	"github.com/mokitanetwork/aether/x/collector/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	aetherApp := app.Setup(false)
	ctx := aetherApp.BaseApp.NewContext(false, tmproto.Header{})
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k := aetherApp.CollectorKeeper
	collector.InitGenesis(ctx, k, &genesisState)
	got := collector.ExportGenesis(ctx, k)
	require.NotNil(t, got)
}
