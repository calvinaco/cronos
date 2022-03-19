package types

const (
	// ModuleName defines the module name
	ModuleName = "interstaking"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_interstaking"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("interstaking-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
