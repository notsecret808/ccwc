package counters

import (
	"bufio"
	"io"
)

func CountBytes(r io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(r)

	fileScanner.Split(bufio.ScanRunes)

	counter := 0

	for fileScanner.Scan() {
		counter += len(fileScanner.Bytes())
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}

func CountChars(r io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(r)

	fileScanner.Split(bufio.ScanRunes)

	counter := 0

	for fileScanner.Scan() {
		counter++
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}

func CountLines(r io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(r)

	counter := 0

	for fileScanner.Scan() {
		counter++
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}

func CountWords(r io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(r)

	fileScanner.Split(bufio.ScanWords)

	counter := 0

	for fileScanner.Scan() {
		counter++
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}
