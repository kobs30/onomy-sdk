package v043_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onomyprotocol/onomy-sdk/simapp"
	"github.com/onomyprotocol/onomy-sdk/store/prefix"
	"github.com/onomyprotocol/onomy-sdk/testutil"
	"github.com/onomyprotocol/onomy-sdk/testutil/testdata"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	v040bank "github.com/onomyprotocol/onomy-sdk/x/bank/legacy/v040"
	v043bank "github.com/onomyprotocol/onomy-sdk/x/bank/legacy/v043"
	"github.com/onomyprotocol/onomy-sdk/x/bank/types"
)

func TestSupplyMigration(t *testing.T) {
	encCfg := simapp.MakeTestEncodingConfig()
	bankKey := sdk.NewKVStoreKey("bank")
	ctx := testutil.DefaultContext(bankKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(bankKey)

	oldFooCoin := sdk.NewCoin("foo", sdk.NewInt(100))
	oldBarCoin := sdk.NewCoin("bar", sdk.NewInt(200))

	// Old supply was stored as a single blob under the `SupplyKey`.
	var oldSupply v040bank.SupplyI
	oldSupply = &types.Supply{Total: sdk.NewCoins(oldFooCoin, oldBarCoin)}
	oldSupplyBz, err := encCfg.Marshaler.MarshalInterface(oldSupply)
	require.NoError(t, err)
	store.Set(v040bank.SupplyKey, oldSupplyBz)

	// Run migration.
	err = v043bank.MigrateStore(ctx, bankKey, encCfg.Marshaler)
	require.NoError(t, err)

	// New supply is indexed by denom.
	var newFooCoin, newBarCoin sdk.Coin
	supplyStore := prefix.NewStore(store, types.SupplyKey)
	encCfg.Marshaler.MustUnmarshalBinaryBare(supplyStore.Get([]byte("foo")), &newFooCoin)
	encCfg.Marshaler.MustUnmarshalBinaryBare(supplyStore.Get([]byte("bar")), &newBarCoin)

	require.Equal(t, oldFooCoin, newFooCoin)
	require.Equal(t, oldBarCoin, newBarCoin)
}

func TestBalanceKeysMigration(t *testing.T) {
	encCfg := simapp.MakeTestEncodingConfig()
	bankKey := sdk.NewKVStoreKey("bank")
	ctx := testutil.DefaultContext(bankKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(bankKey)

	_, _, addr := testdata.KeyTestPubAddr()
	denom := []byte("foo")
	value := []byte("bar")

	oldKey := append(append(v040bank.BalancesPrefix, addr...), denom...)
	store.Set(oldKey, value)

	err := v043bank.MigrateStore(ctx, bankKey, encCfg.Marshaler)
	require.NoError(t, err)

	newKey := append(types.CreateAccountBalancesPrefix(addr), denom...)
	// -7 because we replaced "balances" with 0x02,
	// +1 because we added length-prefix to address.
	require.Equal(t, len(oldKey)-7+1, len(newKey))
	require.Nil(t, store.Get(oldKey))
	require.Equal(t, value, store.Get(newKey))
}
