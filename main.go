package main

import (
	"log"

	"github.com/wookets/go-cli-example/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
