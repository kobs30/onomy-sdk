package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/onomyprotocol/onomy-sdk/codec"
	codectypes "github.com/onomyprotocol/onomy-sdk/codec/types"
	"github.com/onomyprotocol/onomy-sdk/std"
	"github.com/onomyprotocol/onomy-sdk/testutil/testdata"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	"github.com/onomyprotocol/onomy-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
