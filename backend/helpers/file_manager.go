package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func SaveImageToDirectory(file []byte, name string, ext string, folder string) (string, error) {

	if len(file) == 0 {
		return "", fmt.Errorf("el archivo está vacío")
	}

	root := os.Getenv("Directory_Path")
	if root == "" {
		return "", fmt.Errorf("Directory_Path no definido")
	}

	dir := filepath.Join(root, folder)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	fileName := name + ext
	filePath := filepath.Join(dir, fileName)

	if err := os.WriteFile(filePath, file, 0644); err != nil {
		return "", err
	}

	baseURL := strings.TrimRight(os.Getenv("Url"), "/")
	publicURL := fmt.Sprintf("%s/%s/%s", baseURL, folder, fileName)

	return publicURL, nil
}

func DeleteImageFromDirectory(fileName string, folder string) error {

	root := os.Getenv("Directory_Path")

	filePath := filepath.Join(root, folder, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("el archivo %s no existe en la ruta especificada", filePath)
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("error al eliminar el archivo %s: %w", filePath, err)
	}

	return nil
}
