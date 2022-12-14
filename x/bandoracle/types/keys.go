package types

const (
	// ModuleName defines the module name.
	ModuleName = "bandoracleV1"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_bandoracle"

	// Version defines the current version the IBC module supports.
	Version = "bandchain-1"

	// PortID is the default port id that module binds to.
	PortID = "bandoracleV1"
)

var (
	// PortKey defines the key to store the port ID in store.
	PortKey        = KeyPrefix("bandoracle-port-")
	MsgDataKey     = []byte{0x02}
	CheckFlagKey   = []byte{0x03}
	DiscardFlagKey = []byte{0x04}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
