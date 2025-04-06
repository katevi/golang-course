// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	getWithHeaders()
}

func getWithHeaders() {
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://go.dev/play/", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add headers
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("User-Agent", "MyGoApp")

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Process response
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("GitHub API Response:", string(body))
}
