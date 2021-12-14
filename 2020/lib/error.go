package lib

import (
	"log"
)

// Check causes the program to terminate if an error is found
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
