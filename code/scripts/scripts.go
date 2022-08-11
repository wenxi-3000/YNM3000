package scripts

import (
	"YNM3000/code/logger"
	"YNM3000/code/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thoas/go-funk"
)

// func (r *Runner) LoadScripts() {
// 	//var output string
// 	vm := otto.New()
// 	logger.Info("xxxx")
// 	vm.Run(`
//     abc = 2 + 2;
//     console.log("The value of abc is " + abc); // 4
// `)
// vm.Set(Append, func(call otto.FunctionCall) otto.Value {
// 	dest := call.Argument(0).String()
// 	src := call.Argument(1).String()
// 	Append(dest, src)
// 	returnValue, _ := otto.ToValue(true)
// 	return returnValue
// })

// }

// func Append(dest string, src string) {
// 	if !utils.FileExists(src) || utils.FileLength(src) <= 0 {
// 		logger.Info("error to append %v", src)
// 	}
// 	data := utils.GetFileContent(src)
// 	utils.AppendToContent(dest, data)
// }

func Append(dest string, src string) {
	logger.Info(dest)
	logger.Info(src)
	if !utils.FileExists(src) || utils.FileLength(src) <= 0 {
		logger.Info(dest, src)
	}
	data := utils.GetFileContent(src)
	utils.AppendToContent(dest, data)
}

//Cleaning the execution directory
func Cleaning(folder string, reports []string) {
	// list all the file
	items, err := filepath.Glob(fmt.Sprintf("%v/*", folder))
	if err != nil {
		return
	}

	for _, item := range items {
		item = utils.NormalizePath(item)
		if funk.Contains(reports, item) {
			continue
		}

		fi, err := os.Stat(item)
		if err != nil {
			continue
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			DeleteFolder(item)
			continue
		case mode.IsRegular():
			DeleteFile(item)
		}
	}
}

// DeleteFolder delete entire folder
func DeleteFolder(path string) {
	os.RemoveAll(utils.NormalizePath(path))
}

// DeleteFile delete file
func DeleteFile(filename string) {
	os.Remove(utils.NormalizePath(filename))
}
