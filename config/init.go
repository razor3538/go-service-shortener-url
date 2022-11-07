package config

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

type env struct {
	Address  string
	FilePath string
	BaseURL  string
}

var Env env

func CheckFlagEnv() {
	var address string
	var filePath string
	var basePath string

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

	var flagAddress = flag.String("a", "", "Server name")
	var flagFilePath = flag.String("f", "", "File path")
	var flagBaseUrl = flag.String("b", "", "Base url dir")

	flag.Parse()

	if *flagAddress != "" {
		address = *flagAddress
	}

	if *flagFilePath != "" {
		filePath = *flagFilePath
	}

	if *flagBaseUrl != "" {
		basePath = *flagBaseUrl
	}

	Env = env{
		Address:  address,
		FilePath: filePath,
		BaseURL:  basePath,
	}
}
