package types

import (
	"github.com/onomyprotocol/onomy-sdk/codec"
	cryptocodec "github.com/onomyprotocol/onomy-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
