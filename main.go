package main

import (
	"fmt"
	"notsecret808/ccwc/assets"
	cmdParser "notsecret808/ccwc/cmd-parser"
	"notsecret808/ccwc/stream"
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
			helpPagePath := filepath.Join("data", "help-page.txt")
			data, err := assets.Files.ReadFile(helpPagePath)

			if err != nil {
				panic(err)
			}

			return string(data)
		case "--version", "-v":
			versionPath := filepath.Join("data", "version.txt")
			data, err := assets.Files.ReadFile(versionPath)

			if err != nil {
				panic(err)
			}

			return string(data)
		}
	}

	return ""
}
