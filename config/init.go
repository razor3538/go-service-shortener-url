package config

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

type env struct {
	Address      string
	FilePath     string
	BaseURL      string
	BdConnection string
}

var Env env

func CheckFlagEnv() {
	var address string
	var filePath string
	var basePath string
	var dbConnection string

	_ = godotenv.Load()

	if os.Getenv("SERVER_ADDRESS") != "" {
		address = os.Getenv("SERVER_ADDRESS")
	} else {
		address = "localhost:8080"
	}

	if os.Getenv("FILE_STORAGE_PATH") != "" {
		filePath = os.Getenv("FILE_STORAGE_PATH")
	} else {
		filePath = ""
	}

	if os.Getenv("DATABASE_DSN") != "" {
		println("получил из енви -> " + os.Getenv("DATABASE_DSN"))

		dbConnection = os.Getenv("DATABASE_DSN")
	} else {
		println("не получил из енви")
		dbConnection = ""
	}

	var flagAddress = flag.String("a", "", "Server name")
	var flagFilePath = flag.String("f", "", "File path")
	var flagBaseURL = flag.String("b", "", "Base url dir")
	var flagDSN = flag.String("d", "", "Base dsn connection")

	flag.Parse()

	if *flagAddress != "" {
		address = *flagAddress
	}

	if *flagFilePath != "" {
		filePath = *flagFilePath
	}

	if *flagBaseURL != "" {
		basePath = *flagBaseURL
	}

	if *flagDSN != "" {
		println("получил из флага -> " + *flagDSN)

		dbConnection = *flagDSN
	} else {
		println("не получил из флага")
	}

	Env = env{
		Address:      address,
		FilePath:     filePath,
		BaseURL:      basePath,
		BdConnection: dbConnection,
	}
}
