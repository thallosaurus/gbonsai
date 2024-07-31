package main

import (
	"fmt"
	gbonsai "gbonsai/pkg"
)

func main() {
	bonsai, _ := gbonsai.Run(1)
	fmt.Print(bonsai)
}
