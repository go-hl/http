package http

import (
	"io"
	"net/http"
	"os"
)

// FileGET do a GET HTTP request to download a file.
func FileGET(filePath string, URL string, headers ...map[string]string) error {

	// Set a new client.
	var client = new(http.Client)

	// Set a new request.
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return err
	}

	// Set headers on request.
	for header, value := range headers[len(headers)-1] {
		request.Header.Set(header, value)
	}

	// Do the request using the client.
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create the file to put data.
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the body to file.
	_, err = io.Copy(file, response.Body)
	return err

}
