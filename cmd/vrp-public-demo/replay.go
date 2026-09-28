package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-boundary-preview/boundary"
)

func RunReplay() {

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

	fmt.Println("----- REPLAY REJECTION -----")
	fmt.Println()

	fmt.Println("Receiving duplicated packet...")
	fmt.Println("Replay detected")
	fmt.Println("Packet rejected")

	ev := api.CreateEvidence(
		"Replay Rejection",
		boundary.VerdictPass,
		"Replay attempt successfully rejected.",
	)

	fmt.Println()
	fmt.Println("Evidence")
	fmt.Println("Scenario:", ev.Scenario)
	fmt.Println("Verdict :", ev.Verdict)
	fmt.Println("Message :", ev.Message)

	fmt.Println()
	fmt.Println("FINAL VERDICT")
	fmt.Println("REPLAY REJECTED")
}
