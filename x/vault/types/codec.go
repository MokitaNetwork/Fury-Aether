package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateRequest{}, "aether/vault/MsgCreateRequest", nil)
	cdc.RegisterConcrete(&MsgCloseRequest{}, "aether/vault/MsgCloseRequest", nil)
	cdc.RegisterConcrete(&MsgDepositRequest{}, "aether/vault/MsgDepositRequest", nil)
	cdc.RegisterConcrete(&MsgRepayRequest{}, "aether/vault/MsgRepayRequest", nil)
	cdc.RegisterConcrete(&MsgWithdrawRequest{}, "aether/vault/MsgWithdrawRequest", nil)
	cdc.RegisterConcrete(&MsgDrawRequest{}, "aether/vault/MsgDrawRequest", nil)
	cdc.RegisterConcrete(&MsgDepositAndDrawRequest{}, "aether/vault/MsgDepositAndDrawRequest", nil)
	cdc.RegisterConcrete(&MsgCreateStableMintRequest{}, "aether/vault/MsgCreateStableMintRequest", nil)
	cdc.RegisterConcrete(&MsgDepositStableMintRequest{}, "aether/vault/MsgDepositStableMintRequest", nil)
	cdc.RegisterConcrete(&MsgWithdrawStableMintRequest{}, "aether/vault/MsgWithdrawStableMintRequest", nil)
	cdc.RegisterConcrete(&MsgVaultInterestCalcRequest{}, "aether/vault/MsgVaultInterestCalcRequest", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateRequest{},
		&MsgDepositRequest{},
		&MsgWithdrawRequest{},
		&MsgDrawRequest{},
		&MsgRepayRequest{},
		&MsgCloseRequest{},
		&MsgDepositAndDrawRequest{},
		&MsgCreateStableMintRequest{},
		&MsgDepositStableMintRequest{},
		&MsgWithdrawStableMintRequest{},
		&MsgVaultInterestCalcRequest{},
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
