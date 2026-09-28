package boundary

// Transport represents the public view of a transport.
//
// The Runtime Boundary intentionally exposes only observable
// transport information.
type Transport struct {
	ID string

	Type TransportType

	State TransportState
}

// TransportType identifies the transport technology.
type TransportType string

const (
	TransportUDP  TransportType = "UDP"
	TransportTCP  TransportType = "TCP"
	TransportQUIC TransportType = "QUIC"
)

// TransportState describes the observable transport state.
type TransportState string

const (
	TransportDetached TransportState = "DETACHED"
	TransportAttached TransportState = "ATTACHED"
	TransportActive   TransportState = "ACTIVE"
	TransportLost     TransportState = "LOST"
)

// NewTransport creates a public transport representation.
func NewTransport(id string, t TransportType) *Transport {

	return &Transport{
		ID:    id,
		Type:  t,
		State: TransportDetached,
	}
}

// Attach makes the transport available.
func (t *Transport) Attach() {

	t.State = TransportAttached
}

// Activate marks the transport as active.
func (t *Transport) Activate() {

	t.State = TransportActive
}

// Lose marks the transport as unavailable.
func (t *Transport) Lose() {

	t.State = TransportLost
}
