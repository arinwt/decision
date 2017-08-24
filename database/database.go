package database

import (
	"os"
	"path"
	"path/filepath"
	"io/ioutil"
)

const choiceRoot = "choices"

// initialize choices
func Init() {
	os.MkdirAll(choiceRoot, 0700)
}

func GetAllChoices() []string {
	var files []string
	filepath.Walk(choiceRoot, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			files = append(files, path)
		}
		return err
	})
	return files
}

func GetChoice(relPath string) string {
	relPath = path.Join(choiceRoot, relPath)
	b, _ := ioutil.ReadFile(relPath)

	return string(b)
}

func SetChoice(relPath, content string) {
	relPath = path.Join(choiceRoot, relPath)
	os.MkdirAll(path.Dir(relPath), 0700)
	ioutil.WriteFile(relPath, []byte(content), 0700)
}

func DeleteChoice(relPath string) {
	relPath = path.Join(choiceRoot, relPath)
	os.Remove(relPath)
}