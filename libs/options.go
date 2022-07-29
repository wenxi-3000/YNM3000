package libs

type Options struct {
	Inputs   map[string]struct{}
	Results  string
	CmdInput CmdInput
	Paths    Paths
	Scan     Scan
	Org      string
	Clean    bool
}

type CmdInput struct {
	Input      string
	Inputs     []string
	InputFile  string
	ResultPath string
}

type Paths struct {
	Result string
	Root   string
	Org    string
}

type Scan struct {
	Flow       string
	FlowFolder string
}
