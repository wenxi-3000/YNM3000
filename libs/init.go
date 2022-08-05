package libs

import (
	"YNM3000/logger"
	"YNM3000/utils"
	"path"
	"path/filepath"
)

func InitOptions(opt *Options) {
	//日志初始化
	// loggerInit(opt)

	//路径初始化
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

	//确定org目录
	if opt.Org == "" {
		paths.Org = path.Join(resultPath, "no-org")
	} else {
		paths.Org = path.Join(resultPath, "org", opt.Org)
	}
	if !utils.FolderExists(paths.Org) {
		utils.MakeDir(paths.Org)
	}
	logger.Println("文件保存路径: ", paths.Org)

	return paths
}

func InitInput(opt *Options) map[string]struct{} {
	inputs := make(map[string]struct{})
	if opt.CmdInput.InputFile != "" {
		var err error
		if inputs, err = utils.FileSet(opt.CmdInput.InputFile); err != nil {
			logger.Error(err)
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
