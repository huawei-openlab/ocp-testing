package utils

import (
	"os"
)

func FileExist(filename string) bool {
	stat, err := os.Stat(filename)
	return err == nil || os.IsExist(err) || (!stat.IsDir())
}

func DirExist(path string) bool {
	stat, err := os.Stat(path)
	return err == nil || os.IsExist(err) || stat.IsDir()
}

func RemoveFileIfExist(path string) {
	if _, err := os.Stat(path); err == nil {
		os.Remove(path)
	}
}
