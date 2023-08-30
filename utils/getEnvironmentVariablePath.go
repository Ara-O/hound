package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetEnvironmentVariablePath() string {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal("Error logging working directory")
	}

	return filepath.Join(wd, ".env")

}
