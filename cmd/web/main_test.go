package main

import (
	"log"
	"testing"
)

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		log.Fatal("failed run()")
	}
}
