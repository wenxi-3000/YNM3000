package utils

import (
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
)

//将文件中的~/xx 替换为绝对路径
func NormalizePath(path string) string {
	if strings.HasPrefix(path, "~") {
		path, _ = homedir.Expand(path)
	}
	return path
}

//判断文件夹是否存在
func FolderExists(foldername string) bool {
	foldername = NormalizePath(foldername)
	if _, err := os.Stat(foldername); os.IsNotExist(err) {
		return false
	}
	return true
}

// 创建一个文件夹
func MakeDir(folder string) {
	folder = NormalizePath(folder)
	os.MkdirAll(folder, 0750)
}
