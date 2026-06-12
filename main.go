package main

import (
	"fmt"
	"log"

	"github.com/itency/blog_aggregator/internal/config"
)

func main() {
	file, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = file.SetUser("Alex")
	if err != nil {
		log.Fatal(err)
	}

	file, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
}
