package modules

import (
	"log"
	"os"
	"time"
)

func init() {

	//register module to module handler
	registerFetchHandler("hammerfluxhome", HammerfluxHomeFetch)
}

func HammerfluxHomeFetch() {
	t := time.Now().UTC()
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("hammerfluxhome.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(t.String())); err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
