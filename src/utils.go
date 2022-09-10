package util

import (
	"log"
)

func CheckErr(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func JoinStrings(sep string, str ...string) string {
	resp := ""
	for _, s := range str {
		resp += s + sep
	}
	return resp
}
