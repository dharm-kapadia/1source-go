// Package api provides functions for HTTP verb access to 1Source REST API.
package api

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v13"
	"github.com/dharm-kapadia/1source-go/models"
)

// GetAuthToken logs into KeyCloak using credentials from the configuration
// TOML file to retrieve an Auth Token, which is used in subsequent calls to
// the 1Source REST API
func GetAuthToken(cfg *models.AppConfig) (*gocloak.JWT, error) {
	var token *gocloak.JWT
	var err error

	// Log into KeyCloak to get Auth Token
	log.Println("Logging into KeyCloak to get Auth Token")
	client := gocloak.NewClient(cfg.General.Auth_URL)
	ctx := context.Background()

	token, err = client.Login(
		ctx,
		cfg.Authentication.Client_Id,
		cfg.Authentication.Client_Secret,
		cfg.General.Realm_Name,
		cfg.Authentication.Username,
		cfg.Authentication.Password)

	if err != nil {
		log.Panic("Error retrieving Auth token", err)
	} else {
		log.Println("Successfully received Auth token")
	}

	return token, err
}
