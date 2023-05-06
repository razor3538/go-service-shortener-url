package config

import (
	"flag"
	"os"

	"example.com/m/v2/internal/tools"
	"github.com/joho/godotenv"
)

// env Структура для хранения переменных среды
type env struct {
	Address      string
	FilePath     string
	BaseURL      string
	BdConnection string
	EnableHttps  string
}

// Env глобальная переменная для доступа к переменным среды
var Env env

// CheckFlagEnv Метод проверяющий флаги
func CheckFlagEnv() {
	var address string
	var filePath string
	var basePath string
	var dbConnection string
	var enableHttps string

	err := godotenv.Load()

	if err != nil {
		tools.ErrorLog.Println(err)
	}

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
		dbConnection = os.Getenv("DATABASE_DSN")
	} else {
		dbConnection = ""
	}

	if os.Getenv("ENABLE_HTTPS") != "" {
		enableHttps = os.Getenv("ENABLE_HTTPS")
	} else {
		enableHttps = ""
	}

	var flagAddress = flag.String("a", "", "Server name")
	var flagFilePath = flag.String("f", "", "File path")
	var flagBaseURL = flag.String("b", "", "Base url dir")
	var flagDSN = flag.String("d", "", "Base dsn connection")
	var flagHttps = flag.String("s", "", "Enable TLS connection")

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

		dbConnection = *flagDSN
	}

	if *flagHttps != "" {

		enableHttps = *flagHttps
	}

	Env = env{
		Address:      address,
		FilePath:     filePath,
		BaseURL:      basePath,
		BdConnection: dbConnection,
		EnableHttps:  enableHttps,
	}
}
