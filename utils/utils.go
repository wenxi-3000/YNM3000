package utils

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

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

// CalcTimeout calculate timeout
func CalcTimeout(raw string) int {
	raw = strings.ToLower(strings.TrimSpace(raw))
	seconds := raw
	multiply := 1

	matched, _ := regexp.MatchString(`.*[a-z]`, raw)
	if matched {
		unitTime := fmt.Sprintf("%c", raw[len(raw)-1])
		seconds = raw[:len(raw)-1]
		switch unitTime {
		case "s":
			multiply = 1
			break
		case "m":
			multiply = 60
			break
		case "h":
			multiply = 3600
			break
		}
	}

	timeout, err := strconv.Atoi(seconds)
	if err != nil {
		return 0
	}
	return timeout * multiply
}

// RunCommandWithErr Run a command
func RunCommandWithErr(command string, timeoutRaw ...string) (string, error) {
	if len(timeoutRaw) == 0 {
		return runCommandWithError(command)
	}
	var output string
	var err error

	timeout := CalcTimeout(timeoutRaw[0])
	log.Println("Run command with %v seconds timeout", timeout)
	var out string

	c := context.Background()
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	c, cancel := context.WithDeadline(c, deadline)
	defer cancel()
	go func() {
		out, err = runCommandWithError(command)
		cancel()
	}()

	select {
	case <-c.Done():
		return output, err
	case <-time.After(time.Duration(timeout) * time.Second):
		return out, fmt.Errorf("command got timeout")
	}
}

func runCommandWithError(cmd string) (string, error) {
	log.Println(cmd)
	command := []string{
		"bash",
		"-c",
		cmd,
	}
	var output string
	realCmd := exec.Command(command[0], command[1:]...)

	// output command output to std too
	cmdReader, _ := realCmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			out := scanner.Text()
			log.Println(out)
			output += out + "\n"
		}
	}()
	if err := realCmd.Start(); err != nil {
		return output, err
	}
	if err := realCmd.Wait(); err != nil {
		return output, err
	}
	return output, nil
}
