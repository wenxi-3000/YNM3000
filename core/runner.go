package core

import (
	"YNM3000/libs"
	"bufio"
	"log"
	"os/exec"
)

type Runner struct {
	Input     string
	InputType string //domain, ip, url, domain-file, url-file,ip-file
	Routines  []libs.Routine
	Paths     libs.Paths
}

func InitRuner(input string, opt libs.Options) Runner {
	var runner Runner
	runner.Input = input
	runner.Paths = opt.Paths
	runner.Routines = Parse(opt)
	return runner
	//解析module模板
	//moduleFolder := path.Join(opt.Scan.FlowFolder, opt.Scan.Flow)

}

func Run(input string, opt libs.Options) {
	runner := InitRuner(input, opt)
	runner.PrepareModule()
	runner.Start()

}

func (r *Runner) Start() {
	for _, routine := range r.Routines {
		// log.Println("=========================")
		// log.Println(routine.ParsedModules)
		for _, module := range routine.ParsedModules {
			// log.Println("=========================")
			// log.Println(module)
			for _, step := range module.Steps {
				// log.Println("=========================")
				// log.Println(step)
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
					log.Println("=========================")
					log.Println(script)
				}
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
			output += out
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
