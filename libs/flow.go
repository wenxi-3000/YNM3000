package libs

//[Flow[name, Params,Routines[Modules,ParsedModules[] ] ]]

// Step struct to define component about a command
type Step struct {
	// timeout for commands and script
	Timeout    string
	Threads    string
	Conditions []string
	Required   []string
	Commands   []string
	Scripts    []string
}

type Module struct {
	Name    string
	Desc    string
	Params  []map[string]string
	Steps   []Step
	PostRun []string `yaml:"last_run"`
	Report  []string
}

type Routine struct {
	RoutineName   string
	ParsedModules []Module
	Modules       []string
}

type Flow struct {
	Name      string
	Validator string // domain, cidr, ip or domain-file, cidr-file and so on
	Desc      string

	Params   []map[string]string
	Routines []Routine
	Steps    []Step
}
