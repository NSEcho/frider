package main

import (
	"github.com/apex/log"
	"github.com/nsecho/frider/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error ocurred: %v", err)
	}
}
