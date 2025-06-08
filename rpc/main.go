package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body string
}

type API int

var database []Item

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item
	for _, val := range database {
		if val.Title == title {
			getItem = val
			break
		}
	}
	*reply = getItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) GetAll(value int, allItem *[]Item) error {
	res := make([]Item, len(database))
	copy(res, database)
	*allItem = res
	return nil
}

func EditItem(edit Item, reply *Item) error {
	var changed Item
	for i, val := range database {
		if val.Title == edit.Title {
			database[i] = Item{edit.Title, edit.Body}
			changed = database[i]

		}
	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem(title string, reply *Item) error {
	for i, val := range database {
		if val.Title == title {
			database = append(database[:i], database[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item with title '%s' not found", title)
}

// main function to demonstrate the functionality
func main() {
	// Initialize the API and database
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error registering API:", err)
		return
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Error starting listener:", err)
		return
	}
	log.Println("RPC server listening on port 4040")
	listener, err = net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Error starting listener:", err)
		return
	}
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving HTTP:", err)
		return
	}
	log.Println("RPC server started successfully")
}