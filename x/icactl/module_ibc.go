package icactl

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	proto "github.com/gogo/protobuf/proto"

	"github.com/crypto-org-chain/cronos/x/icactl/keeper"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule implements the ICS26 interface for interchain accounts controller chains
type IBCModule struct {
	keeper keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (am IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	// Claim channel capability passed back by IBC module
	if err := am.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return err
	}

	return nil
}

// OnChanOpenTry implements the IBCModule interface
func (am IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	// https://github.com/cosmos/ibc-go/blob/v3.0.0/docs/apps/interchain-accounts/auth-modules.md#ibcmodule-implementation
	return "", sdkerrors.Wrap(icatypes.ErrInvalidChannelFlow, "channel handshake must be initiated by controller chain")
}

// OnChanOpenAck implements the IBCModule interface
func (am IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (am IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// https://github.com/cosmos/ibc-go/blob/v3.0.0/docs/apps/interchain-accounts/auth-modules.md#ibcmodule-implementation
	return sdkerrors.Wrap(icatypes.ErrInvalidChannelFlow, "channel handshake must be initiated by controller chain")
}

// OnChanCloseInit implements the IBCModule interface
func (am IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// https://github.com/cosmos/ibc-go/blob/v3.0.0/docs/apps/interchain-accounts/auth-modules.md#ibcmodule-implementation
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (am IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (am IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	// https://github.com/cosmos/ibc-go/blob/v3.0.0/docs/apps/interchain-accounts/auth-modules.md#ibcmodule-implementation
	return channeltypes.NewErrorAcknowledgement("cannot receive packet on controller chain")
}

// OnAcknowledgementPacket implements the IBCModule interface
func (am IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	fmt.Printf("received ICA packet %s:%s:%d acknowledgement: %s\n", packet.SourceChannel, packet.SourcePort, packet.Sequence, acknowledgement)

	var ack channeltypes.Acknowledgement
	if err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-27 packet acknowledgement: %v", err)
	}

	contract, found := am.keeper.GetContractByPacketID(ctx, packet.SourceChannel, packet.SourcePort, packet.Sequence)
	if !found {
		fmt.Printf("ICA packet %s:%s:%d contract origin not found\n", packet.SourceChannel, packet.SourcePort, packet.Sequence)
		return nil
	}
	fmt.Printf("ICA packet %s:%s:%d contract origin found: %s\n", packet.SourceChannel, packet.SourcePort, packet.Sequence, contract)

	if ackError, ok := ack.GetResponse().(*channeltypes.Acknowledgement_Error); ok {
		fmt.Printf("received ICA packet error acknowledgement: %s\n", ackError)
		am.keeper.CallOnICAPacketError(ctx, contract, packet.SourceChannel, packet.Sequence, ackError.Error)
		return nil
	}

	txMsgData := &sdk.TxMsgData{}
	if err := proto.Unmarshal(ack.GetResult(), txMsgData); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-27 tx message data: %v", err)
	}

	fmt.Println("received ICA packet result acknowledgement", txMsgData)
	if _, err := am.keeper.CallOnICAPacketResult(ctx, contract, packet.SourceChannel, packet.Sequence); err != nil {
		fmt.Printf("cannot call contract onICAPacketResult: %s\n", err)
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot call contract onICAPacketResult: %v", err)
	}
	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (am IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	fmt.Printf("received ICA packet timeout %s:%s:%d\n", packet.SourceChannel, packet.SourcePort, packet.Sequence)
	contract, found := am.keeper.GetContractByPacketID(ctx, packet.SourceChannel, packet.SourcePort, packet.Sequence)
	if !found {
		fmt.Printf("ICA packet %s:%s:%d contract origin not found\n", packet.SourceChannel, packet.SourcePort, packet.Sequence)
		return nil
	}
	fmt.Printf("ICA packet %s:%s:%d contract origin found: %s\n", packet.SourceChannel, packet.SourcePort, packet.Sequence, contract)

	if _, err := am.keeper.CallOnICAPacketTimeout(ctx, contract, packet.SourceChannel, packet.Sequence); err != nil {
		fmt.Printf("cannot call contract onICAPacketTimeout: %s\n", err)
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot call contract onICAPacketTimeout: %v", err)
	}
	return nil
}

func (am IBCModule) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (version string, err error) {
	return proposedVersion, nil
}
