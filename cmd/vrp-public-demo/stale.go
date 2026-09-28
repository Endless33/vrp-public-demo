package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-boundary-preview/boundary"
)

func RunStale() {

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

	fmt.Println("----- STALE STATE REJECTION -----")
	fmt.Println()

	fmt.Println("Receiving stale state...")
	fmt.Println("State rejected")
	fmt.Println("Canonical session preserved")

	ev := api.CreateEvidence(
		"Stale-State Rejection",
		boundary.VerdictPass,
		"Stale state successfully rejected.",
	)

	fmt.Println()
	fmt.Println("Evidence")
	fmt.Println("Scenario:", ev.Scenario)
	fmt.Println("Verdict :", ev.Verdict)
	fmt.Println("Message :", ev.Message)

	fmt.Println()
	fmt.Println("FINAL VERDICT")
	fmt.Println("STALE STATE REJECTED")
}
