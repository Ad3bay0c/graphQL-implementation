package main

import (
	"github.com/Ad3bay0c/graphqlTesting/application"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	application.Start()
}