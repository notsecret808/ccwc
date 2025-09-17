package errors

import "fmt"

type FileNotExistError struct {
	Path string
}

func (e *FileNotExistError) Error() string {
	return fmt.Sprintf("File does not exits: %s", e.Path)
}
