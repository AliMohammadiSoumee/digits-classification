package class

import (
	"net/http"

	"os"
)

func isPng(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}

	typ, err := getFileContentType(file)
	if err != nil {
		return false
	}

	if typ == "image/png" {
		return true
	}

	return false
}

func getFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
