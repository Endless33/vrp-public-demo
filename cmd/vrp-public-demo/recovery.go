package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-boundary-preview/boundary"
)

func RunRecovery() {

	api := boundary.New()

	session := api.CreateSession("demo-session")

	primary := api.CreateTransport(
		"udp:A",
		boundary.TransportUDP,
	)

	backup := api.CreateTransport(
		"udp:B",
		boundary.TransportUDP,
	)

	api.SwitchTransport(session, primary)
	session.Activate()

	fmt.Println("======================================")
	fmt.Println(api.Name())
	fmt.Println("======================================")
	fmt.Println()

	fmt.Println("Version:", api.Version())
	fmt.Println("Principle:", api.DesignPrinciple())
	fmt.Println()

	fmt.Println("----- SESSION RECOVERY -----")
	fmt.Println()

	fmt.Println("Transport lost")

	session.BeginRecovery()

	fmt.Println("Recovery started")

	api.SwitchTransport(session, backup)

	session.Activate()

	fmt.Println("Replacement transport attached")
	fmt.Println("Recovered transport:", session.ActiveTransport)
	fmt.Println("Session state:", session.State)

	ev := api.CreateEvidence(
		"Session Recovery",
		boundary.VerdictPass,
		"Logical session successfully recovered.",
	)

	fmt.Println()
	fmt.Println("Evidence")
	fmt.Println("Scenario:", ev.Scenario)
	fmt.Println("Verdict :", ev.Verdict)
	fmt.Println("Message :", ev.Message)

	fmt.Println()
	fmt.Println("FINAL VERDICT")
	fmt.Println("RECOVERY PRESERVED")
}
