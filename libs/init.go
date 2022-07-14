package libs

import (
	"YNM3000/utils"
	"log"
	"path"
	"path/filepath"
)

func InitOptions(opt *Options) {
	//设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	opt.Paths = InitPath(opt)

	//初始化命令行输入
	opt.Inputs = InitInput(opt)

}

func InitPath(opt *Options) Paths {
	var paths Paths
	//判断文件是否存在
	resultPath := utils.NormalizePath(opt.CmdInput.ResultPath)
	if !utils.FolderExists(resultPath) {
		utils.MakeDir(resultPath)
	}
	resultPath, _ = filepath.Abs(resultPath)
	paths.Result = resultPath

	//项目根目录
	rootPath := path.Dir(resultPath)
	paths.Root = rootPath
	return paths
}

func InitInput(opt *Options) map[string]struct{} {
	inputs := make(map[string]struct{})
	if opt.CmdInput.InputFile != "" {
		var err error
		if inputs, err = utils.FileSet(opt.CmdInput.InputFile); err != nil {
			log.Println(err)
		}
	}
	if opt.CmdInput.Input != "" {
		inputs[opt.CmdInput.Input] = struct{}{}
	}

	if len(opt.CmdInput.Inputs) != 0 {
		for _, input := range opt.CmdInput.Inputs {
			inputs[input] = struct{}{}
		}
	}
	return inputs
}
