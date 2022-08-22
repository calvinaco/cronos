package icactl

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	"github.com/evmos/ethermint/x/evm/keeper/precompiles"

	"github.com/crypto-org-chain/cronos/x/icactl/keeper"
)

var _ precompiles.ICAModule = &ICAModule{}

type ICAModule struct {
	keeper keeper.Keeper
}

// ICAModule creates a new ICAModule given the keeper
func NewICAModule(k keeper.Keeper) ICAModule {
	return ICAModule{
		keeper: k,
	}
}

func (im *ICAModule) OnRegisterInterchainAccount(
	ctx sdk.Context,
	precompileCtx precompiles.ModuleContext,
	connectionID string,
	owner string,
) error {
	fmt.Println("OnRegisterInterchainAccount", connectionID, owner, precompileCtx.Caller)
	return nil
}

func (im *ICAModule) OnSendTx(
	ctx sdk.Context,
	precompileCtx precompiles.ModuleContext,
	chanCap *capabilitytypes.Capability,
	connectionID string,
	channelID string,
	portID string,
	icaPacketData icatypes.InterchainAccountPacketData,
	timeoutTimestamp uint64,
	packetSequence uint64,
) error {
	fmt.Println("OnSendTx", connectionID, channelID, portID, icaPacketData, timeoutTimestamp, packetSequence, precompileCtx.Caller)
	im.keeper.SetContractByPacketID(
		ctx, channelID, portID, packetSequence,
		precompileCtx.Caller,
	)
	return nil
}
