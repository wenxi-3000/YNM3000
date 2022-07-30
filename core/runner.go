package core

import (
	"YNM3000/libs"
	"YNM3000/utils"
	"bufio"
	"fmt"
	"log"
	"os/exec"

	"github.com/robertkrimen/otto"
)

type Runner struct {
	Input     string
	InputType string //domain, ip, url, domain-file, url-file,ip-file
	Routines  []libs.Routine
	Paths     libs.Paths
	Reports   []string
	VM        *otto.Otto
	Clean     bool
}

func InitRunner(input string, opt libs.Options) Runner {
	var runner Runner
	runner.Input = input
	runner.Paths = opt.Paths
	runner.Routines = Parse(opt)
	runner.Clean = opt.Clean
	return runner
	//解析module模板
	//moduleFolder := path.Join(opt.Scan.FlowFolder, opt.Scan.Flow)

}

func Run(input string, opt libs.Options) {
	runner := InitRunner(input, opt)
	runner.PrepareModule()
	runner.InitVM()
	runner.Start()

}

func (r *Runner) Start() {
	for _, routine := range r.Routines {
		// log.Println(routine.ParsedModules)
		for _, module := range routine.ParsedModules {
			// log.Println("=========================")
			// log.Println(module)

			//pre_run
			for _, pre := range module.PreRun {
				r.VM.Run(pre)
			}

			//run steps
			for _, step := range module.Steps {
				// log.Println("=========================")
				// log.Println(step)
				err := r.CheckRequired(step.Required)
				if err != nil {
					log.Println(err)
				}

				if len(step.Commands) > 0 {
					for _, command := range step.Commands {
						results, err := RunCommand(command)
						if err != nil {
							log.Println(err)
						}
						log.Println(results)
					}
				}

				for _, script := range step.Scripts {
					r.VM.Run(script)
					//Append(script)
				}
			}

			for _, postRun := range module.PostRun {
				r.VM.Run(postRun)
			}
		}
	}
}

func RunCommand(cmd string) (string, error) {
	command := []string{
		"bash",
		"-c",
		cmd,
	}
	var output string
	cmdx := exec.Command(command[0], command[1:]...)
	log.Println(cmdx)
	cmdReader, _ := cmdx.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			out := scanner.Text()
			output += out + "\n"
		}
	}()

	if err := cmdx.Start(); err != nil {
		return output, err
	}

	if err := cmdx.Wait(); err != nil {
		return output, err
	}

	return output, nil

}

// CheckRequired check if required file exist or not
func (r *Runner) CheckRequired(requires []string) error {
	if len(requires) == 0 {
		return nil
	}

	for _, require := range requires {
		//require = ResolveData(require, options.Scan.ROptions)
		if !utils.FileExists(require) {
			return fmt.Errorf("Missing Requirement: %s", require)
		}
		continue
	}
	return nil
}

// func (r *Runner) ConditionExecScript(script string) bool {
// 	value, err := r.VM.Run(script)

// 	if err == nil {
// 		out, nerr := value.ToBoolean()
// 		if nerr == nil {
// 			return out
// 		}
// 	}

// 	return false
// }
