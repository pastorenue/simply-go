package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	firstName string
	lastName string
	age uint64
	fullName string
}

func (User) new(fname string, lname string, age uint64) User {
	fullname := fname + " " + lname
	return User {
		firstName: fname,
		lastName: lname,
		age: age,
		fullName: fullname,
	}
}

func createUser() User {
	var userData = make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your first name: ")
	firstName, _ := reader.ReadString('\n')
	userData["firstName"] = strings.TrimSpace(firstName)
	fmt.Print("Enter your last name: ")
	lastName, _ := reader.ReadString('\n')
	userData["lastName"] = strings.TrimSpace(lastName)
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')
	a, _ := strconv.ParseUint(age, 10, 64)
	fullName := strings.TrimSpace(firstName) + " " + strings.TrimSpace(lastName)
	userData["fullName"] = fullName
	if !isValidName(fullName) {
		fmt.Println("Invalid name. Please enter a valid name")
		return createUser()
	}
	// user := User {
	// 	firstName: strings.TrimSpace(firstName),
	// 	lastName: strings.TrimSpace(lastName),
	// 	age: a,
	// 	fullName: fullName,
	// }
	user := User.new(User{}, userData["firstName"], userData["lastName"], a)

	return user
}

func isValidName(val string) bool {
	return len(val) >= 5
}

func createTickets(users []User) {
	fmt.Println("All tickets have been booked. See receipt below:")
	fmt.Println()
	fmt.Println("--------------------")
	for idx, user := range users {
		fmt.Printf("Ticket %v\n", idx + 1)
		fmt.Printf("First Name: %v\n", user.firstName)
		fmt.Printf("Last Name: %v\n", user.lastName)
		fmt.Printf("Full Name: %v\n", user.fullName)
		fmt.Printf("Age: %v\n", user.age)
		println("--------------------")
	}
}

func getOps() {
	operations := make([]float32, 0)
	if len(operations) != 0 {
		handle(operations)
	}
}

func handle(operations []float32) {}

