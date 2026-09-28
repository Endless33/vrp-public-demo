package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-boundary-preview/boundary"
)

func RunSession() {

	api := boundary.New()

	session := api.CreateSession("demo-session")
	session.Activate()

	fmt.Println("======================================")
	fmt.Println(api.Name())
	fmt.Println("======================================")
	fmt.Println()

	fmt.Println("Version:", api.Version())
	fmt.Println("Principle:", api.DesignPrinciple())
	fmt.Println()

	fmt.Println("----- SESSION ESTABLISHMENT -----")
	fmt.Println()

	fmt.Println("Session created")
	fmt.Println("Session ID:", session.ID)
	fmt.Println("Session state:", session.State)

	ev := api.CreateEvidence(
		"Session Establishment",
		boundary.VerdictPass,
		"Logical session successfully established.",
	)

	fmt.Println()
	fmt.Println("Evidence")
	fmt.Println("Scenario:", ev.Scenario)
	fmt.Println("Verdict :", ev.Verdict)
	fmt.Println("Message :", ev.Message)

	fmt.Println()
	fmt.Println("FINAL VERDICT")
	fmt.Println("SESSION ESTABLISHED")
}
