package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSendCreatePool{}, "amm/SendCreatePool", nil)
	cdc.RegisterConcrete(&MsgSendAddLiquidity{}, "amm/SendAddLiquidity", nil)
	cdc.RegisterConcrete(&MsgSendRemoveLiquidity{}, "amm/SendRemoveLiquidity", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendAddLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendRemoveLiquidity{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
