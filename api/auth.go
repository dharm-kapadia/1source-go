package auth

import (
	"context"
	"log"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/dharm-kapadia/1source-go/models"
)

func GetAuthToken(cfg *models.AppConfig) (*gocloak.JWT, error) {
	var token *gocloak.JWT
	var err error

	// Log into KeyCloak to get Auth Token
	log.Println("Logging into Keyclock to get Auth Token")
	client := gocloak.NewClient(cfg.General.Auth_URL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	token, err = client.LoginClient(ctx, cfg.Authentication.Client_Id, cfg.Authentication.Client_Secret, cfg.General.Realm_Name)

	if err != nil {
		log.Panic("Error retrieving Auth token", err)
		panic("Error retrieving Auth token")
	}

	return token, err
}
