package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	if 1 == 2 {
		return errors.New("random error")
	}
	return nil
}
