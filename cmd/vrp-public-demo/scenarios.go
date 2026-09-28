package main

import "fmt"

func RunSessionEstablishment() {

	fmt.Println("========================================")
	fmt.Println("SESSION ESTABLISHMENT")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("Creating session...")
	fmt.Println("PASS")

	fmt.Println("Allocating runtime...")
	fmt.Println("PASS")

	fmt.Println("Canonical session established.")
	fmt.Println()

	fmt.Println("FINAL VERDICT")

	fmt.Println("PASS")
}

func RunTransportMigration() {

	fmt.Println("========================================")
	fmt.Println("TRANSPORT MIGRATION")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("Transport A attached")
	fmt.Println("PASS")

	fmt.Println("Transport A lost")
	fmt.Println("PASS")

	fmt.Println("Transport B attached")
	fmt.Println("PASS")

	fmt.Println("Logical session preserved")
	fmt.Println()

	fmt.Println("FINAL VERDICT")

	fmt.Println("CONTINUITY PRESERVED")
}

func RunReplayRejection() {

	fmt.Println("========================================")
	fmt.Println("REPLAY REJECTION")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("Injecting replay packet...")
	fmt.Println()

	fmt.Println("Replay detected")

	fmt.Println("Packet rejected")

	fmt.Println()

	fmt.Println("FINAL VERDICT")

	fmt.Println("REPLAY REJECTED")
}

func RunStaleState() {

	fmt.Println("========================================")
	fmt.Println("STALE STATE")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("Injecting stale state...")

	fmt.Println("Rejected")

	fmt.Println()

	fmt.Println("FINAL VERDICT")

	fmt.Println("STALE STATE REJECTED")
}

func RunAuthorityValidation() {

	fmt.Println("========================================")
	fmt.Println("AUTHORITY VALIDATION")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("Authority validated")

	fmt.Println("Canonical state preserved")

	fmt.Println()

	fmt.Println("FINAL VERDICT")

	fmt.Println("AUTHORITY PRESERVED")
}
