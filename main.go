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

		switch param {
		// Get all of a particular type from the API
		case "-g":
			switch entity {
			case "events":
				header := "1Source Events"
				events, err := api.GetEntity(appConfig.Endpoints.Events, bearer, header)
				utils.PrintResults(err, events, "Error retrieving 1Source Events: ", header)

			case "parties":
				header := "1Source Parties"
				parties, err := api.GetEntity(appConfig.Endpoints.Parties, bearer, header)
				utils.PrintResults(err, parties, "Error retrieving 1Source Parties: ", header)

			case "agreements":
				header := "1Source Trade Agreements"
				tas, err := api.GetEntity(appConfig.Endpoints.Agreements, bearer, header)
				utils.PrintResults(err, tas, "Error retrieving 1Source Trade Agreements: ", header)

			case "contracts":
				header := "1Source Contracts"
				contracts, err := api.GetEntity(appConfig.Endpoints.Contracts, bearer, header)
				utils.PrintResults(err, contracts, "Error retrieving 1Source Contracts: ", header)

			case "rerates":
				header := "1Source Rerates"
				rerates, err := api.GetEntity(appConfig.Endpoints.Rerates, bearer, header)
				utils.PrintResults(err, rerates, "Error retrieving 1Source Rerates: ", header)

			case "returns":
				header := "1Source Returns"
				returns, err := api.GetEntity(appConfig.Endpoints.Returns, bearer, header)
				utils.PrintResults(err, returns, "Error retrieving 1Source Returns: ", header)

			case "recalls":
				header := "1Source Recalls"
				recalls, err := api.GetEntity(appConfig.Endpoints.Recalls, bearer, header)
				utils.PrintResults(err, recalls, "Error retrieving 1Source Recalls: ", header)

			case "buyins":
				header := "1Source Buyins"
				buyins, err := api.GetEntity(appConfig.Endpoints.Buyins, bearer, header)
				utils.PrintResults(err, buyins, "Error retrieving 1Source Buyins: ", header)

			default:
				log.Println("Unknown command-line entity entered: ", entity)
				fmt.Println("Unknown command-line entity entered: ", entity)
			}

		// Get trade agreement by agreement_id
		case "-a":
			header := "1Source Trade Agreement"
			agreement, err := api.GetEntityById(appConfig.Endpoints.Agreements, entity, bearer, header)
			prompt := fmt.Sprintf("Error retrieving 1Source with Trade Agreement Id = [%s]: ", entity)
			utils.PrintResults(err, agreement, prompt, header)

		// Get event agreement by event_id
		case "-e":
			header := "1Source Event"
			event, err := api.GetEntityById(appConfig.Endpoints.Events, entity, bearer, header)
			prompt := fmt.Sprintf("Error retrieving 1Source with Event Id = [%s]: ", entity)
			utils.PrintResults(err, event, prompt, header)

		// Get contract by contract_id
		case "-c":
			header := "1Source Contract"
			contract, err := api.GetEntityById(appConfig.Endpoints.Contracts, entity, bearer, header)
			prompt := fmt.Sprintf("Error retrieving 1Source with Contract Id = [%s]: ", entity)
			utils.PrintResults(err, contract, prompt, header)

		// Get party by party_id
		case "-p":
			header := "1Source Party"
			party, err := api.GetEntityById(appConfig.Endpoints.Parties, entity, bearer, "Party")
			prompt := fmt.Sprintf("Error retrieving 1Source with party Id = [%s]: ", entity)
			utils.PrintResults(err, party, prompt, header)

		// Propose contract
		case "-cp":
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

		// Cancel and proposed contract
		case "-ca":
			// Get the Contract by contract_id to check that it is in the proposed state
			contract, err := api.GetEntityById(appConfig.Endpoints.Contracts, entity, bearer, "1Source Contract")

			if err != nil {
				log.Printf("Error GET %s by id [%s]: %s", "Contract", entity, err)
			} else {
				// Check the state of the contract
				if strings.Contains(contract, "PROPOSED") {
					fmt.Println(contract)
				} else {
					fmt.Printf("Contract with id [%s] is not in PROPOSED state\n", entity)
				}
			}

		default:
			log.Println("Unknown command-line switch entered: ", argsWithoutProg)
		}
	}
}
