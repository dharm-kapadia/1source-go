package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nerzal/gocloak/v13"
	api "github.com/dharm-kapadia/1source-go/api"
	"github.com/dharm-kapadia/1source-go/models"
	"github.com/dharm-kapadia/1source-go/utils"
)

var (
	LOG_FILE string = "1source-go.log"
	fileName string
	token    *gocloak.JWT
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
	var appConfig *models.AppConfig

	// Open the log file
	var logFile, err = os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic(err)
	}

	defer logFile.Close()
	log.SetOutput(logFile)

	if len(os.Args) < 1 {
		displayHelp()

		// Graceful exit after displaying help
		os.Exit(0)
	}

	argsWithoutProg := os.Args[1:]

	// Command line of length 1 usually means help or version info requested
	if len(argsWithoutProg) == 1 {
		switch argsWithoutProg[0] {
		case "--help", "help", "-h":
			displayHelp()
		case "--version", "-v":
			displayVersion()
		default:
			displayHelp()
		}

		// Graceful exit after displaying help
		os.Exit(0)
	}

	if len(argsWithoutProg) == 2 {
		// Command line of length 2 means -t TOML file
		if argsWithoutProg[0] == "-t" {
			fileName = argsWithoutProg[1]

			// Read and parse configuration TOML file
			appConfig, err = utils.ReadTOML(fileName)

			if err != nil {
				log.Println("Error reading and parsing configuration TOML file: ", err)
				os.Exit(100)
			}

			fmt.Println(appConfig.General.Auth_URL)
		} else {
			log.Println("Unknown command line flag combination")
			os.Exit(200)
		}

		// Graceful exit after reading and parsing configuration TOML file
		os.Exit(0)
	}

	if len(argsWithoutProg) == 3 {
		log.Println("Unknown command line flag combination")
		os.Exit(300)
	}

	if len(argsWithoutProg) == 4 {
		fileName = argsWithoutProg[1]

		// Read and parse configuration TOML file
		appConfig, err = utils.ReadTOML(fileName)

		if err != nil {
			log.Println("Error reading and parsing configuration TOML file: ", err)
			os.Exit(100)
		}

		token, err = api.GetAuthToken(appConfig)

		if err != nil {
			log.Panic("Error retrieving Auth Token: ", err)
			panic(err)
		} else {
			fmt.Println("Auth token: ", token.AccessToken)
		}

		// Get the 3rd command line parameter
		param := argsWithoutProg[2]
		entity := argsWithoutProg[3]

		switch param {
		// Get all of a particular type from teh API
		case "-o":
			switch entity {
			case "agreements":
				break

			case "contracts":
				break

			case "events":
				break

			case "parties":
				break
			}

		// Get trade agreement by agreement_id
		case "-a":
			break

		// Get event agreement by event_id
		case "-e":
			break

		// Get contract by contract_id
		case "-c":
			break

		// Get party by party_id
		case "-p":
			break
		}
	}
}
