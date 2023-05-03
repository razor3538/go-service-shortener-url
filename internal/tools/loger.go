package tools

import (
	"log"
	"os"
)

// InfoLog global var that give a info logger methods
var InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

// ErrorLog global var that give a error logger methods
var ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
