package core

import (
	"YNM3000/logger"
	script "YNM3000/scripts"
	"YNM3000/utils"

	"github.com/robertkrimen/otto"
)

const (
	Append       = "Append"
	Cleaning     = "Cleaning"
	ExecCmd      = "ExecCmd"
	CreateFolder = "CreateFolder"
	DeleteFile   = "DeleteFile"
)

// InitVM init scripting engine
func (r *Runner) InitVM() {
	r.VM = otto.New()
	r.LoadEngineScripts()
}

func (r *Runner) ExecScript(script string) string {
	logger.Info(script)
	value, err := r.VM.Run(script)
	if err == nil {
		out, nerr := value.ToString()
		if nerr == nil {
			return out
		}
	}

	return ""
}

func (r *Runner) LoadEngineScripts() {
	r.LoadScripts()
	// r.LoadDBScripts()
	// r.LoadImportScripts()
	// r.LoadExternalScripts()
	// r.LoadGitScripts()
	// r.LoadNotiScripts()
}

func (r *Runner) LoadScripts() string {
	var output string
	vm := r.VM

	vm.Set(Append, func(call otto.FunctionCall) otto.Value {
		dest := call.Argument(0).String()
		src := call.Argument(1).String()
		script.Append(dest, src)
		returnValue, _ := otto.ToValue(true)
		return returnValue
	})

	vm.Set(Cleaning, func(call otto.FunctionCall) otto.Value {
		if !r.Clean {
			return otto.Value{}
		}
		script.Cleaning(call.Argument(0).String(), r.Reports)
		return otto.Value{}
	})

	// ExecCmd execute command
	vm.Set(ExecCmd, func(call otto.FunctionCall) otto.Value {
		cmd := call.Argument(0).String()
		_, err := utils.RunCommandWithErr(cmd)
		var validate bool
		if err != nil {
			validate = true
		}
		result, err := vm.ToValue(validate)
		if err != nil {
			return otto.Value{}
		}
		return result
	})

	vm.Set(CreateFolder, func(call otto.FunctionCall) otto.Value {
		utils.MakeDir(call.Argument(0).String())
		return otto.Value{}
	})

	vm.Set(DeleteFile, func(call otto.FunctionCall) otto.Value {
		script.DeleteFile(call.Argument(0).String())
		return otto.Value{}
	})

	r.VM = vm

	return output
}
