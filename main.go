package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nerzal/gocloak/v13"
	"github.com/dharm-kapadia/1source-go/api"
	"github.com/dharm-kapadia/1source-go/models"
	"github.com/dharm-kapadia/1source-go/utils"
)

var (
	LogFile   = "1source-go.log"
	fileName  string
	token     *gocloak.JWT
	appConfig *models.AppConfig
)

func main() {
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
		utils.DisplayHelp()

		// Graceful exit after displaying help
		os.Exit(0)
	}

	argsWithoutProg := os.Args[1:]

	// Command line of length 1 usually means help or version info requested
	if len(argsWithoutProg) == 1 {
		switch argsWithoutProg[0] {
		case "--help", "help", "-h":
			utils.DisplayHelp()
		case "--version", "-v":
			utils.DisplayVersion()
		default:
			utils.DisplayHelp()
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

		var endPoint string

		switch param {
		// Get all of a particular type from the API
		case "-g":
			switch entity {
			case "events":
				api.GetEntity(appConfig.Endpoints.Parties, bearer, "1Source Events")

			case "parties":
				api.GetEntity(appConfig.Endpoints.Parties, bearer, "1Source Parties")

			case "agreements":
				api.GetEntity(appConfig.Endpoints.Agreements, bearer, "1Source Agreements")

			case "contracts":
				api.GetEntity(appConfig.Endpoints.Contracts, bearer, "1Source Contracts")

			case "rerates":
				api.GetEntity(appConfig.Endpoints.Rerates, bearer, "1Source Rerates")

			case "returns":
				api.GetEntity(appConfig.Endpoints.Returns, bearer, "1Source Returns")

			case "recalls":
				api.GetEntity(appConfig.Endpoints.Recalls, bearer, "1Source Recalls")

			case "buyins":
				api.GetEntity(appConfig.Endpoints.Buyins, bearer, "1Source Buyins")

			default:
				log.Println("Unknown command-line entity entered: ", entity)
			}

		// Get trade agreement by agreement_id
		case "-a":
			endPoint = appConfig.Endpoints.Agreements + "/" + entity
			api.GetEntityById(endPoint, entity, bearer, "Agreement")

		// Get event agreement by event_id
		case "-e":
			endPoint = appConfig.Endpoints.Events + "/" + entity
			api.GetEntityById(endPoint, entity, bearer, "Event")

		// Get contract by contract_id
		case "-c":
			endPoint = appConfig.Endpoints.Contracts + "/" + entity
			api.GetEntityById(endPoint, entity, bearer, "Contract")

		// Get party by party_id
		case "-p":
			endPoint = appConfig.Endpoints.Parties + "/" + entity
			api.GetEntityById(endPoint, entity, bearer, "Party")

		// Propose contract
		case "-i":
			// Read on JSON file specified on the command line as bytes
			body, err := os.ReadFile(entity)
			if err != nil {
				fmt.Printf("Error JSON reading file [%s]: %s\n", entity, err)
				log.Printf("Error JSON reading file [%s]: %s\n", entity, err)
			}

			// Do HTTP PostProposeContract to initiate the contract
			// Errors are handled in the function
			_, err = api.PostProposeContract(appConfig.Endpoints.Contracts, bearer, body)

			if err == nil {
				fmt.Println("Successfully created Proposed Contract")
			}
		default:
			log.Println("Unknown command-line switch entered: ", argsWithoutProg)
		}
	}
}
