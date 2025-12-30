package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/image/webp"
	"image"
	"os"
	"path/filepath"
	"strings"
)

func ConvertToBytes(base64Str string) ([]byte, string, error) {

	if base64Str == "" {
		return nil, "", errors.New("la cadena base64 está vacía")
	}

	data, err := base64.StdEncoding.DecodeString(base64Str)

	if err != nil {
		return nil, "", fmt.Errorf("error al decodificar la cadena base64: %w", err)
	}

	reader := bytes.NewReader(data)

	if _, format, err := image.DecodeConfig(reader); err == nil {
		switch format {
		case "png":
			return data, ".png", nil
		case "jpeg":
			return data, ".jpg", nil
		default:
		}
	}

	reader.Seek(0, 0)

	if _, err := webp.DecodeConfig(reader); err == nil {
		return data, ".webp", nil
	}

	return nil, "", errors.New("formato de imagen no permitido: solo PNG, JPG o WEBP")

}

func SaveImageToDirectory(file []byte, name string, ext string, folderEnv string) (string, error) {

	if len(file) == 0 {
		return "", fmt.Errorf("el archivo está vacío")
	}

	var folder string

	switch folderEnv {
	case "images":
		folder = os.Getenv("Images")
	case "pdf":
		folder = os.Getenv("Videos")
	default:
		return "", fmt.Errorf("el tipo de carpeta especificado (%s) no es válido", folderEnv)
	}

	root := os.Getenv("Directory_Path")

	dir := filepath.Join(root, folder)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("no se pudo crear el directorio %s: %w", dir, err)
	}

	fileName := name + ext

	filePath := filepath.Join(dir, fileName)

	if err := os.WriteFile(filePath, file, 0644); err != nil {
		return "", fmt.Errorf("error al guardar el archivo: %w", err)
	}

	baseURL := strings.TrimRight(os.Getenv("Url"), "/")

	publicFolder := strings.Trim(folder, "/")

	publicURL := fmt.Sprintf("%s/%s/%s", baseURL, publicFolder, fileName)

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
