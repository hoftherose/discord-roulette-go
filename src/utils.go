package util

import (
	"log"
)

func CheckErr(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
