package core

import (
	"bytes"
	"log"
	"path"
	"text/template"
)

//读取模板里的Params部分
func (r *Runner) PreparTemPara() map[string]string {
	temPara := make(map[string]string)
	for _, routine := range r.Routines {
		for _, module := range routine.ParsedModules {
			if len(module.Params) > 0 {
				log.Println("xxxxxxxxxxxxxxx")
				log.Println(module.Params)
				for _, param := range module.Params {
					for k, v := range param {
						log.Println(k, v)
						temPara[k] = v
					}
				}
			}
		}
	}
	return temPara
}

//定义模板里的参数
func (r *Runner) PrepareTem() map[string]string {

	//先解析模板里的params
	temParam := r.PreparTemPara()
	temp := temParam

	temp["Binaries"] = path.Join(r.Paths.Root, "binaries")
	temp["Output"] = path.Join(r.Paths.Root, "results", "org")
	temp["Target"] = r.Input

	return temp
}

//替换模板中的{{.xxxx}}等
func (r *Runner) PrepareModule() {
	//要替换的参数
	temp := r.PrepareTem()

	for _, routine := range r.Routines {
		// log.Println("=========================")
		// log.Println(routine.ParsedModules)
		for _, module := range routine.ParsedModules {
			// log.Println("=========================")
			// log.Println(module)
			for i, step := range module.Steps {
				module.Steps[i].Threads = ResolveData(step.Threads, temp)
				module.Steps[i].Commands = ResolveSlice(step.Commands, temp)
				module.Steps[i].Scripts = ResolveSlice(step.Scripts, temp)
			}

		}
	}
}

func ResolveData(format string, data map[string]string) string {
	t := template.Must(template.New("").Parse(format))
	buf := &bytes.Buffer{}
	err := t.Execute(buf, data)
	if err != nil {
		log.Println(err)
		return format
	}
	return buf.String()
}

func ResolveSlice(slice []string, data map[string]string) (resolveSlice []string) {
	for _, s := range slice {
		resolveSlice = append(resolveSlice, ResolveData(s, data))
	}
	return resolveSlice
}
