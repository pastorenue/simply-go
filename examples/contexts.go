package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)


func MakeContextCall() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string {
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/posts",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal(err)
		results <- fmt.Sprintf("Error creating request for %s:%s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request for %s:%s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	log.Default().Println(resp.StatusCode)
	results <- fmt.Sprintf("Success for %s: %d", url, resp.StatusCode)
}

func GetPolygonTradeData() {
	c := polygon.New("oZ4e004INQ3expS7sh8CntKp6NbVf8Kp")

	// Set params
	params := models.GetExchangesParams{}.
		WithAssetClass((models.AssetClass("stocks"))).
		WithLocale(models.MarketLocale("us"))

	// Make the request
	resp, err := c.GetExchanges(context.Background(), params)
	if err != nil {
		log.Fatalf("Error fetching trade data: %v", err)
	}
	// Print the response
	fmt.Printf("Exchange data:\n", resp.Results)
}