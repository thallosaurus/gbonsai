package gbonsai

import (
	"bytes"
	"fmt"
	"strings"
)

type CharCell struct {
	c     []byte
	color Color
}

func NewGrowingVector(width, height int) GrowingVector {
	locked := false
	chardata := make(map[int]*CharCell)
	return GrowingVector{
		chardata: chardata,
		width:    width,
		height:   height,
		locked:   locked,
	}
}

type GrowingVector struct {
	chardata map[int]*CharCell
	width    int
	height   int
	locked   bool
}

func (g *GrowingVector) Mvwprintw(x, y int, s string, color Color) {
	g.SetString(x, y, s, color)
}
func (g *GrowingVector) Set(x, y int, c *CharCell) {
	if x < g.width && y < g.height {
		index := xy_to_index(g.width, x, y)
		g.SetIndex(index, c)
	}
}
func (g *GrowingVector) SetIndex(index int, c *CharCell) {
	g.chardata[index] = c
}
func (g *GrowingVector) SetString(x, y int, s string, color Color) {
	for pos, char := range s {
		g.Set(x+pos, y, &CharCell{
			c:     []byte(string(char)),
			color: color,
		})
	}
}

func (g *GrowingVector) Get(x, y int) *CharCell {
	index := xy_to_index(g.width, x, y)
	return g.chardata[index]
}
func (g *GrowingVector) String() string {
	buf := make([]byte, g.width*g.height)
	w := bytes.NewBuffer(buf)

	for y := range g.height {
		for x := range g.width {
			if d := g.Get(x, y); d != nil {
				w.Write(d.c)
			} else {
				w.Write([]byte(" "))
			}
		}
		w.Write([]byte("\n"))
	}

	return w.String()
}

func (g *GrowingVector) HtmlString() string {
	buf := make([]byte, 0)
	w := bytes.NewBuffer(buf)

	s := g.String()

	in_space := false

	x := 0
	y := -1

	leaf_counter := 0

	var split_buf []string
	sd := strings.Split(s, "\n")
	for _, v := range sd {
		if len(strings.TrimSpace(v)) == 0 {
			y++
			continue
		}

		split_buf = append(split_buf, v)
	}

	for _, char := range strings.Join(split_buf, "\n") {

		switch string(char) {
		case " ":
			if !in_space {
				w.WriteString("<span>")
			}
			in_space = true
			w.WriteString("&nbsp;")
			x++

		case "\n":
			if in_space {
				w.WriteString("</span>")
				leaf_counter = 0
			}
			in_space = false
			w.WriteString("<br>")
			y++
			x = 0

		default:
			if in_space {
				w.WriteString("</span>")
			}
			in_space = false
			cell := g.Get(x, y)

			leaf_counter++

			if cell != nil {
				html := fmt.Sprintf("<span style=\"background-color: black; color: %s;\">%s</span>", DecodeColorHtml(cell.color), string(char))
				w.WriteString(html)
			} else {

				html := fmt.Sprintf("<span style=\"background-color: black; color: white;\">%s</span>", string(char))
				w.WriteString(html)
			}

			x++
		}

	}

	return w.String()
}
