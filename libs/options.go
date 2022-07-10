package libs

type Options struct {
	Input Input
	Paths Paths
}

type Input struct {
	Input      string
	Inputs     []string
	InputFile  string
	ResultPath string
}

type Paths struct {
	Result string
	Org    string
}
