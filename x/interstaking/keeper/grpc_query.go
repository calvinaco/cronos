package keeper

import (
	"github.com/crypto-org-chain/cronos/x/interstaking/types"
)

var _ types.QueryServer = Keeper{}
