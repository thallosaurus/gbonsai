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

	last_left_offset := g.width
	last_right_offset := g.width

	var split_buf []string
	{

		sd := strings.Split(s, "\n")
		for _, v := range sd {
			if len(strings.TrimSpace(v)) == 0 {
				y++
				continue
			}

			//get left and right offset
			left_offset := count_whitespace(v, false)
			right_offset := count_whitespace(v, true)

			//if left offset is smaller than last one, set last one to offset
			if left_offset < last_left_offset {
				last_left_offset = left_offset
			}
			//if right offset is bigger than last one, set last one to offset
			if right_offset < last_right_offset {
				last_right_offset = right_offset
			}

			fmt.Printf("%d, %d: %s\n", last_left_offset, last_right_offset, v)

			split_buf = append(split_buf, v)
		}
	}

	x = last_left_offset

	fmt.Printf("Left Offset: %d, Right Offset: %d\n", last_left_offset, last_right_offset)

	for i, s := range split_buf {
		split_buf[i] = s[last_left_offset : g.width-last_right_offset]
		//fmt.Println(split_buf[i])
	}

	w.WriteString("<p class='row'>")

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
			w.WriteString("</p><br><p class='row'>")
			y++
			x = last_left_offset

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

	w.WriteString("</p>")

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

func count_whitespace(s string, right bool) int {
	c := 0

	if right {
		s = reverse(s)
	}

	for i, v := range s {
		if string(v) != " " {
			break
		}

		c = i
	}

	return c
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
