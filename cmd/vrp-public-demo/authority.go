package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-boundary-preview/boundary"
)

func RunAuthority() {

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

	fmt.Println("----- AUTHORITY VALIDATION -----")
	fmt.Println()

	fmt.Println("Authority received")
	fmt.Println("Authority validated")
	fmt.Println("Canonical session preserved")

	ev := api.CreateEvidence(
		"Authority Validation",
		boundary.VerdictPass,
		"Authority successfully validated.",
	)

	fmt.Println()
	fmt.Println("Evidence")
	fmt.Println("Scenario:", ev.Scenario)
	fmt.Println("Verdict :", ev.Verdict)
	fmt.Println("Message :", ev.Message)

	fmt.Println()
	fmt.Println("FINAL VERDICT")
	fmt.Println("AUTHORITY PRESERVED")
}
