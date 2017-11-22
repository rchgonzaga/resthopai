package core

import "os"

// Config struct
type Config struct {
	DBConnection string
	DBType       string
}

// Fetch from env variables
func (c *Config) Fetch() {
	c.DBConnection = os.Getenv("DB_CONNECTION")
	c.DBType = os.Getenv("DB_TYPE")
}
