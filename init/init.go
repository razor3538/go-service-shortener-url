package vars

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

type env struct {
	Address  string
	FilePath string
}

type flags struct {
	Address  *string
	FilePath *string
	BaseURL  *string
}

var Env env
var Flag flags

func init() {
	_ = godotenv.Load()
	Env = env{
		Address:  os.Getenv("SERVER_ADDRESS"),
		FilePath: os.Getenv("FILE_STORAGE_PATH"),
	}
	Flag = flags{
		Address:  flag.String("a", "localhost:8080", "Server name"),
		FilePath: flag.String("f", "tmp", "File path"),
		BaseURL:  flag.String("b", "", "Base url dir"),
	}
}
