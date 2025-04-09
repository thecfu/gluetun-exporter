package types

type Statistics struct {
	Interface  string // Name of the Interface
	RxBytes    uint64 // Number of bytes received
	TxBytes    uint64 // Number of bytes transmitted
	RxPackets  uint64 // Number of packets received
	TxPackets  uint64 // Number of packets transmitted
	RxErrors   uint64 // Number of receive errors
	TxErrors   uint64 // Number of transmit errors
	RxDropped  uint64 // Number of packets dropped when receiving
	TxDropped  uint64 // Number of packets dropped when transmitting
	Collisions uint64 // Number of collisions in transmission
}
