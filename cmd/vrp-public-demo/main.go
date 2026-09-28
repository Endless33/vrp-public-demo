package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		RunMenu()
		return
	}

	switch os.Args[1] {

	case "session":
		RunSession()

	case "migration":
		RunMigration()

	case "recovery":
		RunRecovery()

	case "replay":
		RunReplay()

	case "stale":
		RunStale()

	case "authority":
		RunAuthority()

	case "full":
		RunFullDemo()

	case "help":
		fmt.Println("VRP Public Demo")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  go run ./cmd/vrp-public-demo")
		fmt.Println("  go run ./cmd/vrp-public-demo session")
		fmt.Println("  go run ./cmd/vrp-public-demo migration")
		fmt.Println("  go run ./cmd/vrp-public-demo recovery")
		fmt.Println("  go run ./cmd/vrp-public-demo replay")
		fmt.Println("  go run ./cmd/vrp-public-demo stale")
		fmt.Println("  go run ./cmd/vrp-public-demo authority")
		fmt.Println("  go run ./cmd/vrp-public-demo full")
		fmt.Println("  go run ./cmd/vrp-public-demo help")

	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println()
		fmt.Println("Run:")
		fmt.Println("  go run ./cmd/vrp-public-demo help")
	}
}
