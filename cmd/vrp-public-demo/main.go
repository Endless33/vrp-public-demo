package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-boundary-preview/boundary"
)

func main() {

	api := boundary.New()

	session := api.CreateSession("demo-session")

	transport := api.CreateTransport(
		"udp:A",
		boundary.TransportUDP,
	)

	transport.Attach()
	transport.Activate()

	session.AttachTransport(transport.ID)
	session.Activate()

	evidence := api.CreateEvidence(
		"Session Establishment",
		boundary.VerdictPass,
		"Logical session successfully established.",
	)

	fmt.Println("======================================")
	fmt.Println(api.Name())
	fmt.Println("======================================")
	fmt.Println()

	fmt.Println("Version:", api.Version())
	fmt.Println("Principle:", api.DesignPrinciple())
	fmt.Println()

	fmt.Println("Session ID:", session.ID)
	fmt.Println("Session State:", session.State)
	fmt.Println("Active Transport:", session.ActiveTransport)
	fmt.Println()

	fmt.Println("Evidence")
	fmt.Println("Scenario:", evidence.Scenario)
	fmt.Println("Verdict:", evidence.Verdict)
	fmt.Println("Message:", evidence.Message)
	fmt.Println()

	fmt.Println("PUBLIC RUNTIME BOUNDARY INITIALIZED")
}
