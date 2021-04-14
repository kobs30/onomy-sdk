package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onomyprotocol/onomy-sdk/codec"
	codectypes "github.com/onomyprotocol/onomy-sdk/codec/types"
	"github.com/onomyprotocol/onomy-sdk/types/module"
	simtypes "github.com/onomyprotocol/onomy-sdk/types/simulation"
	"github.com/onomyprotocol/onomy-sdk/x/authz/simulation"
	"github.com/onomyprotocol/onomy-sdk/x/authz/types"
)

func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var authzGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &authzGenesis)

	require.Len(t, authzGenesis.Authorization, 0)
}
