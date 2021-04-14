package testutil

import (
	"github.com/onomyprotocol/onomy-sdk/testutil"
	clitestutil "github.com/onomyprotocol/onomy-sdk/testutil/cli"
	"github.com/onomyprotocol/onomy-sdk/testutil/network"
	"github.com/onomyprotocol/onomy-sdk/x/authz/client/cli"
)

func ExecGrantAuthorization(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
