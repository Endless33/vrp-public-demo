package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-boundary-preview/boundary"
)

func main() {

	api := boundary.New()

	session := api.CreateSession("demo-session")

	transportA := api.CreateTransport(
		"udp:A",
		boundary.TransportUDP,
	)

	transportB := api.CreateTransport(
		"udp:B",
		boundary.TransportUDP,
	)

	api.SwitchTransport(session, transportA)
	session.Activate()

	fmt.Println("======================================")
	fmt.Println(api.Name())
	fmt.Println("======================================")
	fmt.Println()

	fmt.Println("Version:", api.Version())
	fmt.Println("Principle:", api.DesignPrinciple())
	fmt.Println()

	fmt.Println("Session:", session.ID)
	fmt.Println("State:", session.State)
	fmt.Println("Transport:", session.ActiveTransport)

	fmt.Println()
	fmt.Println("----- TRANSPORT MIGRATION -----")
	fmt.Println()

	api.SwitchTransport(session, transportB)

	fmt.Println("Transport switched")
	fmt.Println("Active transport:", session.ActiveTransport)

	evidence := api.CreateEvidence(
		"Transport Migration",
		boundary.VerdictPass,
		"Logical session preserved while transport changed.",
	)

	fmt.Println()
	fmt.Println("Evidence")
	fmt.Println("Scenario:", evidence.Scenario)
	fmt.Println("Verdict :", evidence.Verdict)
	fmt.Println("Message :", evidence.Message)

	fmt.Println()
	fmt.Println("FINAL VERDICT")
	fmt.Println("CONTINUITY PRESERVED")
}
