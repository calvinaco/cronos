package types

import (
	"bytes"
	"encoding/binary"
)

const (
	// ModuleName defines the module name
	ModuleName = "icactl"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_icactl"

	// Version defines the current version the IBC module supports
	Version = "icactl-1"
)

// prefix bytes for the cronos persistent store
const (
	prefixPacketIDToContract = iota + 1
)

// KVStore key prefixes
var (
	KeyPrefixPacketIDToContract = []byte{prefixPacketIDToContract}
)

// PacketIDToContractKey defines the store key for packet to contract mapping
func PacketIDToContractKey(connectionID, portID string, sequence uint64) []byte {
	sequenceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(sequenceBytes, sequence)

	// "KeyPrefixPacketToContract{connectionID}:{portID}:{sequence(BigEndian)}
	// Use colon(`:`) as separator because it is not a valid character in IBC Identifier
	// Ref: https://github.com/cosmos/ibc/tree/master/spec/core/ics-024-host-requirements
	return bytes.Join([][]byte{
		KeyPrefixPacketIDToContract,
		{':'},
		[]byte(connectionID),
		{':'},
		[]byte(portID),
		{':'},
		sequenceBytes,
	}, []byte{})
}
