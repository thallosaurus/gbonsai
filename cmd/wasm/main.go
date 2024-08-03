//go:build wasm

package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {
	js.Global().Set("gbonsai", js.FuncOf(gbonsai_func))
	<-make(chan bool)
}

func gbonsai_func(this js.Value, args []js.Value) interface{} {
	seed := args[0].Int()
	life := args[1].Int()

	conf := gbonsai.NewConfig(175, 100, int64(seed), life)
	bonsai, pot := gbonsai.Run(conf, nil)

	s := fmt.Sprintf("<div id=\"tree\">%s</div><pre id=\"pot\" style=\"color: white\">%s</pre>", bonsai.HtmlString(), pot.String())
	return s
}

//export gbonsai_string
func get_as_string(seed int, life int) string {
	conf := gbonsai.NewConfig(200, 100, int64(seed), life)
	//bonsai, _ := gbonsai.Run(conf)
	bonsai, _ := gbonsai.Run(conf, nil)

	return bonsai.String()
}

//export generation
func generation() {
	b := js.Global().Get("document").Call("getElementById", "tree")
	p := js.Global().Get("document").Call("getElementById", "pot")
	fmt.Println(b)

	for i := range 200 {
		conf := gbonsai.NewConfig(200, 100, rand.Int63(), i)
		bonsai, pot := gbonsai.Run(conf, nil)
		b.Set("innerHTML", bonsai.HtmlString())
		p.Set("innerHTML", pot.HtmlString())
		time.Sleep(300)
	}
}
