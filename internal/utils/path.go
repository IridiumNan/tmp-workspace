package utils

import (
	"fmt"
	"os"
	"path"
	"time"
)

var HomePath, _ = os.UserHomeDir()

func ensureDirExist(dirPath string) (err error) {
	err = os.MkdirAll(dirPath, 0o755)

	return
}

func ensureFileExist(filePath string) error {
	file, err := os.Create(filePath)
	defer file.Close()

	return err
}

func EnsureExist(targetPath string, isDir bool) (err error) {
	if _, err = os.Stat(targetPath); err == nil {
		return
	}

	if isDir {
		fmt.Println("path -> ", targetPath, "not exist, try to create")
		err = ensureDirExist(targetPath)
		if err != nil {
			return fmt.Errorf("error when create dir -> %s, error -> %w", targetPath, err)
		}

		// ensure the dir then return
		return
	}

	// if not dir but a file

	// ensure the parent dir first
	dirPath := path.Dir(targetPath)
	err = ensureDirExist(dirPath)
	if err != nil {
		return fmt.Errorf("error when create dir -> %s, error -> %w", dirPath, err)
	}

	// then create the file
	filePath := targetPath

	// backup if the file exist
	BakIfExist(filePath, false)

	err = ensureFileExist(filePath)
	if err != nil {
		return fmt.Errorf("error when create file -> %s, error -> %w", filePath, err)
	}

	return
}

func BakIfExist(targetPath string, isDir bool) {
	_, err := os.Stat(targetPath)
	if os.IsExist(err) && isDir {
		backPath := targetPath + time.Now().String() + "-bak"
		fmt.Printf("path -> %s exist, backup it as %s\n", targetPath, backPath)
	}

	if os.IsExist(err) && !isDir {
		backPath := targetPath + time.Now().String() + ".bak"
		fmt.Printf("path -> %s exist, backup it as %s\n", targetPath, backPath)
	}
}
