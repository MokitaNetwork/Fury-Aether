package tokenmint

// import (
// 	"math/rand"

// 	"github.com/mokitanetwork/aether/testutil/sample"
// 	tokenmintsimulation "github.com/mokitanetwork/aether/x/tokenmint/simulation"
// 	"github.com/mokitanetwork/aether/x/tokenmint/types"
// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/module"
// 	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
// 	"github.com/cosmos/cosmos-sdk/x/simulation"
// )

// // avoid unused import issue
// var (
// 	_ = sample.AccAddress
// 	_ = tokenmintsimulation.FindAccount
// 	_ = simappparams.StakePerAccount
// 	_ = simulation.MsgEntryKind
// 	_ = baseapp.Paramspace
// )

// const (
// )

// // GenerateGenesisState creates a randomized GenState of the module
// func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
// 	accs := make([]string, len(simState.Accounts))
// 	for i, acc := range simState.Accounts {
// 		accs[i] = acc.Address.String()
// 	}
// 	tokenmintGenesis := types.GenesisState{
// 	}
// 	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenmintGenesis)
// }

// // ProposalContents doesn't return any content functions for governance proposals
// func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
// 	return nil
// }

// // RandomizedParams creates randomized  param changes for the simulator
// func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

// 	return []simtypes.ParamChange{
// 	}
// }

// // RegisterStoreDecoder registers a decoder
// func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// // WeightedOperations returns the all the gov module operations with their respective weights.
// func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
// 	operations := make([]simtypes.WeightedOperation, 0)

// 	return operations
// }