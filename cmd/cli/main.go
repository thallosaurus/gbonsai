package main

import (
	"fmt"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {
	conf := gbonsai.NewConfig(100, 100, 1, 10)
	bonsai, _ := gbonsai.Run(conf)
	fmt.Print(bonsai)
}
