package interstaking_test

import (
	"testing"

	keepertest "github.com/crypto-org-chain/cronos/testutil/keeper"
	"github.com/crypto-org-chain/cronos/testutil/nullify"
	"github.com/crypto-org-chain/cronos/x/interstaking"
	"github.com/crypto-org-chain/cronos/x/interstaking/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InterstakingKeeper(t)
	interstaking.InitGenesis(ctx, *k, genesisState)
	got := interstaking.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
