package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)


func Checks() error {
	i, err := strconv.Atoi(os.Getenv("VSCODE_PROFILE_INITIALIZED"))
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
		panic(err)
	}
	if i != 1 {
		return errors.New("vscode is not initialized yet")
	}
	return nil
}

func init() {
	err := Checks()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Init VS Code: ", os.Getenv("VSCODE_PROFILE_INITIALIZED"))
}


func somethingGood() {
	const maxWorkers  = 5
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= maxWorkers; w++ {
		go workers(w, jobs, results)
	}

	for j := 1; j <= 20; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 20; a++ {
		<-results
	}
}

func workers(id int, jobs <-chan int, results chan<- int ) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(200 * time.Millisecond)
		results <- j * 2
	}
}