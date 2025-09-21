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
	streams := make([]stream.FileStream, 0)

	if data := getStaticFlagContent(options); data != "" {
		fmt.Println(data)
		return
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

	if len(streams) > 1 {
		countTotal(streams)
	}
}

func countTotal(streams []stream.FileStream) {
	fileStream := stream.FileStream{}

	for _, s := range streams {
		if s.Bytes != nil {
			if fileStream.Bytes == nil {
				fileStream.Bytes = new(int)
			}

			*fileStream.Bytes += *s.Bytes
		}
		if s.Chars != nil {
			if fileStream.Chars == nil {
				fileStream.Chars = new(int)
			}

			*fileStream.Chars += *s.Chars
		}
		if s.Lines != nil {
			if fileStream.Lines == nil {
				fileStream.Lines = new(int)
			}

			*fileStream.Lines += *s.Lines
		}
		if s.Words != nil {
			if fileStream.Words == nil {
				fileStream.Words = new(int)
			}

			*fileStream.Words += *s.Words
		}
	}

	fileStream.FilePath = "total"

	fileStream.WriteOutput()
}

func getStaticFlagContent(options []string) (content string) {
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

			return string(data)
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

			return string(data)
		}
	}

	return ""
}
