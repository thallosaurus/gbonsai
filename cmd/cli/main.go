package main

import (
	"fmt"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {
	bonsai, _ := gbonsai.Run(1)
	fmt.Print(bonsai)
}
