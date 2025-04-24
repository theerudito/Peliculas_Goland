package helpers

import "strings"

func QuitarGuiones(texto string) string {
	return strings.ReplaceAll(texto, "-", " ")
}
