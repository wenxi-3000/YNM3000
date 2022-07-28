package utils

import (
	"bufio"
	"io/ioutil"
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

//判断文件是否存在
func FileExists(filename string) bool {
	filename = NormalizePath(filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 创建一个文件夹
func MakeDir(folder string) {
	folder = NormalizePath(folder)
	os.MkdirAll(folder, 0750)
}

//按行读取一个文件，返回map[string]{}{}
func FileSet(path string) (map[string]struct{}, error) {
	results := map[string]struct{}{}
	if strings.HasPrefix(path, "~") {
		path, _ = homedir.Expand(path)
	}
	file, err := os.Open(path)
	if err != nil {
		return results, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		results[line] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}

	return results, nil
}

// FileLength count len of file
func FileLength(filename string) int {
	filename = NormalizePath(filename)
	if !FileExists(filename) {
		return 0
	}
	return CountLines(filename)
}

// CountLines Return the lines amount of the file
func CountLines(filename string) int {
	var amount int
	if strings.HasPrefix(filename, "~") {
		filename, _ = homedir.Expand(filename)
	}
	file, err := os.Open(filename)
	if err != nil {
		return amount
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := strings.TrimSpace(scanner.Text())
		if val == "" {
			continue
		}
		amount++
	}

	if err := scanner.Err(); err != nil {
		return amount
	}
	return amount
}

// GetFileContent Reading file and return content of it
func GetFileContent(filename string) string {
	var result string
	if strings.Contains(filename, "~") {
		filename, _ = homedir.Expand(filename)
	}
	file, err := os.Open(filename)
	if err != nil {
		return result
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return result
	}
	return string(b)
}

// AppendToContent append string to a file
func AppendToContent(filename string, data string) (string, error) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	if _, err := f.Write([]byte(data)); err != nil {
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}
	return filename, nil
}
