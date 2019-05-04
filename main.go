package main

import (
	"log"
	"os"

	"spottyci/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
