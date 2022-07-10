package libs

import (
	"YNM3000/utils"
	"log"
)

func InitOptions(opt *Options) {
	//设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	opt.Paths = InitPath(opt)

}

func InitPath(opt *Options) Paths {
	var paths Paths
	//判断文件是否存在
	if !utils.FolderExists(opt.Input.ResultPath) {
		utils.MakeDir(opt.Input.ResultPath)
	}

	//结果文件
	paths.Result = opt.Input.ResultPath
	return paths
}
