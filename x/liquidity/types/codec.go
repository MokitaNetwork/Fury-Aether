package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers the necessary x/liquidity interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePair{}, "aether/liquidity/MsgCreatePair", nil)
	cdc.RegisterConcrete(&MsgCreatePool{}, "aether/liquidity/MsgCreatePool", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "aether/liquidity/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "aether/liquidity/MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgLimitOrder{}, "aether/liquidity/MsgLimitOrder", nil)
	cdc.RegisterConcrete(&MsgMarketOrder{}, "aether/liquidity/MsgMarketOrder", nil)
	cdc.RegisterConcrete(&MsgCancelOrder{}, "aether/liquidity/MsgCancelOrder", nil)
	cdc.RegisterConcrete(&MsgCancelAllOrders{}, "aether/liquidity/MsgCancelAllOrders", nil)
	cdc.RegisterConcrete(&MsgFarm{}, "aether/liquidity/MsgFarm", nil)
	cdc.RegisterConcrete(&MsgUnfarm{}, "aether/liquidity/MsgUnfarm", nil)
	cdc.RegisterConcrete(&UpdateGenericParamsProposal{}, "aether/liquidity/UpdateGenericParamsProposal", nil)
	cdc.RegisterConcrete(&CreateNewLiquidityPairProposal{}, "aether/liquidity/CreateNewLiquidityPairProposal", nil)
}

// RegisterInterfaces registers the x/liquidity interfaces types with the
// interface registry.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&UpdateGenericParamsProposal{},
		&CreateNewLiquidityPairProposal{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreatePair{},
		&MsgCreatePool{},
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgLimitOrder{},
		&MsgMarketOrder{},
		&MsgCancelOrder{},
		&MsgCancelAllOrders{},
		&MsgFarm{},
		&MsgUnfarm{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
