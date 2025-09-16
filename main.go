package main

import (
	"fmt"
	"log"
	"os"
)

type Command struct {
	Option string
	Path   string
}

func main() {
	cmd := parseArgs()

	switch cmd.Option {
	case "-c":
		size := getNumberOfBytes(cmd.Path)
		fmt.Printf("Number of bytes: %d \n", size)
	}
}

func parseArgs() Command {
	args := os.Args[1:]

	return Command{args[0], args[1]}
}

func getNumberOfBytes(path string) int64 {
	data, err := os.Stat(path)

	if err != nil {
		log.Fatal(err)
	}

	return data.Size()
}
