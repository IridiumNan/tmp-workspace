package utils

import (
	"fmt"
	"os"
)

var HomePath, _ = os.UserHomeDir()

func EnsureExist(TargetPath string) (err error) {
	if _, err = os.Stat(TargetPath); err != nil {
		fmt.Println("path -> ", TargetPath, " not exist, try to create")

		err = os.MkdirAll(TargetPath, 0o755)
		if err != nil {
			return
		}

		return nil
	}

	return
}
