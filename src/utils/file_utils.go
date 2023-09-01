package utils

import (
	"os"
)

func Open(file string) (*os.File, error) {
	return os.Open(file)
}
