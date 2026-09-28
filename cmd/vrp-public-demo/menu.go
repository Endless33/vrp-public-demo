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
		fmt.Println("1. Transport Migration")
		fmt.Println("2. Session Recovery")
		fmt.Println("0. Exit")
		fmt.Println()

		fmt.Print("Select: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		fmt.Println()

		switch input {

		case "1":
			RunMigration()

		case "2":
			RunRecovery()

		case "0":
			fmt.Println("Goodbye.")
			return

		default:
			fmt.Println("Unknown selection.")
		}
	}
}