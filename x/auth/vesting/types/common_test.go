package types_test

import (
	"github.com/onomyprotocol/onomy-sdk/simapp"
)

var (
	app      = simapp.Setup(false)
	appCodec = simapp.MakeTestEncodingConfig().Marshaler
)
