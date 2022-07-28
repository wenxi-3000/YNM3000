package core

import (
	"YNM3000/libs"
	"log"
	"testing"
)

func TestLoadScripts(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	var opt = libs.Options{}
	input := "hackerone.com"
	runner := InitRunner(input, opt)
	runner.LoadScripts()
}
