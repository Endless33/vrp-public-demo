package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		RunMigration()
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
		fmt.Println("  vrp-public-demo migration")
		fmt.Println("  vrp-public-demo recovery")

	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println("  migration")
		fmt.Println("  recovery")
		fmt.Println("  help")
	}
}