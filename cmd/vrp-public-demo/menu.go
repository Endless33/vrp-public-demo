package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("======================================")
		fmt.Println("VRP PUBLIC DEMO")
		fmt.Println("======================================")
		fmt.Println()
		fmt.Println("1. Session Establishment")
		fmt.Println("2. Transport Migration")
		fmt.Println("3. Session Recovery")
		fmt.Println("4. Replay Rejection")
		fmt.Println("5. Stale-State Rejection")
		fmt.Println("6. Authority Validation")
		fmt.Println("7. Full Demonstration")
		fmt.Println("0. Exit")
		fmt.Println()

		fmt.Print("Select: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		input = strings.TrimSpace(input)
		fmt.Println()

		switch input {

		case "1":
			RunSession()

		case "2":
			RunMigration()

		case "3":
			RunRecovery()

		case "4":
			RunReplay()

		case "5":
			RunStale()

		case "6":
			RunAuthority()

		case "7":
			RunFullDemo()

		case "0":
			fmt.Println("Goodbye.")
			return

		default:
			fmt.Println("Unknown selection.")
		}
	}
}
