package main

import (
	"fmt"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {
	conf := gbonsai.NewConfig(100, 50, 1, 50)
	bonsai, _ := gbonsai.Run(conf)
	fmt.Print(bonsai)
}
