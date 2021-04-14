package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/onomyprotocol/onomy-sdk/codec"
	cryptoAmino "github.com/onomyprotocol/onomy-sdk/crypto/codec"
	"github.com/onomyprotocol/onomy-sdk/testutil/testdata"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	"github.com/onomyprotocol/onomy-sdk/x/auth/legacy/legacytx"
	"github.com/onomyprotocol/onomy-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "onomy-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
