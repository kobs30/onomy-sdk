package main

import (
	"os"

	"github.com/onomyprotocol/onomy-sdk/server"
	svrcmd "github.com/onomyprotocol/onomy-sdk/server/cmd"
	"github.com/onomyprotocol/onomy-sdk/simapp"
	"github.com/onomyprotocol/onomy-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
