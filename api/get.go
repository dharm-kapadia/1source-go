package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type PartiesApiService service

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
			fmt.Println(string(data))
			return string(data), err
		}
	}

	return "", err
}
