package keeper

import (
	"github.com/cosbtc/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}
