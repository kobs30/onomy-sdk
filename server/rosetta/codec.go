package rosetta

import (
	"github.com/onomyprotocol/onomy-sdk/codec"
	codectypes "github.com/onomyprotocol/onomy-sdk/codec/types"
	cryptocodec "github.com/onomyprotocol/onomy-sdk/crypto/codec"
	authcodec "github.com/onomyprotocol/onomy-sdk/x/auth/types"
	bankcodec "github.com/onomyprotocol/onomy-sdk/x/bank/types"
)

// MakeCodec generates the codec required to interact
// with the cosmos APIs used by the rosetta gateway
func MakeCodec() (*codec.ProtoCodec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	authcodec.RegisterInterfaces(ir)
	bankcodec.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)

	return cdc, ir
}
