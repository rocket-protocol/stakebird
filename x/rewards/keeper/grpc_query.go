package keeper

import (
	"github.com/public-awesome/stargaze/x/rewards/types"
)

var _ types.QueryServer = Keeper{}
