package keeper

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/crypto-org-chain/cronos/x/icactl/types"
)

// DefaultGasCap defines the gas limit used to run internal evm call
const DefaultGasCap uint64 = 25000000

func (k Keeper) CallOnICAPacketResult(ctx sdk.Context, contract common.Address, channelID string, sequence uint64) ([]byte, error) {
	sequenceBigInt := new(big.Int).SetUint64(sequence)
	data, err := types.ModuleICAContract.ABI.Pack("onICAPacketResult", channelID, sequenceBigInt)
	if err != nil {
		return nil, err
	}
	_, res, err := k.CallEVM(ctx, &contract, data, big.NewInt(0))
	if err != nil {
		return nil, err
	}
	if res.Failed() {
		return nil, fmt.Errorf("failed calling contract onICAPacketResult method at address %s: %v", contract.Hex(), res.Ret)
	}
	return res.Ret, nil
}

func (k Keeper) CallOnICAPacketError(ctx sdk.Context, contract common.Address, channelID string, sequence uint64, packetErr string) ([]byte, error) {
	sequenceBigInt := new(big.Int).SetUint64(sequence)
	data, err := types.ModuleICAContract.ABI.Pack("onICAPacketError", channelID, sequenceBigInt, packetErr)
	if err != nil {
		return nil, err
	}
	_, res, err := k.CallEVM(ctx, &contract, data, big.NewInt(0))
	if err != nil {
		return nil, err
	}
	if res.Failed() {
		return nil, fmt.Errorf("failed calling contract onICAPacketError method at address %s: %v", contract.Hex(), res.Ret)
	}
	return res.Ret, nil
}

func (k Keeper) CallOnICAPacketTimeout(ctx sdk.Context, contract common.Address, channelID string, sequence uint64) ([]byte, error) {
	sequenceBigInt := new(big.Int).SetUint64(sequence)
	data, err := types.ModuleICAContract.ABI.Pack("onICAPacketTimeout", channelID, sequenceBigInt)
	if err != nil {
		return nil, err
	}
	_, res, err := k.CallEVM(ctx, &contract, data, big.NewInt(0))
	if err != nil {
		return nil, err
	}
	if res.Failed() {
		return nil, fmt.Errorf("failed calling contract onICAPacketTimeout method at address %s: %v", contract.Hex(), res.Ret)
	}
	return res.Ret, nil
}

// CallEVM execute an evm message from native module
func (k Keeper) CallEVM(ctx sdk.Context, to *common.Address, data []byte, value *big.Int) (*ethtypes.Message, *evmtypes.MsgEthereumTxResponse, error) {
	nonce := k.evmKeeper.GetNonce(ctx, types.EVMModuleAddress)
	msg := ethtypes.NewMessage(
		types.EVMModuleAddress,
		to,
		nonce,
		value, // amount
		DefaultGasCap,
		big.NewInt(0), nil, nil, // gasPrice
		data,
		nil,   // accessList
		false, // isFake
	)

	ret, err := k.evmKeeper.ApplyMessage(ctx, msg, nil, true)
	if err != nil {
		return nil, nil, err
	}
	return &msg, ret, nil
}
