package main

import (
	"fmt"
)

type HttpStatus int

const (
	OK HttpStatus = 200
	NOT_FOUND HttpStatus = 404
	INTERNAL_SERVER_ERROR HttpStatus = 500
	CREATED HttpStatus = 201
	ACK HttpStatus = 301
)

var httpStatuses = map[HttpStatus]string {
	OK: 					"OK",
	NOT_FOUND: 				"NOT_FOUND",
	INTERNAL_SERVER_ERROR: 	"INTERNAL_SERVER_ERROR",
	CREATED: 				"CREATED",
	ACK: 					"ACK",
}

func (h HttpStatus) String() string {
	return httpStatuses[h]
}

type ServerState int

const (
	IDLE ServerState = iota
	STARTING
	STARTED
	CONNECTED
	ERROR
	RETRYING
	STOPPING
	STOPPED
)

var serverStates = map[ServerState]string {
	IDLE: 					"idle",
	STARTING: 				"starting",
	STARTED: 				"started",
	CONNECTED: 				"connected",
	ERROR: 					"error",
	RETRYING: 				"retrying",
	STOPPING: 				"stopping",
	STOPPED: 				"stopped",
}

func (s ServerState) String() string {
	return serverStates[s]
}

func Transition(s ServerState, count ...int) ServerState {
	if s == RETRYING && len(count) == 0  {
		fmt.Println("Still retrying.... Pass a `count` to stop retrying")
		return RETRYING
	}

	fmt.Println("transitioning from ", s, " to ", ServerState(s + 1))
	return ServerState(s + 1)
}

func GetStatus() {
	fmt.Println(HttpStatus(200));

	Transition(RETRYING)
}
