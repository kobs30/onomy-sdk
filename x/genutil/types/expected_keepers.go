package types

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/onomyprotocol/onomy-sdk/codec"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	auth "github.com/onomyprotocol/onomy-sdk/x/auth/types"
	bankexported "github.com/onomyprotocol/onomy-sdk/x/bank/exported"
)

// StakingKeeper defines the expected staking keeper (noalias)
type StakingKeeper interface {
	ApplyAndReturnValidatorSetUpdates(sdk.Context) (updates []abci.ValidatorUpdate, err error)
}

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	NewAccount(sdk.Context, auth.AccountI) auth.AccountI
	SetAccount(sdk.Context, auth.AccountI)
	IterateAccounts(ctx sdk.Context, process func(auth.AccountI) (stop bool))
}

// GenesisAccountsIterator defines the expected iterating genesis accounts object (noalias)
type GenesisAccountsIterator interface {
	IterateGenesisAccounts(
		cdc *codec.LegacyAmino,
		appGenesis map[string]json.RawMessage,
		cb func(auth.AccountI) (stop bool),
	)
}

// GenesisAccountsIterator defines the expected iterating genesis accounts object (noalias)
type GenesisBalancesIterator interface {
	IterateGenesisBalances(
		cdc codec.JSONMarshaler,
		appGenesis map[string]json.RawMessage,
		cb func(bankexported.GenesisBalance) (stop bool),
	)
}
