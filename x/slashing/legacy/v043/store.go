package v043

import (
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	v043distribution "github.com/onomyprotocol/onomy-sdk/x/distribution/legacy/v043"
	v040slashing "github.com/onomyprotocol/onomy-sdk/x/slashing/legacy/v040"
)

// MigrateStore performs in-place store migrations from v0.40 to v0.42. The
// migration includes:
//
// - Change addresses to be length-prefixed.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey) error {
	store := ctx.KVStore(storeKey)
	v043distribution.MigratePrefixAddress(store, v040slashing.ValidatorSigningInfoKeyPrefix)
	v043distribution.MigratePrefixAddressBytes(store, v040slashing.ValidatorMissedBlockBitArrayKeyPrefix)
	v043distribution.MigratePrefixAddress(store, v040slashing.AddrPubkeyRelationKeyPrefix)

	return nil
}
