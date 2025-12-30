package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolders() error {

	folders := []string{
		filepath.Join(os.Getenv("Directory_Path"), os.Getenv("Images")),
		filepath.Join(os.Getenv("Directory_Path"), os.Getenv("Videos")),
	}

	for _, folder := range folders {
		err := os.MkdirAll(folder, os.ModePerm)

		if err != nil {
			return fmt.Errorf("no se pudo crear la carpeta %s: %w", folder, err)
		}
	}

	return nil
}
