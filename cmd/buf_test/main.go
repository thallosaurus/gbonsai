package main

import (
	"fmt"

	"github.com/thallosaurus/gbonsai/pkg/gbonsai"
)

func main() {
	buf := gbonsai.NewTwoDimStringBuf(100, 100)

	buf.Mvwprintw(0, 0, "*")
	/*
		buf.Mvwprintw(0, 1, "*")
		buf.Mvwprintw(0, 2, "*")
		buf.Mvwprintw(0, 3, "*")
		buf.Mvwprintw(0, 4, "*")
		buf.Mvwprintw(1, 4, "*")
		buf.Mvwprintw(2, 4, "*")
		buf.Mvwprintw(3, 4, "*")
		buf.Mvwprintw(4, 4, "*")*/

	fmt.Print(buf.String())
}
