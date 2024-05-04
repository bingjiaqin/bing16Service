package file

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func WriteFile(path string, data string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
		return errors.New("write file error")
	}

	return nil
}

func ReadJson[T any](path string, ignore string) ([]T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
		return nil, err
	}

	var result []T
	err = json.Unmarshal(data[len(ignore):], &result)
	if err != nil {
		log.Fatalf("无法解析JSON: %v", err)
		return nil, err
	}

	return result, nil
}

func WriteJson[T any](path string, prefix string, content T) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(content)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	WriteFile(path, prefix+buffer.String())
}
