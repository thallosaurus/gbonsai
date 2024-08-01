package main

import (
	"fmt"
	"math/rand"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {
	conf := gbonsai.NewConfig(200, 100, rand.Int63(), 100)
	bonsai, _ := gbonsai.Run(conf)
	fmt.Print(bonsai)
}
