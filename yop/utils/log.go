package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "yop-go-sdk: ", log.LstdFlags)
