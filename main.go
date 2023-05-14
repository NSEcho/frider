package main

import (
	"fmt"
	"github.com/nsecho/frider/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error ocurred: %v\n", err)
		os.Exit(1)
	}
}
