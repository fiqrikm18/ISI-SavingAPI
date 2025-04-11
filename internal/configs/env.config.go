package configs

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var DBHost string = ""
var DBName string = ""
var DBUsername string = ""
var DBPassword string = ""
var DBEnableSSL bool = false
var AppPort string = ":8080"

func init() {
	parseEnv()

	// get port from command line arguments
	// default port is 8080
	// if port is not provided, use the default port
	// if port is provided, use the provided port
	// if port is less than 1 or greater than 65535, return error
	portParam := flag.Int("port", 8080, "Port to run the server on")
	if *portParam != 8080 {
		AppPort = fmt.Sprintf(":%d", *portParam)
	}

	if *portParam < 1 || *portParam > 65535 {
		panic("Port must be between 1 and 65535")
	}
}

func parseEnv() {
	enableSSl, _ := strconv.ParseBool(os.Getenv("DB_ENABLE_SSL"))

	DBHost = os.Getenv("DB_HOST")
	DBName = os.Getenv("DB_NAME")
	DBUsername = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBEnableSSL = enableSSl
}
