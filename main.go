package main

import (
	"fmt"
	"io"
	"log"
	"notsecret808/ccwc/counters"
	"notsecret808/ccwc/utils"
	"os"
)

func main() {
	options, path, err := utils.ParseCmdParams()

	if err != nil {
		log.Fatal(err)
	}

	readFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	f := Stream{Reader: readFile}

	if len(options) == 0 {
		f.countBytes()
		f.countLines()
		f.countWords()
	}

	for _, option := range options {
		switch option {
		case "-c", "--bytes":
			f.countBytes()
		case "-m", "--chars":
			f.countChars()
		case "-l", "--lines":
			f.countLines()
		case "-w", "--words":
			f.countWords()
		default:
			log.Fatal("Incorrect option")
		}
	}

	f.writeOutput()
}

type Stream struct {
	Reader   io.Reader
	FilePath string
	Bytes    *int
	Chars    *int
	Lines    *int
	Words    *int
}

func (f *Stream) countBytes() {
	bytesCount, error := counters.CountBytes(f.Reader)

	if error != nil {
		log.Fatal(error)
	}

	if f.Bytes == nil {
		f.Bytes = new(int)
	}

	*f.Bytes = bytesCount
}

func (f *Stream) countChars() {
	charsCount, error := counters.CountChars(f.Reader)

	if error != nil {
		log.Fatal(error)
	}

	if f.Chars == nil {
		f.Chars = new(int)
	}

	*f.Chars = charsCount

}

func (f *Stream) countLines() {
	linesCount, error := counters.CountLines(f.Reader)

	if error != nil {
		log.Fatal(error)
	}

	if f.Lines == nil {
		f.Lines = new(int)
	}

	*f.Lines = linesCount
}

func (f *Stream) countWords() {
	wordsCount, error := counters.CountWords(f.Reader)

	if error != nil {
		log.Fatal(error)
	}

	if f.Words == nil {
		f.Words = new(int)
	}

	*f.Words = wordsCount
}

func (f *Stream) writeOutput() {
	output := ""

	var values []*int

	values = append(values, f.Lines)
	values = append(values, f.Words)
	values = append(values, f.Bytes)
	values = append(values, f.Chars)

	for _, value := range values {
		if value == nil {
			continue
		}

		if output == "" {
			output = fmt.Sprintf("%d", *value)
		} else {
			output = fmt.Sprintf("%s %d", output, *value)
		}
	}

	fmt.Printf("%s %s\n", output, f.Reader)
}
