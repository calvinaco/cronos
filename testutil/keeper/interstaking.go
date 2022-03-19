package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	"github.com/crypto-org-chain/cronos/x/interstaking/keeper"
	"github.com/crypto-org-chain/cronos/x/interstaking/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func InterstakingKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	portKey := types.PortKey
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)

	ss := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"InterstakingSubSpace",
	)
	IBCKeeper := ibckeeper.NewKeeper(
		appCodec,
		storeKey,
		ss,
		nil,
		nil,
		capabilityKeeper.ScopeToModule("InterstakingIBCKeeper"),
	)

	paramsSubspace := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"InterstakingParams",
	)

	icaStoreKey := sdk.NewKVStoreKey(icacontrollertypes.StoreKey)
	icaMemStoreKey := storetypes.NewMemoryStoreKey(icacontrollertypes.SubModuleName)
	icaSubSpace := typesparams.NewSubspace(appCodec,
		types.Amino,
		icaStoreKey,
		icaMemStoreKey,
		"InterstakingSubSpace",
	)
	icaControllerKeeper := icacontrollerkeeper.NewKeeper(
		appCodec, icaStoreKey, icaSubSpace,
		IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		IBCKeeper.ChannelKeeper, &IBCKeeper.PortKeeper,
		capabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName), baseapp.NewMsgServiceRouter(),
	)
	k := keeper.NewKeeper(
		appCodec,
		portKey,
		storeKey,
		paramsSubspace,
		capabilityKeeper.ScopeToModule("InterstakingScopedKeeper"),
		IBCKeeper.ChannelKeeper,
		&IBCKeeper.PortKeeper,
		icaControllerKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
