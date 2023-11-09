package utils

import (
	"fmt"
	"io"
	"log"
	"os"

	models "github.com/dharm-kapadia/1source-go/models"
	"github.com/pelletier/go-toml/v2"
)

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
