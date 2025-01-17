package simulation

import (
	"bytes"
	"fmt"

	"github.com/onomyprotocol/onomy-sdk/codec"
	"github.com/onomyprotocol/onomy-sdk/types/kv"
	"github.com/onomyprotocol/onomy-sdk/x/feegrant/types"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding feegrant type.
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.FeeAllowanceKeyPrefix):
			var grantA, grantB types.FeeAllowanceGrant
			cdc.MustUnmarshalBinaryBare(kvA.Value, &grantA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &grantB)
			return fmt.Sprintf("%v\n%v", grantA, grantB)
		default:
			panic(fmt.Sprintf("invalid feegrant key %X", kvA.Key))
		}
	}
}
