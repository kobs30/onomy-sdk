package legacybech32

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onomyprotocol/onomy-sdk/crypto/hd"
	"github.com/onomyprotocol/onomy-sdk/crypto/ledger"
	"github.com/onomyprotocol/onomy-sdk/testutil"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
)

func TestBeach32ifPbKey(t *testing.T) {
	require := require.New(t)
	path := *hd.NewFundraiserParams(0, sdk.CoinType, 0)
	priv, err := ledger.NewPrivKeySecp256k1Unsafe(path)
	require.Nil(err, "%s", err)
	require.NotNil(priv)

	pubKeyAddr, err := MarshalPubKey(AccPK, priv.PubKey())
	require.NoError(err)
	require.Equal("cosmospub1addwnpepqd87l8xhcnrrtzxnkql7k55ph8fr9jarf4hn6udwukfprlalu8lgw0urza0",
		pubKeyAddr, "Is your device using test mnemonic: %s ?", testutil.TestMnemonic)

}
