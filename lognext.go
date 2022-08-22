package main

import (
	"log"
	"os"
)

func main() {
	oldLocation := "/var/log/sys.log"
	newLocation := "/writelog/cp/sys.log"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
