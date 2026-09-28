// CreateSession creates a new public session.
func (a *API) CreateSession(id string) *Session {
	return NewSession(id)
}

// CreateTransport creates a new public transport.
func (a *API) CreateTransport(
	id string,
	t TransportType,
) *Transport {

	return NewTransport(id, t)
}

// CreateEvidence creates a public evidence record.
func (a *API) CreateEvidence(
	scenario string,
	verdict Verdict,
	message string,
) Evidence {

	return NewEvidence(
		scenario,
		verdict,
		message,
	)
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