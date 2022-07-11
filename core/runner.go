package core

import (
	"YNM3000/libs"
	"log"
)

type Runner struct {
	Input     string
	InputType string //domain, ip, url, domain-file, url-file,ip-file
	Routines  []libs.Routine
}

func InitRuner(input string, opt libs.Options) Runner {
	var runner Runner
	runner.Input = input
	runner.Routines = Parse(opt)
	return runner
	//解析module模板
	//moduleFolder := path.Join(opt.Scan.FlowFolder, opt.Scan.Flow)

}

func Run(input string, opt libs.Options) {
	runner := InitRuner(input, opt)
	runner.Prepare()
	runner.Start()

}

func (r *Runner) Start() {
	log.Println("Start")
	log.Println(r.Routines)
	for _, routine := range r.Routines {
		// log.Println("=========================")
		// log.Println(routine.ParsedModules)
		for _, module := range routine.ParsedModules {
			// log.Println("=========================")
			// log.Println(module)
			for _, step := range module.Steps {
				// log.Println("=========================")
				// log.Println(step)
				for _, command := range step.Commands {
					log.Println("=========================")
					log.Println(command)
				}
				for _, script := range step.Scripts {
					log.Println("=========================")
					log.Println(script)
				}
			}
		}
	}
}
