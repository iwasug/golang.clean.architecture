package main

import (
	"log"

	"golang.clean.architecture/api"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	api.Init()
}
