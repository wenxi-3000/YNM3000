package libs

type Options struct {
	Inputs   map[string]struct{}
	CmdInput CmdInput
	Paths    Paths
	Scan     Scan
}

type CmdInput struct {
	Input      string
	Inputs     []string
	InputFile  string
	ResultPath string
}

type Paths struct {
	Result string
	Org    string
	Root   string
}

type Scan struct {
	Flow       string
	FlowFolder string
}
