package utils

import (
	"os"
)

const defaultContent = `{
  "data": {
    "1": {
      "id": "1",
      "text": "Hello, World!"
    }
  },
  "current_id": 1
}`

func Open(file string) (*os.File, error) {
	return os.Open(file)
}

func CreateIfNotExists(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		_, err := file.WriteString(defaultContent)
		if err != nil {
			return err
		}
	}

	return nil
}
