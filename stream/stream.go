package stream

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"notsecret808/ccwc/counters"
	"os"
)

type FileStream struct {
	FilePath string
	Bytes    *int
	Chars    *int
	Lines    *int
	Words    *int
}

func (f *FileStream) countOptions(filePath string, options []string) {

	if len(options) == 0 {
		f.countBytes(filePath)
		f.countLines(filePath)
		f.countWords(filePath)
	}

	for _, option := range options {
		switch option {
		case "-c", "--bytes":
			f.countBytes(filePath)
		case "-m", "--chars":
			f.countChars(filePath)
		case "-l", "--lines":
			f.countLines(filePath)
		case "-w", "--words":
			f.countWords(filePath)
		default:
			message := fmt.Sprintf("invalid option: %s \n", option)
			log.Fatal(message)
		}
	}
}

func (f *FileStream) ReadFromFile(path string, options []string) {
	f.countOptions(path, options)
	f.FilePath = path
}

func (f *FileStream) ReadFromStdIn(options []string) {
	stdin := bufio.NewReader(os.Stdin)
	tempFile, tempErr := os.CreateTemp("", "stdin")

	if tempErr != nil {
		log.Print(tempErr)
		return
	}

	_, copyError := io.Copy(tempFile, stdin)

	if copyError != nil {
		log.Print(copyError)
		return
	}

	tempFile.Close()

	f.countOptions(tempFile.Name(), options)

	defer os.Remove(tempFile.Name())
}

func (f *FileStream) countBytes(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Print(err)
		return
	}

	bytesCount, error := counters.CountBytes(file)

	if error != nil {
		log.Fatal(error)
	}

	if f.Bytes == nil {
		f.Bytes = new(int)
	}

	*f.Bytes = bytesCount

	defer file.Close()
}

func (f *FileStream) countChars(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Print(err)
		return
	}

	charsCount, error := counters.CountChars(file)

	if error != nil {
		log.Fatal(error)
	}

	if f.Chars == nil {
		f.Chars = new(int)
	}

	*f.Chars = charsCount

	defer file.Close()
}

func (f *FileStream) countLines(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Print(err)
		return
	}

	linesCount, error := counters.CountLines(file)

	if error != nil {
		log.Fatal(error)
		return
	}

	if f.Lines == nil {
		f.Lines = new(int)
	}

	*f.Lines = linesCount

	defer file.Close()
}

func (f *FileStream) countWords(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Print(err)
		return
	}

	wordsCount, error := counters.CountWords(file)

	if error != nil {
		log.Fatal(error)
	}

	if f.Words == nil {
		f.Words = new(int)
	}

	*f.Words = wordsCount

	defer file.Close()
}

func (f *FileStream) WriteOutput() {
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

	fmt.Printf("%s %s\n", output, f.FilePath)
}
