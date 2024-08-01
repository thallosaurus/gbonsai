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

func (g *GrowingVector) Mvwprintw(x, y int, c *CharCell) {
	g.Set(x, y, c)
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
	for i, v := range strings.Split(s, "") {

		g.Set(x+i, y, &CharCell{
			c:     []byte(v),
			color: color,
			//x:     x + i,
			//y:     y,
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

func (g *GrowingVector) isRowEmpty(y int) bool {
	return true
}

func (g *GrowingVector) HtmlString() string {
	buf := make([]byte, g.width*g.height)
	w := bytes.NewBuffer(buf)

	for y := range g.height {

		// if line is only spaces skip it

		//if rr :=

		rowbuf := make([]byte, g.width)
		wb := bytes.NewBuffer(rowbuf)

		open_span := false

		for x := range g.width {
			index := xy_to_index(g.width, x, y)
			v := g.chardata[index]

			if open_span {
				wb.WriteString("&nbsp;")

				if x+1 == g.width || g.chardata[x+1] != nil {
					open_span = false
					wb.WriteString("</span>")
				}
			} else {

				if v != nil {
					html := fmt.Sprintf("<span style=\"background-color: black; color: %s;\">%s</span>", DecodeColorHtml(v.color), string(v.c))
					wb.WriteString(html)
				} else {
					open_span = true
					wb.WriteString("<span>")
					// nothing
					wb.WriteString("&nbsp;")
					//</span>")

				}
			}

		}

		w.WriteString(wb.String())
		w.WriteString("<br>")
	}

	return w.String()
}
