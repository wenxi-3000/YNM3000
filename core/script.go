package core

import (
	script "YNM3000/scripts"
	"log"

	"github.com/robertkrimen/otto"
)

const (
	Append = "Append"
)

// InitVM init scripting engine
func (r *Runner) InitVM() {
	r.VM = otto.New()
	r.LoadEngineScripts()
}

func (r *Runner) ExecScript(script string) string {
	log.Println("[Run-Scripts] %v", script)
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

	// set attribute
	vm.Set("Target", r.Target)

	vm.Set(Append, func(call otto.FunctionCall) otto.Value {
		dest := call.Argument(0).String()
		src := call.Argument(1).String()
		script.Append(dest, src)
		returnValue, _ := otto.ToValue(true)
		return returnValue
	})

	r.VM = vm

	return output
}
