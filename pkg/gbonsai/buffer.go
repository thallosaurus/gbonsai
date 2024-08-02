package gbonsai

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

type CharCell struct {
	c     []byte
	color Color
	t     *int
}

func NewGrowingVector(width, height int) GrowingVector {
	locked := false
	chardata := make(map[int]*CharCell)
	return GrowingVector{
		chardata: chardata,
		width:    width,
		height:   height,
		locked:   locked,
		x:        0,
		y:        0,
	}
}

type GrowingVector struct {
	chardata map[int]*CharCell
	width    int
	height   int
	locked   bool

	x int
	y int
}

func (g *GrowingVector) Mvwprintw(x, y int, s string, color Color, t *int) {
	g.SetString(x, y, s, color, t)
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
func (g *GrowingVector) SetString(x, y int, s string, color Color, t *int) {
	for pos, char := range s {
		g.Set(x+pos, y, &CharCell{
			c:     []byte(string(char)),
			color: color,
			t:     t,
		})
	}
}

func (g *GrowingVector) Get(x, y int) *CharCell {
	index := xy_to_index(g.width, x, y)
	return g.chardata[index]
}
func (g *GrowingVector) String() string {
	buf := make([]byte, 0)
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
	{

		sd := strings.Split(s, "\n")
		for _, v := range sd {
			if len(strings.TrimSpace(v)) == 0 {
				y++
				continue
			}

			split_buf = append(split_buf, v)
		}
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

			if cell != nil && cell.t != nil {
				delay := rand.Int() % 1000
				html := fmt.Sprintf("<span style=\"animation-delay: -%dms; \"class=\"color-%d type-%d\">%s</span>", delay, cell.color, *cell.t, string(char))
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

func (g *GrowingVector) Movptr(x, y int) {
	g.x = x
	g.y = y
}

func (g *GrowingVector) Wprintw(s string, c Color) {
	g.SetString(g.x, g.y, s, c, nil)
	g.x += len(s)
}

func xy_to_index(w, x, y int) int {
	return y*w + x
}
