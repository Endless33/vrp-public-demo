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

	case "migration":
		RunMigration()

	case "recovery":
		RunRecovery()

	case "help":
		fmt.Println("VRP Public Demo")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  go run ./cmd/vrp-public-demo")
		fmt.Println("  go run ./cmd/vrp-public-demo migration")
		fmt.Println("  go run ./cmd/vrp-public-demo recovery")

	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println()
		fmt.Println("Use:")
		fmt.Println("  go run ./cmd/vrp-public-demo help")
	}
}