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
func gbonsai_dom(seed int64, life int) {
	fmt.Println("Called go code from js")

	conf := gbonsai.NewConfig(175, 100, seed, life)
	bonsai, pot := gbonsai.Run(conf)

	js.Global().Get("document").Call("getElementById", "tree").Set("innerHTML", bonsai.HtmlString())
	js.Global().Get("document").Call("getElementById", "pot").Set("innerHTML", pot.HtmlString())
}

//export generation
func generation() {
	b := js.Global().Get("document").Call("getElementById", "tree")
	p := js.Global().Get("document").Call("getElementById", "pot")
	fmt.Println(b)

	for i := range 200 {
		conf := gbonsai.NewConfig(200, 100, rand.Int63(), i)
		bonsai, pot := gbonsai.Run(conf)
		b.Set("innerHTML", bonsai.HtmlString())
		p.Set("innerHTML", pot.HtmlString())
		time.Sleep(300)
	}
}
