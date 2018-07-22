package main

import (
	"fmt"
	"os"

	"github.com/object88/writing-compilers-and-interpreters/cmd"
)

func main() {
	rootCmd := cmd.InitializeCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
