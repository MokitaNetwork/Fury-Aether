package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&AddAssetsProposal{}, "aether/asset/AddAssetsProposal", nil)
	cdc.RegisterConcrete(&UpdateAssetProposal{}, "aether/asset/UpdateAssetProposal", nil)
	cdc.RegisterConcrete(&AddPairsProposal{}, "aether/asset/AddPairsProposal", nil)
	cdc.RegisterConcrete(&UpdatePairProposal{}, "aether/asset/UpdatePairProposal", nil)
	cdc.RegisterConcrete(&AddAppProposal{}, "aether/asset/AddAppProposal", nil)
	cdc.RegisterConcrete(&AddAssetInAppProposal{}, "aether/asset/AddAssetInAppProposal", nil)
	cdc.RegisterConcrete(&UpdateGovTimeInAppProposal{}, "aether/asset/UpdateGovTimeInAppProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&AddAssetsProposal{},
		&UpdateAssetProposal{},
		&AddPairsProposal{},
		&UpdatePairProposal{},
		&AddAppProposal{},
		&AddAssetInAppProposal{},
		&UpdateGovTimeInAppProposal{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
