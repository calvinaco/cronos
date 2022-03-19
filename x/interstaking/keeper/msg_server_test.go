package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/crypto-org-chain/cronos/x/interstaking/types"
    "github.com/crypto-org-chain/cronos/x/interstaking/keeper"
    keepertest "github.com/crypto-org-chain/cronos/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InterstakingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
