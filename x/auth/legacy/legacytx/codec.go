package legacytx

import (
	"github.com/onomyprotocol/onomy-sdk/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "onomy-sdk/StdTx", nil)
}
