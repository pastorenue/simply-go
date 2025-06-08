package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var (
	userID = 100
	clientTimeout = 200
	apiKey = "simple_api_key"
	projectID = "simple_project_id"
)

func main() {
	uri := "http://localhost:81/"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("X-Project-ID", projectID)
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User-ID", strconv.Itoa(userID))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error occured", err)
	}

	defer res.Body.Close()
}