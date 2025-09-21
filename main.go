package main

import (
	cmdParser "notsecret808/ccwc/cmd-parser"
	"notsecret808/ccwc/stream"
)

func main() {
	filePaths, options := cmdParser.ParseCommandParams()
	streams := make([]stream.Stream, 1)

	if len(filePaths) == 0 {
		s := stream.Stream{}
		s.ReadFromStdIn(options)
		streams = append(streams, s)
	} else {
		for _, filePath := range filePaths {
			s := stream.Stream{}
			s.ReadFromFile(filePath, options)
			streams = append(streams, s)
		}
	}

	for _, stream := range streams {
		stream.WriteOutput()
	}
}

func calcTotal(streams []stream.Stream) {
}
