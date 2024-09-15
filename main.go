package main

import (
	"fmt"
	"log"

	"github.com/sriram-kailasam/equalize/config"
)

func main() {
	fmt.Print("starting equalize")

	filename := "./conf.yaml"

	conf, err := config.Parse(filename)
	if err != nil {
		log.Fatal(err)
	}

	startServers(conf)
	for {
	}
}
