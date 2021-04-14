package types

import (
	"github.com/onomyprotocol/onomy-sdk/codec"
	"github.com/onomyprotocol/onomy-sdk/codec/types"
	cryptocodec "github.com/onomyprotocol/onomy-sdk/crypto/codec"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	"github.com/onomyprotocol/onomy-sdk/types/msgservice"
	govtypes "github.com/onomyprotocol/onomy-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers the necessary x/distribution interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgWithdrawDelegatorReward{}, "onomy-sdk/MsgWithdrawDelegationReward", nil)
	cdc.RegisterConcrete(&MsgWithdrawValidatorCommission{}, "onomy-sdk/MsgWithdrawValidatorCommission", nil)
	cdc.RegisterConcrete(&MsgSetWithdrawAddress{}, "onomy-sdk/MsgModifyWithdrawAddress", nil)
	cdc.RegisterConcrete(&MsgFundCommunityPool{}, "onomy-sdk/MsgFundCommunityPool", nil)
	cdc.RegisterConcrete(&CommunityPoolSpendProposal{}, "onomy-sdk/CommunityPoolSpendProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgWithdrawDelegatorReward{},
		&MsgWithdrawValidatorCommission{},
		&MsgSetWithdrawAddress{},
		&MsgFundCommunityPool{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&CommunityPoolSpendProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/distribution module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/distribution and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
