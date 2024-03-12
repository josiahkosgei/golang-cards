package main

import (
	"cards/api"
	"log"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	api.Init()
}
