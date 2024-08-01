//go:build wasm

package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {}

//export gbonsai_string
func get_as_string(seed int, life int) string {
	conf := gbonsai.NewConfig(200, 100, int64(seed), life)
	//bonsai, _ := gbonsai.Run(conf)
	bonsai, _ := gbonsai.Run(conf)

	return bonsai.String()
}

//export gbonsai_dom
func write_to_document(seed int64, life int) {
	fmt.Println("Called go code from js")

	conf := gbonsai.NewConfig(200, 100, seed, life)
	//bonsai, _ := gbonsai.Run(conf)
	bonsai, _ := gbonsai.Run(conf)

	js.Global().Get("document").Call("getElementById", "tree").Set("innerHTML", bonsai.HtmlString())
	//js.Global().Get("document").Call("getElementById", "pot").Set("innerHTML", pot.HtmlString())
}

//export generation
func generation() {
	e := js.Global().Get("document").Call("getElementById", "tree")
	fmt.Println(e)

	for i := range 200 {
		conf := gbonsai.NewConfig(200, 100, rand.Int63(), i)
		//bonsai, _ := gbonsai.Run(conf)
		bonsai, _ := gbonsai.Run(conf)
		e.Set("innerHTML", bonsai.HtmlString())
		time.Sleep(300)
	}
}
