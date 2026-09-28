package boundary

import "time"

// Verdict represents the observable outcome of a public
// runtime verification scenario.
type Verdict string

const (
	VerdictPass Verdict = "PASS"
	VerdictFail Verdict = "FAIL"
)

// Evidence represents the public verification result.
//
// The Runtime Boundary intentionally exports only observable
// engineering evidence. Internal runtime implementation,
// canonical state, cryptographic material, and protected
// algorithms are never exposed.
type Evidence struct {

	Scenario string

	Verdict Verdict

	Message string

	Timestamp time.Time
}

// NewEvidence creates a public evidence record.
func NewEvidence(
	scenario string,
	verdict Verdict,
	message string,
) Evidence {

	return Evidence{
		Scenario:  scenario,
		Verdict:   verdict,
		Message:   message,
		Timestamp: time.Now().UTC(),
	}
}