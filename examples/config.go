package main

type DB struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Config struct {
	Port     	int
	Host     	string
	Secret		string
	LogLevel	string
	Mode		string
}