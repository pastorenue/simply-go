package main

type OptFunc func(a *Config)

func WithPort(p int) OptFunc {
	return func(a *Config) {
		a.Port = p
	}
}

func WithHost(h string) OptFunc {
	return func(a *Config) {
		a.Host = h
	}
}

func WithSecret(s string) OptFunc {
	return func(a *Config) {
		a.Secret = s
	}
}

func WithLogLevel(l string) OptFunc {
	return func(a *Config) {
		a.LogLevel = l
	}
}

func WithMode(m string) OptFunc {
	return func(a *Config) {
		a.Mode = m
	}
}

type SQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type DbOptFunc func(a *SQLConfig)

func WithH(h string) DbOptFunc {
	return func(a *SQLConfig) {
		a.Host = h
	}
}
func WithP(p int) DbOptFunc {
	return func(a *SQLConfig) {
		a.Port = p
	}
}
func WithU(u string) DbOptFunc {
	return func(a *SQLConfig) {
		a.Username = u
	}
}
func WithPass(p string) DbOptFunc {
	return func(a *SQLConfig) {
		a.Password = p
	}
}
func WithD(d string) DbOptFunc {
	return func(a *SQLConfig) {
		a.Database = d
	}
}