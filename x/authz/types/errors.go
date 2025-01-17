package types

import (
	sdkerrors "github.com/onomyprotocol/onomy-sdk/types/errors"
)

// x/authz module sentinel errors
var (
	ErrInvalidExpirationTime = sdkerrors.Register(ModuleName, 3, "expiration time of authorization should be more than current time")
)
