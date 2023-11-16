// utils contains utility functions
package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	models "github.com/dharm-kapadia/1source-go/models"
	"github.com/pelletier/go-toml/v2"
)

// FileExists checks that the specified file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		log.Printf("Configuration TOML file '%s' does not exist", filename)
		return false
	} else {
		log.Printf("Configuration TOML file '%s' exists!", filename)
	}

	return !info.IsDir()
}

// ReadTOML opens and reads in application configuration TOML file
func ReadTOML(filename string) (*models.AppConfig, error) {
	var appConfig models.AppConfig
	var err error
	var file *os.File

	if FileExists(filename) {
		// TOML file exists  ... open it
		file, err = os.Open(filename)

		if err != nil {
			log.Panic(err)
			panic(err)
		}

		defer file.Close()

		// Read file contents
		b, err := io.ReadAll(file)

		if err != nil {
			log.Panic(err)
			panic(err)
		}

		// Unmarshall it from TOML to defined struct
		err = toml.Unmarshal(b, &appConfig)

		if err != nil {
			log.Panic(err)
			panic(err)
		} else {
			log.Printf("Successfully read and parsed '%s'\n", filename)
		}
	} else {
		fmt.Printf("Configuration TOML file '%s' does not exists\n", filename)
	}

	return &appConfig, err
}

// DisplayVersion prints the program version
func DisplayVersion() {
	fmt.Println("1source-go V0.2")
}

// DisplayHelp creates the complete help string output for the command line
func DisplayHelp() {
	fmt.Print("Usage: 1Source [--help] [--version] -t VAR [-g VAR] [-a agreement_id] [-e events] [-c contract_id] [-p party_id] [-i JSON]\n")
	fmt.Print("Note: -t is required\n\n")
	fmt.Println("Optional arguments:")
	fmt.Println("-h, --help\tshows help message and exits")
	fmt.Print("-v, --version\tprints version information and exits\n\n")
	fmt.Println("-t\t\t1Source configuration TOML file [required]")
	fmt.Println("-g\t\t1Source API Endpoint to query [agreements, contracts, events, parties, returns, rerates, recalls, buyins]")

	fmt.Println("-a\t\t1Source API Endpoint to query trade agreements by agreement_id")
	fmt.Println("-e\t\t1Source API Endpoint to query events by event_id")
	fmt.Println("-c\t\t1Source API Endpoint to query contracts by contract_id")
	fmt.Print("-p\t\t1Source API Endpoint to query parties by party_id\n\n")

	fmt.Println("-cp\t\t1Source API Endpoint to propose a contract from a JSON file")
	fmt.Println("-cc\t\t1Source API Endpoint to cancel a proposed contract by contract_id")
	fmt.Println("-ca\t\t1Source API Endpoint to approve a proposed contract by contract_id")
	fmt.Println("-cd\t\t1Source API Endpoint to decline a proposed contract by contract_id")
	fmt.Println("")
}

// PrintResults outputs entities fetched from the 1Source REST API to the console
func PrintResults(err error, entity string, prompt string, header string) {
	if err != nil {
		fmt.Println(prompt, err)
	} else {
		fmt.Println(header)
		fmt.Println(strings.Repeat("=", len(header)))
		fmt.Println(entity)
	}
}
