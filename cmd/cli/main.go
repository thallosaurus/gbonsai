package main

import (
	"fmt"
	gbonsai "gbonsai/pkg"
)

func main() {

	bonsai := gbonsai.Run(1)
	fmt.Print(bonsai)
	//log.Printf("%d %d %d", rand.Int(), rand.Int(), rand.Int())
}
