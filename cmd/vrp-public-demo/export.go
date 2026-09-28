package main

import (
	"github.com/Endless33/vrp-public-demo/internal/evidence"
)

func exportTransportMigration() error {

	report := evidence.Report{
		Version:   "v0.1.0",
		Scenario:  "Transport Migration",
		Verdict:   "PASS",
		Principle: "SESSION ≠ TRANSPORT",
		SessionID: "demo-session",
		Transport: "udp:B",
	}

	return evidence.Export(
		"evidence/transport-migration.json",
		report,
	)
}
