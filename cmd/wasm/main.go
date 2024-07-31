//go:build wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {}

//export gbonsai_dom
func write_to_document(seed int64) {
	fmt.Println("Called go code from js")

	conf := gbonsai.NewConfig(200, 100, seed, 100)
	//bonsai, _ := gbonsai.Run(conf)
	bonsai, pot := gbonsai.Run(conf)

	js.Global().Get("document").Call("getElementById", "tree").Set("innerText", bonsai)
	js.Global().Get("document").Call("getElementById", "pot").Set("innerText", pot)
}
