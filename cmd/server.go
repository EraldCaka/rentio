package main

import (
	"github.com/EraldCaka/rentio/router"
	"github.com/EraldCaka/rentio/util"
	"log"
)

func main() {
	util.ReadEnvFile()
	router.NewRouter()
	if err := router.Start(":5555"); err != nil {
		log.Fatal(err)
	}

}
