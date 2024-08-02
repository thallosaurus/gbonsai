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

	conf := gbonsai.NewConfig(150, 100, seed, life)
	bonsai, pot := gbonsai.Run(conf)

	//var buf []byte
	//w := bytes.NewBuffer(buf)
	//w.WriteString(fmt.Sprintf("<div style=\"font-family: monospace; text-align: center;\"><div id=\"tree\">%s</div><pre id=\"pot\" style=\"color: white\">%s</pre></div>", bonsai.HtmlString(), pot.HtmlString()))

	js.Global().Get("document").Call("getElementById", "tree").Set("innerHTML", bonsai.HtmlString())
	js.Global().Get("document").Call("getElementById", "pot").Set("innerHTML", pot.String())

	//return w.String()
}

//export generation
func generation() {
	b := js.Global().Get("document").Call("getElementById", "tree")
	p := js.Global().Get("document").Call("getElementById", "pot")
	fmt.Println(b)

	for i := range 200 {
		conf := gbonsai.NewConfig(200, 100, rand.Int63(), i)
		//bonsai, _ := gbonsai.Run(conf)
		bonsai, pot := gbonsai.Run(conf)
		b.Set("innerHTML", bonsai.HtmlString())
		p.Set("innerHTML", pot.HtmlString())
		time.Sleep(300)
	}
}
