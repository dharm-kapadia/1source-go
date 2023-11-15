package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/dharm-kapadia/1source-go/api"
	"github.com/dharm-kapadia/1source-go/models"
	"github.com/dharm-kapadia/1source-go/utils"
)

var (
	LogFile  = "1source-go.log"
	fileName string
	token    *gocloak.JWT
	endPoint string
)

// displayVersion prints the program version
func displayVersion() {
	fmt.Println("1source-go V0.1")
}

// displayHelp creates the complete help string output for the command line
func displayHelp() {
	fmt.Print("Usage: 1Source [--help] [--version] -t VAR [-o VAR] [-a VAR] [-e VAR] [-c VAR] [-p VAR]\n")
	fmt.Print("Note: -t is required\n\n")
	fmt.Println("Optional arguments:")
	fmt.Println("-h, --help\tshows help message and exits")
	fmt.Println("-v, --version\tprints version information and exits")
	fmt.Println("-t\t\t1Source configuration TOML file [required]")
	fmt.Println("-g\t\t1Source API Endpoint to query [agreements, contracts, events, parties, returns, rerates, recalls, buyins ]")

	fmt.Println("-a\t\t1Source API Endpoint to query trade agreements by agreement_id")
	fmt.Println("-e\t\t1Source API Endpoint to query events by event_id")
	fmt.Println("-c\t\t11Source API Endpoint to query contracts by contract_id")
	fmt.Println("-p\t\t1Source API Endpoint to query parties by party_id")
}

func main() {
	var appConfig *models.AppConfig

	// Open the log file
	var logFile, err = os.OpenFile(LogFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic("Error opening the Log file: ", err)
	}

	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Printf("Error closing the log file	 '%s'", LogFile)
		}
	}(logFile)

	// Set the log to output the Log file
	log.SetOutput(logFile)

	// Begin parsing the command line arguments
	if len(os.Args) == 1 {
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

	// Command line of length 2 means either help or version
	if len(argsWithoutProg) == 2 {
		// Command line of length 2 means -t TOML file
		if argsWithoutProg[0] == "-t" {
			fileName = argsWithoutProg[1]

			// Read and parse configuration TOML file
			appConfig, err = utils.ReadTOML(fileName)

			if err != nil {
				log.Println("Error reading and parsing configuration TOML file: ", err)
				os.Exit(10)
			}
		} else {
			log.Println("Unknown command line flag combination")
			os.Exit(20)
		}

		// Graceful exit after reading and parsing configuration TOML file
		os.Exit(0)
	}

	// Command line of length 3 is not supported
	if len(argsWithoutProg) == 3 {
		log.Println("Unknown command line flag combination")
		os.Exit(30)
	}

	// Command line of length 4 contains the actual command to execute
	if len(argsWithoutProg) == 4 {
		fileName = argsWithoutProg[1]

		// Read and parse configuration TOML file
		appConfig, err = utils.ReadTOML(fileName)

		if err != nil {
			log.Println("Error reading and parsing configuration TOML file: ", err)
			os.Exit(15)
		}

		// Get Auth Token using credentials from config file
		token, err = api.GetAuthToken(appConfig)
		var bearer string

		if err != nil {
			log.Panic("Error retrieving Auth Token: ", err)
		} else {
			bearer = `Bearer ` + token.AccessToken
		}

		// Get the 3rd and 4th command line parameters
		// The 3rd parameter will be a switch, the 4th parameter will be the entity
		param := argsWithoutProg[2]
		entity := argsWithoutProg[3]

		switch param {
		// Get all of a particular type from the API
		case "-g":
			switch entity {
			case "events":
				getEntity(appConfig.Endpoints.Parties, bearer, "1Source Events")

			case "parties":
				getEntity(appConfig.Endpoints.Parties, bearer, "1Source Parties")

			case "agreements":
				getEntity(appConfig.Endpoints.Agreements, bearer, "1Source Agreements")

			case "contracts":
				getEntity(appConfig.Endpoints.Contracts, bearer, "1Source Contracts")

			case "rerates":
				getEntity(appConfig.Endpoints.Rerates, bearer, "1Source Rerates")

			case "returns":
				getEntity(appConfig.Endpoints.Returns, bearer, "1Source Returns")

			case "recalls":
				getEntity(appConfig.Endpoints.Recalls, bearer, "1Source Recalls")

			case "buyins":
				getEntity(appConfig.Endpoints.Buyins, bearer, "1Source Buyins")
			}

		// Get trade agreement by agreement_id
		case "-a":
			endPoint = appConfig.Endpoints.Agreements + "/" + entity
			getEntityById(endPoint, entity, bearer, "Agreement")

		// Get event agreement by event_id
		case "-e":
			endPoint = appConfig.Endpoints.Events + "/" + entity
			getEntityById(endPoint, entity, bearer, "Event")

		// Get contract by contract_id
		case "-c":
			endPoint = appConfig.Endpoints.Contracts + "/" + entity
			getEntityById(endPoint, entity, bearer, "Contract")

		// Get party by party_id
		case "-p":
			endPoint = appConfig.Endpoints.Parties + "/" + entity
			getEntityById(endPoint, entity, bearer, "Party")

		// Propose contract
		case "-i":
			endPoint = appConfig.Endpoints.Contracts
		}
	}
}

// getEntityById is a helper function to perform an HTTP GET to
// retrieve a particular entity by Id from the 1Source REST API
func getEntityById(endPoint string, id, bearer string, header string) {
	agreement, err := api.Get(endPoint, bearer)
	if err == nil {
		fmt.Println(header)
		fmt.Println(strings.Repeat("=", len(header)))
		fmt.Println(agreement)
	} else {
		log.Printf("Error GET %s by id [%s]: %s", header, id, err)
	}
}

// getEntity is a helper function to perform an HTTP GET
// to retrieve entity-level data from the 1Source REST API
func getEntity(endPoint string, bearer string, header string) {
	entity, err := api.Get(endPoint, bearer)
	if err == nil {
		fmt.Println(header)
		fmt.Println(strings.Repeat("=", len(header)))
		fmt.Println(entity)
	} else {
		log.Printf("Error GET %s: %s", header, err)
	}
}
