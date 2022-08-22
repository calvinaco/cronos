package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channelkeeper "github.com/cosmos/ibc-go/v3/modules/core/04-channel/keeper"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/crypto-org-chain/cronos/x/icactl/types"
	"github.com/ethereum/go-ethereum/common"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		paramStore paramtypes.Subspace

		channelKeeper       channelkeeper.Keeper
		icaControllerKeeper icacontrollerkeeper.Keeper
		scopedKeeper        capabilitykeeper.ScopedKeeper
		evmKeeper           *evmkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	paramStore paramtypes.Subspace,
	channelKeeper channelkeeper.Keeper,
	icaControllerKeeper icacontrollerkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	evmKeeper *evmkeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !paramStore.HasKeyTable() {
		paramStore = paramStore.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramStore: paramStore,

		channelKeeper:       channelKeeper,
		icaControllerKeeper: icaControllerKeeper,
		scopedKeeper:        scopedKeeper,
		evmKeeper:           evmKeeper,
	}
}

// GetContractByPacketID find the corresponding contract for the packet identity.
func (k *Keeper) GetContractByPacketID(ctx sdk.Context, channelID, portID string, sequence uint64) (common.Address, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PacketIDToContractKey(channelID, portID, sequence))
	if len(bz) == 0 {
		return common.Address{}, false
	}

	return common.BytesToAddress(bz), true
}

// GetContractByPacketID find the corresponding contract for the packet identity.
func (k *Keeper) SetContractByPacketID(ctx sdk.Context, channelID, portID string, sequence uint64, contract common.Address) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PacketIDToContractKey(channelID, portID, sequence), contract.Bytes())
}

// DeleteContractByPacketID deletes the kvpair of the packet identity.
func (k *Keeper) DeleteContractByPacketID(ctx sdk.Context, channelID, portID string, sequence uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PacketIDToContractKey(channelID, portID, sequence))
}

// DoSubmitTx submits a transaction to the host chain on behalf of interchain account
func (k *Keeper) DoSubmitTx(ctx sdk.Context, connectionID, owner string, msgs []sdk.Msg, timeoutDuration time.Duration) error {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelID, found := k.icaControllerKeeper.GetActiveChannelID(ctx, connectionID, portID)
	if !found {
		return sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	channelCapability, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, msgs)
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutDuration should be constraited by MinTimeoutDuration parameter.
	timeoutTimestamp := ctx.BlockTime().Add(timeoutDuration).UnixNano()

	_, err = k.icaControllerKeeper.SendTx(ctx, channelCapability, connectionID, portID, packetData, uint64(timeoutTimestamp))
	if err != nil {
		return err
	}

	return nil
}

// RegisterInterchainAccount registers an interchain account with the given `connectionId` and `owner` on the host chain
func (k *Keeper) RegisterInterchainAccount(ctx sdk.Context, connectionID, owner string) error {
	return k.icaControllerKeeper.RegisterInterchainAccount(ctx, connectionID, owner)
}

// GetInterchainAccountAddress fetches the interchain account address for given `connectionId` and `owner`
func (k *Keeper) GetInterchainAccountAddress(ctx sdk.Context, connectionID, owner string) (string, error) {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return "", status.Errorf(codes.InvalidArgument, "invalid owner address: %s", err)
	}

	icaAddress, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, connectionID, portID)

	if !found {
		return "", status.Errorf(codes.NotFound, "could not find account")
	}

	return icaAddress, nil
}

func (k *Keeper) GetChannelConnection(ctx sdk.Context, portID string, channelID string) (string, error) {
	connectionID, _, err := k.channelKeeper.GetChannelConnection(ctx, portID, channelID)
	if err != nil {
		return "", err
	}

	return connectionID, nil
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
