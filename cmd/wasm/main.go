//go:build wasm

package main

import (
	"fmt"
	gbonsai "gbonsai/pkg"
	"syscall/js"
)

func main() {}

//export call_me_from_js
func write_to_document(seed int64) {
	fmt.Println("Called go code from js")

	bonsai, pot := gbonsai.Run(1)

	js.Global().Get("document").Call("getElementById", "tree").Set("innerText", bonsai)
	js.Global().Get("document").Call("getElementById", "pot").Set("innerText", pot)
}
