package manager

import (
	"io"
	"log"
	"net/http"
)

func getHTTPList(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading the response body: %v", err)
		return "", err
	}

	bodyString := string(bodyBytes)
	return bodyString, nil
}
