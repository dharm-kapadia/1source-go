package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type PartiesApiService service

func Get(apiEndpoint string, bearer string) (string, error) {

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, apiEndpoint, nil)

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}

		log.Println("Received Redirect Error:", err)
		return err
	}

	request.Header.Set("Authorization", bearer)
	request.Header.Set("User-Agent", "1source-go Command Line")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("ContentType", "application/x-www-form-urlencoded; charset=UTF-8")

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
