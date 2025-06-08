package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ProcessPipe() {
	stats, err := os.Stdin.Stat()
	if err != nil {
		fmt.Errorf("An error occurred: %v", err)
		return
	}
	fmt.Printf("IsPiped: %d\n", (stats.Mode() & os.ModeCharDevice))
	isPiped := (stats.Mode() & os.ModeCharDevice) == 0
	if isPiped {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	} else if len(os.Args) > 1{
		file, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Errorf("An error occurred: %v", err)
			return
		}
		fmt.Println(string(file))
	} else {
		fmt.Println("No input provided")
		fmt.Println("Example usage: go run piping.go < file.txt")
		fmt.Println("or: go run piping.go file.txt")
		fmt.Println("or: cat file.txt | go run piping.go")
	}
}

type Entry struct {
	Level string `json:"level"`
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

