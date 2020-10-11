package types

import (
	sdkerrors "github.com/ivansukach/modified-cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrInvalid = sdkerrors.Register(ModuleName, 1, "custom error message")
)
