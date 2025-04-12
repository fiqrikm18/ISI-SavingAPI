package configs

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	DBHost      string
	DBName      string
	DBUsername  string
	DBPassword  string
	DBEnableSSL string

	AppPort string
	AppHost string
}

func init() {

}

func NewConfig() *Config {
	DBHost := os.Getenv("DB_HOST")
	DBName := os.Getenv("DB_NAME")
	DBUsername := os.Getenv("DB_USERNAME")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBEnableSSL := os.Getenv("DB_ENABLE_SSL")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	portParam := flag.Int("port", 8080, "Port to run the server on")
	hostPataram := flag.String("host", "localhost", "Host to run the server on")
	flag.Parse()

	AppPort := ":8080"
	AppHost := "localhost"

	if *portParam != 8080 {
		AppPort = fmt.Sprintf(":%d", *portParam)
	}

	if *portParam < 1 || *portParam > 65535 {
		panic("Port must be between 1 and 65535")
	}

	if *hostPataram != "localhost" {
		AppHost = *hostPataram
	}

	return &Config{
		DBHost:      DBHost,
		DBName:      DBName,
		DBUsername:  DBUsername,
		DBPassword:  DBPassword,
		DBEnableSSL: DBEnableSSL,
		AppPort:     AppPort,
		AppHost:     AppHost,
	}
}
