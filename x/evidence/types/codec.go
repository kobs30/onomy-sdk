package types

import (
	"github.com/onomyprotocol/onomy-sdk/codec"
	"github.com/onomyprotocol/onomy-sdk/codec/types"
	cryptocodec "github.com/onomyprotocol/onomy-sdk/crypto/codec"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	"github.com/onomyprotocol/onomy-sdk/types/msgservice"
	"github.com/onomyprotocol/onomy-sdk/x/evidence/exported"
)

// RegisterLegacyAminoCodec registers all the necessary types and interfaces for the
// evidence module.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.Evidence)(nil), nil)
	cdc.RegisterConcrete(&MsgSubmitEvidence{}, "onomy-sdk/MsgSubmitEvidence", nil)
	cdc.RegisterConcrete(&Equivocation{}, "onomy-sdk/Equivocation", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSubmitEvidence{})
	registry.RegisterInterface(
		"cosmos.evidence.v1beta1.Evidence",
		(*exported.Evidence)(nil),
		&Equivocation{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/evidence module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/evidence and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
