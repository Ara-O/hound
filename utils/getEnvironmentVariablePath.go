package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetEnvironmentVariablePath() string {
	// TODO: Make it a fixed directory so that the code will work regardless of
	//the director the user is in, also do it in db.go
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal("Error logging working directory")
	}

	return filepath.Join(wd, ".env")

}
