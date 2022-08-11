package logger

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	Info("xxx %s", "sss")
	Println("xxxx")
	Silent("aaaaa")

	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	NewLogger(LSilent)
	Info("xxx %s", "sss")
	Println("xxxx")
	Silent("aaaaa")

	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	SetLevel(LSilent)
	Info("xxx %s", "sss")
	Println("xxxx")
	Silent("aaaaa")

	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	log.Println("ffff")
	log.New(os.Stderr, "[VERBOSE]", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("ffff")
	log.Println("ffff")

	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("zzzzzzzzzz")

}
