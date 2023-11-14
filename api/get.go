package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type PartiesApiService service

func Get(apiEndpoint string, bearer string) (string, error) {
	client := &http.Client{Timeout: time.Duration(5) * time.Second}

	request, err := http.NewRequest(http.MethodGet, apiEndpoint, nil)
	request.Header.Set("Authorization", bearer)

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		req.Header.Set("Authorization", via[0].Header.Get("Authorization"))
		return nil
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
			fmt.Println(string(data))
			return string(data), err
		}
	}

	return "", err
}
