//go:build wasm

package main

import (
	"fmt"
	gbonsai "gbonsai/pkg"
	"log"
	"math/rand"
	"syscall/js"
)

func main() {}

//export call_me_from_js
func write_to_document() {
	fmt.Println("Called go code from js")

	bonsai := gbonsai.Run(1)

	log.Printf("%d %d %d", rand.Int(), rand.Int(), rand.Int())

	js.Global().Get("document").Call("getElementById", "output").Set("innerText", bonsai)
}
