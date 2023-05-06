package config

import (
	"encoding/json"
	"example.com/m/v2/internal/tools"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"os"
	"strconv"
)

// env Структура для хранения переменных среды
type env struct {
	Address      string `json:"server_address"`
	FilePath     string `json:"file_storage_path"`
	BaseURL      string `json:"base_url"`
	BdConnection string `json:"database_dsn"`
	EnableHttps  bool   `json:"enable_https"`
}

// Env глобальная переменная для доступа к переменным среды
var Env env

// CheckFlagEnv Метод проверяющий флаги
func CheckFlagEnv() {
	var address string
	var filePath string
	var basePath string
	var dbConnection string
	var enableHttps bool
	var configFile string

	err := godotenv.Load()

	if err != nil {
		tools.ErrorLog.Println(err)
	}

	if os.Getenv("CONFIG") != "" {
		configFile = os.Getenv("CONFIG")
	} else {
		configFile = ""
	}

	var flagConfigFile = flag.String("c", "", "Path to config file")
	var flagAddress = flag.String("a", "", "Server name")
	var flagFilePath = flag.String("f", "", "File path")
	var flagBaseURL = flag.String("b", "", "Base url dir")
	var flagDSN = flag.String("d", "", "Base dsn connection")
	var flagHttps = flag.Bool("s", false, "Enable TLS connection")

	flag.Parse()

	if *flagConfigFile != "" {

		configFile = *flagConfigFile
	}

	if configFile != "" {
		jsonFile, errJson := os.Open(configFile)
		if errJson != nil {
			fmt.Println(err)
		}
		byteValue, _ := io.ReadAll(jsonFile)

		var envJson env

		errJson = json.Unmarshal(byteValue, &envJson)
		if errJson != nil {
			return
		}

		enableHttps = envJson.EnableHttps
		address = envJson.Address
		filePath = envJson.FilePath
		basePath = envJson.BaseURL
		enableHttps = envJson.EnableHttps

		defer func(jsonFile *os.File) {
			errJson = jsonFile.Close()
			if err != nil {
				return
			}
		}(jsonFile)
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

	if os.Getenv("DATABASE_DSN") != "" {
		https, err := strconv.ParseBool(os.Getenv("ENABLE_HTTPS"))
		if err != nil {
			return
		}
		if https {
			enableHttps = https
		} else {
			enableHttps = false
		}
	}

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

	if *flagHttps {

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
