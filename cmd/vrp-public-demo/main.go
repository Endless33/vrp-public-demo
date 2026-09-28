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
<<<<<<< HEAD
		fmt.Println("  go run ./cmd/vrp-public-demo migration")
		fmt.Println("  go run ./cmd/vrp-public-demo recovery")

	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println()
		fmt.Println("Use:")
=======
		fmt.Println("  go run ./cmd/vrp-public-demo session")
		fmt.Println("  go run ./cmd/vrp-public-demo migration")
		fmt.Println("  go run ./cmd/vrp-public-demo recovery")
		fmt.Println("  go run ./cmd/vrp-public-demo replay")
		fmt.Println("  go run ./cmd/vrp-public-demo stale")
		fmt.Println("  go run ./cmd/vrp-public-demo authority")
		fmt.Println("  go run ./cmd/vrp-public-demo full")

	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Run:")
>>>>>>> c22a71e (Refactor public demo into interactive scenario runner)
		fmt.Println("  go run ./cmd/vrp-public-demo help")
	}
}
