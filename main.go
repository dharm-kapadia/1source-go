package main

import (
	"fmt"
	"os"
)

func displayVersion() {
	fmt.Println("1source-go V1.0.2")
}

func displayHelp() {
	fmt.Print("Usage: 1Source [--help] [--version] -t VAR [-o VAR] [-a VAR] [-e VAR] [-c VAR] [-p VAR]\n\n")
	fmt.Println("Optional arguments:")
	fmt.Println("-h, --help\tshows help message and exits")
	fmt.Println("-v, --version\tprints version information and exits")
	fmt.Println("-t\t\t1Source configuration TOML file [required]")
	fmt.Println("-o\t\t1Source API Endpoint to query [agreements, contracts, events, parties ]")

	fmt.Println("-a\t\t1Source API Endpoint to query trade agreements by agreement_id")
	fmt.Println("-e\t\t1Source API Endpoint to query events by event_id")
	fmt.Println("-c\t\t11Source API Endpoint to query contracts by contract_id")
	fmt.Println("-p\t\t1Source API Endpoint to query parties by party_id")
}

func main() {
	argsWithoutProg := os.Args[1:]

	// Command line of length 1 usually means help or version info requested
	if len(argsWithoutProg) == 1 {
		switch argsWithoutProg[0] {
		case "--help":
			displayHelp()
		case "-h":
			displayHelp()
		case "--version":
			displayVersion()
		case "-v":
			displayVersion()
		}

		// Graceful exit after displaying help
		os.Exit(0)
	}

	if len(argsWithoutProg) == 2 {
		// Command line of length 2 means -t TOML file
		if argsWithoutProg[0] == "-t" {
			filename := argsWithoutProg[1]
			fmt.Printf("TOML file name = %s\n", filename)

		} else {
			fmt.Println("Unknown command line flag combination")
		}
	}
}
