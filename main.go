package main

import (
	"fmt"
	cmdParser "notsecret808/ccwc/cmd-parser"
	"notsecret808/ccwc/stream"
	"notsecret808/ccwc/utils"
	"os"
	"path/filepath"
)

func main() {
	filePaths, options := cmdParser.ParseCommandParams()
	streams := make([]stream.FileStream, 1)

	for _, option := range options {
		switch option {
		case "--help", "-h":
			root, err := utils.GetModuleRootDirectory()

			if err != nil {
				panic(err)
			}

			helpPage := filepath.Join(root, "assets", "help-page.txt")
			data, err := os.ReadFile(helpPage)

			if err != nil {
				panic(err)
			}

			fmt.Println(string(data))
			return
		case "--version", "-v":
			root, err := utils.GetModuleRootDirectory()

			if err != nil {
				panic(err)
			}

			helpPage := filepath.Join(root, "assets", "version.txt")
			data, err := os.ReadFile(helpPage)

			if err != nil {
				panic(err)
			}

			fmt.Println(string(data))
			return

		}
	}

	if len(filePaths) == 0 {
		s := stream.FileStream{}
		s.ReadFromStdIn(options)
		streams = append(streams, s)
	} else {
		for _, filePath := range filePaths {
			s := stream.FileStream{}
			s.ReadFromFile(filePath, options)
			streams = append(streams, s)
		}
	}

	for _, stream := range streams {
		stream.WriteOutput()
	}
}

func calcTotal(streams []stream.FileStream) {
}
