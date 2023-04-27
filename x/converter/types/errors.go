package types

import errorsmod "cosmossdk.io/errors"

// errors
var (
	ErrABIPack   = errorsmod.Register(ModuleName, 1, "contract ABI pack failed")
	ErrABIUnpack = errorsmod.Register(ModuleName, 2, "contract ABI unpack failed")
)
