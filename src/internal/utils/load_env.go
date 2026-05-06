package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	for _, path := range []string{
		".env",
	} {
		if _, err := os.Stat(path); err == nil {
			return godotenv.Load(path)
		} else if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	return nil
}
