package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onomyprotocol/onomy-sdk/crypto/keys/ed25519"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	bank "github.com/onomyprotocol/onomy-sdk/x/bank/types"
)

var granter = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
var grantee = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
var msgType = bank.SendAuthorization{}.MethodName()

func TestGrantkey(t *testing.T) {
	granter1, grantee1 := ExtractAddressesFromGrantKey(GetAuthorizationStoreKey(grantee, granter, msgType))
	require.Equal(t, granter, granter1)
	require.Equal(t, grantee, grantee1)
}
