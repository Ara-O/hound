package utils

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

func GetEnvironmentVariablePath() string {
	// TODO: Make it a fixed directory so that the code will work regardless of
	//the director the user is in, also do it in db.go
	hd, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Error getting user home directory")
	}

	folderName := path.Join(hd, "hound")

	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			log.Fatal("Error creating directory:", err)
		}
	}

	return filepath.Join(hd, "hound", ".env")

}
