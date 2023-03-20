package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/christian-gama/rename/renamer"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: go run main.go <path> <old_name> <new_name> <replace_inside_files>")
		os.Exit(1)
	}

	path := os.Args[1]
	oldName := os.Args[2]
	newName := os.Args[3]
	replaceInsideFiles := os.Args[4] == "true"

	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(info.Name(), oldName) {
			newFileName := renamer.ReplaceWithCase(info.Name(), oldName, newName)
			newPath := filepath.Join(filepath.Dir(currentPath), newFileName)

			err := os.Rename(currentPath, newPath)
			if err != nil {
				return err
			}

			currentPath = newPath
		}

		if replaceInsideFiles && !info.IsDir() {
			content, err := ioutil.ReadFile(currentPath)
			if err != nil {
				return err
			}

			newContent := renamer.ReplaceWithCase(string(content), oldName, newName)
			err = ioutil.WriteFile(currentPath, []byte(newContent), info.Mode())
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
