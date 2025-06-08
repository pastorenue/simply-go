package main

import (
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var database []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Error connecting to RPC server:", err)
		return
	}
	a := Item{"Sample Title", "Sample Body"}
	b := Item{"Another Title", "Another Body"}
	c := Item{"Edit Title", "Edit Body"}
	d := Item{"Another Title", "Another Body"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.AddItem", d, &reply)

	client.Call("API.GetAll", 0, &database)
	log.Println("Items in database:", database)
}
