// Package api provides functions for HTTP verb access to 1Source REST API.
package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Get performs an HTTP GET operation on the 1Source REST API
// It is used to get all entities of a type (Events, Parties,
// Trade Agreements, Contracts) or it can retrieve one of those
// entities based on an Id
// It returns the entities from the query and any error encountered.
func Get(apiEndPoint string, bearer string) (string, error) {
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

	request, err := http.NewRequestWithContext(ctx, "GET", apiEndPoint, nil)
	request.Header.Set("Authorization", bearer)

	if err != nil {
		log.Println("Error creating new HTTP Request: ", err)
		return "", err
	}

	log.Println("Calling API endpoint: ", apiEndPoint)
	response, err := client.Do(request)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing Body.\n[ERR] -", err)
		}
	}(response.Body)

	if err != nil {
		log.Println("Error in response.\n[ERR] -", err)
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

// getEntityById is a helper function to perform an HTTP GET to
// retrieve a particular entity by Id from the 1Source REST API
func GetEntityById(endPoint string, id, bearer string, header string) {
	agreement, err := Get(endPoint, bearer)
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
func GetEntity(endPoint string, bearer string, header string) {
	entity, err := Get(endPoint, bearer)
	if err == nil {
		fmt.Println(header)
		fmt.Println(strings.Repeat("=", len(header)))
		fmt.Println(entity)
	} else {
		log.Printf("Error GET %s: %s", header, err)
	}
}
