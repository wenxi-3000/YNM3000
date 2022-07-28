package scripts

import (
	"YNM3000/utils"
	"log"
)

// func (r *Runner) LoadScripts() {
// 	//var output string
// 	vm := otto.New()
// 	log.Println("xxxx")
// 	vm.Run(`
//     abc = 2 + 2;
//     console.log("The value of abc is " + abc); // 4
// `)
// vm.Set(Append, func(call otto.FunctionCall) otto.Value {
// 	dest := call.Argument(0).String()
// 	src := call.Argument(1).String()
// 	Append(dest, src)
// 	returnValue, _ := otto.ToValue(true)
// 	return returnValue
// })

// }

// func Append(dest string, src string) {
// 	if !utils.FileExists(src) || utils.FileLength(src) <= 0 {
// 		log.Println("error to append %v", src)
// 	}
// 	data := utils.GetFileContent(src)
// 	utils.AppendToContent(dest, data)
// }

func Append(dest string, src string) {
	if !utils.FileExists(src) || utils.FileLength(src) <= 0 {
		log.Println(dest, src)
	}
	data := utils.GetFileContent(src)
	utils.AppendToContent(dest, data)
}
