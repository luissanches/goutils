package goutils

import "fmt"

func GetName(name string) string {
	newName := fmt.Sprintf("Bom dia %s", name)
	return newName
}
