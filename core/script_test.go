package core

import (
	"YNM3000/libs"
	"testing"
)

func TestLoadScripts(t *testing.T) {
	loggerSetFlags(loggerLdate | loggerLtime | loggerLshortfile)
	var opt = libs.Options{}
	input := "hackerone.com"
	runner := InitRunner(input, opt)
	runner.LoadScripts()
}
