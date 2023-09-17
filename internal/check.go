package internal

import "log"

func Check(err error, message string) {
	if err != nil {
		log.Fatalf("%s, %v\n", message, err)
	}
}
