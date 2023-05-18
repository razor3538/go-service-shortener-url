package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"example.com/m/v2/internal/tools"
	"github.com/joho/godotenv"
)

// env Структура для хранения переменных среды
type env struct {
	Address       string `json:"server_address"`
	Pem           string `json:"pem"`
	Key           string `json:"key"`
	FilePath      string `json:"file_storage_path"`
	BaseURL       string `json:"base_url"`
	BdConnection  string `json:"database_dsn"`
	EnableHTTPS   bool   `json:"enable_https"`
	TrustedSubnet string `json:"trusted_subnet"`
}

// Env глобальная переменная для доступа к переменным среды
var Env env

// CheckFlagEnv Метод проверяющий флаги
func CheckFlagEnv() {
	var address string
	var filePath string
	var basePath string
	var dbConnection string
	var pem string
	var key string
	var enableHTTPS bool
	var configFile string
	var trustedSubnet string

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
	var flagHTTPS = flag.Bool("s", false, "Enable TLS connection")
	var flagPem = flag.String("p", "", "pem")
	var flagKey = flag.String("k", "", "key")
	var flagTrustedSubnet = flag.String("t", "", "trusted subnet")

	flag.Parse()

	if *flagConfigFile != "" {
		configFile = *flagConfigFile
	}

	if configFile != "" {
		jsonFile, errJSON := os.Open(configFile)
		if errJSON != nil {
			fmt.Println(err)
		}
		byteValue, _ := io.ReadAll(jsonFile)

		var envJSON env

		errJSON = json.Unmarshal(byteValue, &envJSON)
		if errJSON != nil {
			return
		}

		enableHTTPS = envJSON.EnableHTTPS
		address = envJSON.Address
		filePath = envJSON.FilePath
		basePath = envJSON.BaseURL
		dbConnection = envJSON.BdConnection
		pem = envJSON.Pem
		key = envJSON.Key
		trustedSubnet = envJSON.TrustedSubnet

		Env = env{
			Address:       address,
			FilePath:      filePath,
			BaseURL:       basePath,
			BdConnection:  dbConnection,
			EnableHTTPS:   enableHTTPS,
			Pem:           pem,
			Key:           key,
			TrustedSubnet: trustedSubnet,
		}

		defer func(jsonFile *os.File) {
			errJSON = jsonFile.Close()
			if err != nil {
				return
			}
		}(jsonFile)
	} else {
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

		if os.Getenv("PEM") != "" {
			pem = os.Getenv("PEM")
		} else {
			pem = ""
		}

		if os.Getenv("KEY") != "" {
			key = os.Getenv("KEY")
		} else {
			key = ""
		}

		if os.Getenv("TRUSTED_SUBNET") != "" {
			trustedSubnet = os.Getenv("TRUSTED_SUBNET")
		} else {
			trustedSubnet = ""
		}

		if os.Getenv("DATABASE_DSN") != "" {
			https, err := strconv.ParseBool(os.Getenv("ENABLE_HTTPS"))
			if err != nil {
				return
			}
			if https {
				enableHTTPS = https
			} else {
				enableHTTPS = false
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

		if *flagHTTPS {

			enableHTTPS = *flagHTTPS
		}

		if *flagPem == "" {

			pem = *flagPem
		}

		if *flagKey == "" {

			key = *flagKey
		}

		if *flagTrustedSubnet == "" {

			trustedSubnet = *flagTrustedSubnet
		}

		Env = env{
			Address:       address,
			FilePath:      filePath,
			BaseURL:       basePath,
			BdConnection:  dbConnection,
			EnableHTTPS:   enableHTTPS,
			Pem:           pem,
			Key:           key,
			TrustedSubnet: trustedSubnet,
		}
	}
}
