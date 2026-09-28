package boundary

// Version identifies the public Runtime Boundary API version.
const Version = "v0.1.0"

// API exposes the public Runtime Boundary.
//
// The Runtime Boundary intentionally exposes observable protocol
// behavior only. Protected runtime implementation remains private.
type API struct{}

// New creates a new public Runtime Boundary API.
func New() *API {
	return &API{}
}

// Name returns the public boundary name.
func (a *API) Name() string {
	return "VRP Runtime Boundary"
}

// Version returns the current public API version.
func (a *API) Version() string {
	return Version
}

// DesignPrinciple returns the architectural principle demonstrated
// by the public boundary.
func (a *API) DesignPrinciple() string {
	return "SESSION ≠ TRANSPORT"
}

// SwitchTransport performs a public transport replacement.
func (a *API) SwitchTransport(
	session *Session,
	transport *Transport,
) {

	if session == nil || transport == nil {
		return
	}

	transport.Attach()
	transport.Activate()

	session.SwitchTransport(transport.ID)
}