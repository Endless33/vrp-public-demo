package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const version = "v0.1.0"

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		printBanner()

		fmt.Print("Select scenario: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			notImplemented("Session Establishment")

		case "2":
			notImplemented("Transport Migration")

		case "3":
			notImplemented("Replay Rejection")

		case "4":
			notImplemented("Stale-State Rejection")

		case "5":
			notImplemented("Authority Validation")

		case "6":
			notImplemented("Full Demonstration")

		case "0":
			fmt.Println()
			fmt.Println("Goodbye.")
			return

		default:
			fmt.Println()
			fmt.Println("Unknown option.")
			pause(reader)
		}
	}
}

func printBanner() {

	fmt.Println("====================================================")
	fmt.Println("VRP PUBLIC DEMO")
	fmt.Println("Veil Routing Protocol")
	fmt.Println("Version:", version)
	fmt.Println("====================================================")
	fmt.Println()
	fmt.Println("SESSION \u2260 TRANSPORT")
	fmt.Println()
	fmt.Println("1. Session Establishment")
	fmt.Println("2. Transport Migration")
	fmt.Println("3. Replay Rejection")
	fmt.Println("4. Stale-State Rejection")
	fmt.Println("5. Authority Validation")
	fmt.Println("6. Full Demonstration")
	fmt.Println()
	fmt.Println("0. Exit")
	fmt.Println()
}

func notImplemented(name string) {

	fmt.Println()
	fmt.Println("----------------------------------------------------")
	fmt.Println(name)
	fmt.Println("----------------------------------------------------")
	fmt.Println()
	fmt.Println("This public demonstration scenario")
	fmt.Println("will be implemented in the next revision.")
	fmt.Println()
	fmt.Println("Architecture is public.")
	fmt.Println("Protected runtime implementation remains private.")
	fmt.Println()
}

func pause(reader *bufio.Reader) {

	fmt.Print("Press ENTER to continue...")
	reader.ReadString('\n')
}