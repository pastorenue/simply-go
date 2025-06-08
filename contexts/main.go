package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num ++ {
		printCh <- num
	}
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("doSomething: finished\n")
}

func doSomethingWithDeadline(ctx context.Context) {
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()
	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 0; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <- ctx.Done():
			break
		}
	}
	cancelCtx()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("doSomethingWithDeadline: finished\n")

}

func doSomethingWithTimeout(ctx context.Context) {
	ctx,cancelCtx := context.WithTimeout(ctx, 1500 * time.Millisecond)
	defer cancelCtx()
	printCh := make(chan int)
	go doAnother(ctx, printCh)
	for num := 0; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <- ctx.Done():
			break
		}
	}
	cancelCtx()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("doSomethingWithTimeout: finished\n")
}

func doAnother(ctx context.Context, printCh <- chan int) {
	for {
		select {
		case <- ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother: %s\n", err)
				return
			} else {
				fmt.Println("doAnother: finished\n")
				return
			}
		case num := <- printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

type ApiKey string
const api_key ApiKey = "api_key"

type Result struct {
	ApiKey string
	UserID int
}

func main() {
	ProcessCSV(os.Stdout)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
		apiKey := r.Header.Get("X-API-Key")
		userID, _ := strconv.Atoi(r.Header.Get("X-User-ID"))
		if apiKey == "" {
			http.Error(w, "API Key is required", http.StatusBadRequest)
			return
		}
		if userID == 0 {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}
		if userID != 101 {
			http.Error(w, "User ID is not authorized", http.StatusForbidden)
			return
		}
		logger.Info("Received request", "userID", userID, "apiKey", apiKey)
		ctx := context.WithValue(r.Context(), api_key, apiKey)
		ctx = context.WithValue(ctx, "userID", userID)
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		logger.Debug("Context created", "apiKey", apiKey, "userID", userID)
		resultChan := make(chan Result)
		go func() {
			resultChan <- callRemoteAPI(ctx, logger)
		}()

		select {
		case <- ctx.Done():
			logger.Error("Request timed out")
		case result := <- resultChan:
			fmt.Printf("Result received: %+v\n", result)
			logger.Debug("Result: %s\n", result)
		}
	})

	fmt.Println("Server running on port 81")
	http.ListenAndServe(":81", nil)
}

func callRemoteAPI(ctx context.Context, log *slog.Logger) Result {
	if err := ctx.Err(); err != nil {
		log.Error("Error occured", err.Error())
		fmt.Sprintf("Error: %s", err)
	}
	key := ctx.Value(api_key).(string)
	id, _ := ctx.Value("userID").(int)

	result := Result{
		ApiKey: key,
		UserID: id,
	}

	// Simulate a successful API call
	log.Debug("Call was successful", result)
	return result
}