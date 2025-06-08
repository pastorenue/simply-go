package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	simpleGoroutine()
}

func simpleGoroutine() {
	n := 1000000000000000
	c := make(chan int)
	t := time.NewTicker(time.Second * 5)

	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()
	
	buffer := make([]int, 0)
	sum := 0
	for {
		select {
		case i, ok := <-c:
			if !ok {
				for _, i := range buffer {
					sum += i
				}
				fmt.Println("sum:", sum)
				fmt.Println("--------------------")
				sum = 0
				buffer = make([]int, 0)
				return
			}
			buffer = append(buffer, i)
		case <-t.C:
			for _, i := range buffer {
				sum += i
			}
			fmt.Println	("sum:", sum)
			fmt.Println("--------------------")
			sum = 0
			buffer = make([]int, 0)
		case <-time.After(time.Second * 20):
			println("timeout")
			return
		}
	}
}

func goroutineWithWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1 finished")
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("Goroutine 2 finished")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 3 finished")
	}()

	wg.Wait()
	fmt.Println("All goroutines finished")
}


type Aggregate struct {
	count int
	buckets map[string]int
}
