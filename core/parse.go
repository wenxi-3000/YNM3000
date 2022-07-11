package core

import (
	"YNM3000/libs"
	"io/ioutil"
	"log"
	"path"

	"github.com/Shopify/yaml"
)

func ParseFlow(flowFile string) (libs.Flow, error) {
	var flow libs.Flow
	yamlFile, err := ioutil.ReadFile(flowFile)
	if err != nil {
		return flow, err
	}
	err = yaml.Unmarshal(yamlFile, &flow)
	if err != nil {
		return flow, err
	}

	return flow, nil
}

// ParseModules parse module file
func ParseModules(moduleFile string) (libs.Module, error) {

	var module libs.Module

	yamlFile, err := ioutil.ReadFile(moduleFile)
	if err != nil {
		return module, err
	}
	err = yaml.Unmarshal(yamlFile, &module)
	if err != nil {
		return module, err
	}
	return module, err
}

func Parse(opt libs.Options) []libs.Routine {
	var routines []libs.Routine
	//解析flow模板
	flowFile := path.Join(opt.Scan.FlowFolder, opt.Scan.Flow+".yaml")
	parseFlow, err := ParseFlow(flowFile)
	if err != nil {
		log.Println(err)
	}

	moduleFolder := path.Join(opt.Scan.FlowFolder, opt.Scan.Flow)
	for _, routine := range parseFlow.Routines {
		for _, module := range routine.Modules {
			moduleFile := path.Join(moduleFolder, module+".yaml")
			parsedModule, err := ParseModules(moduleFile)
			if err != nil {
				log.Println(err)
			}
			routine.ParsedModules = append(routine.ParsedModules, parsedModule)
		}
		routines = append(routines, routine)
	}
	return routines
}
