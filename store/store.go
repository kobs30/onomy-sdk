package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/onomyprotocol/onomy-sdk/store/cache"
	"github.com/onomyprotocol/onomy-sdk/store/rootmulti"
	"github.com/onomyprotocol/onomy-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
