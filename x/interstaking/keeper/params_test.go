package keeper_test

import (
	"testing"

	testkeeper "github.com/crypto-org-chain/cronos/testutil/keeper"
	"github.com/crypto-org-chain/cronos/x/interstaking/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InterstakingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
