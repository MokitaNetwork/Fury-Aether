package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgLend{}, "aether/lend/MsgLend", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "aether/lend/MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "aether/lend/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgCloseLend{}, "aether/lend/MsgCloseLend", nil)
	cdc.RegisterConcrete(&MsgBorrow{}, "aether/lend/MsgBorrow", nil)
	cdc.RegisterConcrete(&MsgDepositBorrow{}, "aether/lend/MsgDepositBorrow", nil)
	cdc.RegisterConcrete(&MsgDraw{}, "aether/lend/MsgDraw", nil)
	cdc.RegisterConcrete(&MsgCloseBorrow{}, "aether/lend/MsgCloseBorrow", nil)
	cdc.RegisterConcrete(&MsgRepay{}, "aether/lend/MsgRepay", nil)
	cdc.RegisterConcrete(&MsgBorrowAlternate{}, "aether/lend/MsgBorrowAlternate", nil)
	cdc.RegisterConcrete(&MsgFundModuleAccounts{}, "aether/lend/MsgFundModuleAccounts", nil)
	cdc.RegisterConcrete(&LendPairsProposal{}, "aether/lend/LendPairsProposal", nil)
	cdc.RegisterConcrete(&AddPoolsProposal{}, "aether/lend/AddPoolsProposal", nil)
	cdc.RegisterConcrete(&AddAssetToPairProposal{}, "aether/lend/AddAssetToPairProposal", nil)
	cdc.RegisterConcrete(&AddAssetRatesParams{}, "aether/lend/AddAssetRatesParams", nil)
	cdc.RegisterConcrete(&AddAuctionParamsProposal{}, "aether/lend/AddAuctionParamsProposal", nil)
	cdc.RegisterConcrete(&MsgCalculateInterestAndRewards{}, "aether/lend/MsgCalculateInterestAndRewards", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&LendPairsProposal{},
		&AddPoolsProposal{},
		&AddAssetToPairProposal{},
		&AddAssetRatesParams{},
		&AddAuctionParamsProposal{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgLend{},
		&MsgWithdraw{},
		&MsgDeposit{},
		&MsgCloseLend{},
		&MsgBorrow{},
		&MsgDepositBorrow{},
		&MsgDraw{},
		&MsgCloseBorrow{},
		&MsgRepay{},
		&MsgBorrowAlternate{},
		&MsgFundModuleAccounts{},
		&MsgCalculateInterestAndRewards{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
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
