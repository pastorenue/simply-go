package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	SwitchMain()
}

var firstClassFunc = func() {
	var album = struct{
		title string
		artist string
		year int
		copies int
	}{
		title: "Rolling Stone",
		artist: "Harry K.",
		year: 1987,
		copies: 400,
	}
	str := "This is my first string slice"
	str_slice := strings.Split(str, " ")
	fmt.Println(str_slice)
	fmt.Println(album, &album)

	strSlice2 := append([]string{}, str_slice...)
	strSlice2 = append(strSlice2, "men", "of", "war")
	fmt.Printf("%#v, %T\n", strSlice2, strSlice2)
	fmt.Printf("Percentage: %d%%", 40)

	var i interface{}

	describe(i)

	i = 42
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func f() {
	fmt.Println("a")
	panic("a")
	fmt.Println("b")
}

func anotherFunc(vals ...int) int {
	var sum int
	for _, val := range vals {
		sum += val
	}
	return sum
}

func simpleFunc() {
	numbers := []int{1, 2, 3, 4, 5}
	s := anotherFunc(numbers...)
	fmt.Println("Sum of numbers is", s)
}

func getRandomDataHere() {
	const apiKey = "3DrIWOOfGZ_Crzt69TQ91A"
	const projectId = "476d6039-2182-4cc4-b965-983be063afc7"
	data, err := GetRandomData(projectId, apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Random data:", string(data))
	result, _ := json.MarshalIndent(data, "", "  ")

	fmt.Println(string(result))
}

func SampleApp() []User {
	const NUM_TICKETS int = 50
	var (
		conferenceName   string = "'GopherCon'"
		requestedTickets int
		bookingsArray    = [NUM_TICKETS]string{}
		bookings         []User
	)

	bookingsArray[0] = "John Doe"
	bookingsArray[1] = "Jane Doe"

	fmt.Println("Memory location of conferenceName:", &conferenceName)
	fmt.Println("Welcome to the", conferenceName, "Conference")
	fmt.Printf("We have %v tickets available\n", NUM_TICKETS)
	fmt.Print(("Please enter the number of tickets you would like to purchase: "))
	fmt.Scan(&requestedTickets)
	remainingTickets := NUM_TICKETS - requestedTickets
	fmt.Printf("You have requested %v tickets\n", requestedTickets)
	fmt.Println("Please enter the user details for each ticket below:")
	fmt.Printf("There are now %v tickets remaining\n", remainingTickets)

	for i := 0; i < requestedTickets; i++ {
		newUser := createUser()
		bookings = append(bookings, newUser)
		fmt.Println("Booking successful for", newUser)
	}

	return bookings
}
