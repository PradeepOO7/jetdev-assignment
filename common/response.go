package common

import (
	"bytes"
	"io"
	"os"
)

type Respone struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    any    `json:"data"`
}

func ReadFromJSONFile(path string) (*bytes.Buffer, error) {
	fileHandler, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer fileHandler.Close()

	byteValue, err := io.ReadAll(fileHandler)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(byteValue), nil
}