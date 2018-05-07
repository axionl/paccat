package error

import (
	"log"
)

func CheckErr (msg string, err error) {
	if err != nil{
		log.Fatal(msg, err)
	}
}