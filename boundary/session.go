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

<<<<<<< HEAD
	session.SwitchTransport(transport.ID)
}
=======
	s.State = SessionActive
}

// BeginRecovery marks the session as recovering.
func (s *Session) BeginRecovery() {

	s.State = SessionRecovering
}

// AttachTransport updates the active transport.
func (s *Session) AttachTransport(name string) {

	s.ActiveTransport = name
}

// Close marks the session as closed.
func (s *Session) Close() {

	s.State = SessionClosed
}
>>>>>>> 82b06d5 (Demonstrate runtime boundary transport migration)
