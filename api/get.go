// Package api provides functions for HTTP verb access to 1Source REST API.
package api

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Get performs an HTTP GET operation on the 1Source REST API
// It is used to get all entities of a type (Events, Parties,
// Trade Agreements, Contracts) or it can retrieve one of those
// entities based on an Id
// It returns the entities from the query and any error encountered.
func Get(apiEndpoint string, bearer string) (string, error) {
	ctx := context.Background()
	transport := &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			r.Header.Set("Authorization", bearer)
			return nil, nil
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(10) * time.Second,
	}

	request, err := http.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	request.Header.Set("Authorization", bearer)

	if err != nil {
		log.Println("Error creating new HTTP Request: ", err)
		return "", err
	}

	log.Println("Calling API endpoint: ", apiEndpoint)
	response, err := client.Do(request)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing Body.\n[ERR] -", err)
		}
	}(response.Body)

	if err != nil {
		log.Println("Error on response.\n[ERR] -", err)
	} else {
		if response.StatusCode != http.StatusOK {
			log.Println("Error in response status. [ERR] -", response.StatusCode)
		} else {
			data, _ := io.ReadAll(response.Body)
			return string(data), err
		}
	}

	return "", err
}
