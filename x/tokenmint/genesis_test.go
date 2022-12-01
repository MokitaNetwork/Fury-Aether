package tokenmint_test

// import (
// 	"testing"

// 	keepertest "github.com/mokitanetwork/aether/testutil/keeper"
// 	"github.com/mokitanetwork/aether/testutil/nullify"
// 	"github.com/mokitanetwork/aether/x/tokenmint"
// 	"github.com/mokitanetwork/aether/x/tokenmint/types"
// 	"github.com/stretchr/testify/require"
// )

// func TestGenesis(t *testing.T) {
// 	genesisState := types.GenesisState{
// 		Params:	types.DefaultParams(),

// 	}

// 	k, ctx := keepertest.TokenmintKeeper(t)
// 	tokenmint.InitGenesis(ctx, *k, genesisState)
// 	got := tokenmint.ExportGenesis(ctx, *k)
// 	require.NotNil(t, got)

// 	nullify.Fill(&genesisState)
// 	nullify.Fill(got)

// }
