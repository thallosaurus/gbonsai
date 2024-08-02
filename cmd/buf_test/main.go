package main

import (
	"fmt"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {
	buf := gbonsai.NewGrowingVector(10, 10)

	buf.Mvwprintw(0, 0, "*", gbonsai.White, nil)

	fmt.Print(buf.String())
	fmt.Print(buf.HtmlString())
}
