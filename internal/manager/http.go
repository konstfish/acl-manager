package manager

import (
	"io"
	"log"
	"net/http"
)

func downloadList(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading the response body: %v", err)
		return nil, err
	}

	bodyString := string(bodyBytes)
	return parseNetList(bodyString), nil
}
