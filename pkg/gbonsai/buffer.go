package gbonsai

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

type StrBuf = map[int][]byte

type BufAttr = []int

type TwoDimStringBuf struct {
	width  int
	height int
	vec    StrBuf
	attrs  BufAttr

	last_color       ColorPair
	last_color_index int
}

func NewTwoDimStringBuf(w, h int) TwoDimStringBuf {
	col := ColorPair{bg: White, fg: 4}
	// Allocate memory for the row buffers
	vc := make(StrBuf, h)
	//attrs := make(BufAttr)
	var attrs BufAttr
	// Initialize children for each row
	for i := range h {
		s_buf := make([]byte, w)
		for j := range s_buf {
			s_buf[j] = []byte(" ")[0]
			attrs = append(attrs, 4)
		}
		vc[i] = s_buf
	}

	return TwoDimStringBuf{
		width:      w,
		height:     h,
		vec:        vc,
		attrs:      attrs,
		last_color: col,
	}
}

func (t TwoDimStringBuf) String() string {
	buf := make([]byte, t.width*t.height)
	w := bytes.NewBuffer(buf)

	for i := range t.height {
		if ts := strings.TrimSpace(string(t.vec[i])); len(ts) > 0 {
			w.WriteString(fmt.Sprintf("%s\n", string(t.vec[i])))
		}
	}
	return w.String()
}

func (t *TwoDimStringBuf) HtmlString() string {
	buf := make([]byte, t.width*t.height)
	w := bytes.NewBuffer(buf)

	anal := analyseBuffer(t)
	fmt.Printf("%+v\n", anal)

	in_space := false

	for y := range t.height {

		leaf_count := 0

		row := strings.Split(string(t.vec[y]), "")
		for x := range len(row) {
			if in_space {
				w.WriteString("&nbsp;")
				//</span>")

				if x+1 == t.width || string(row[x+1]) != " " {
					in_space = false
					w.WriteString("</span>")
				}

			} else {

				if row[x] == " " {
					in_space = true
					w.WriteString("<span>&nbsp;")
				} else {
					index := xy_to_index(t.width, x, y)
					color := t.attrs[index]

					html := fmt.Sprintf("<span style=\"background-color: black; color: %s;\">%s</span>", DecodeColorHtml(color), string(row[x]))
					//w.WriteString(fmt.Sprintf(, string(t.vec[i])))
					w.WriteString(html)
					leaf_count++
				}
			}
		}
		w.WriteString("<br>")
	}

	return w.String()
}

func (t *TwoDimStringBuf) Wattron(colorIndex int) {
	t.last_color_index = colorIndex
}

func (t *TwoDimStringBuf) Mvwprintw(x, y int, s *CharCell) {
	if y < t.height {
		//row := []byte(t.vec[y])
		st := s.c

		for i, v := range st {
			if x+i < len(t.vec[y]) && x+i > 0 {
				t.vec[y][x+i] = v

				//col := t.last_color_index
				xyindex := xy_to_index(t.width, x+i, y)
				t.attrs[xyindex] = s.color
			}
		}
	}

}

func (t TwoDimStringBuf) Wprintw(s string) {

}

func xy_to_index(w, x, y int) int {
	return y*w + x
}

type AnalysisResult struct {
	left   int
	right  int
	top    int
	bottom int
}

func analyseBuffer(buf *TwoDimStringBuf) AnalysisResult {
	left := 1
	right := buf.width
	top := 1
	bottom := buf.height

	//tree_started := false

	for y := range len(buf.vec) {

		row_s := string(buf.vec[y])
		row := strings.Split(row_s, "")

		left_whitespace := len(row)
		right_whitespace := 0

		for x := range len(row) {
			//count whitespace
			v := row[x]
			if v != " " {
				left_whitespace = x
				fmt.Printf("%d\n", x)
				break
			}

		}

		fmt.Printf("left: %d, right: %d\n", left_whitespace, right_whitespace)

		left = max(left_whitespace, left)
		//right = max(right, right_whitespace)

		//		c_left := countWhitespaceArray(row)
		/*fmt.Printf("row: %s", row)

		if left < c_left && !tree_started {
			//top =
			//left = c_left
			tree_started = true
		}

		if left > c_left && tree_started {

		}*/

		//for _ = range len(row) {

		// find left side
		//left = (left,countWhitespaceArray(row, false))
		/*if row[x] != " " && !left_found {
			left = x
			left_found = true
			//space
		} else {
		}*/
		//}
	}

	fmt.Printf("\n")

	return AnalysisResult{
		left, right, top, bottom,
	}
}

func countWhitespaceArray(a []string) int {
	for i, v := range a {
		log.Printf("log %s", v)
		if v != " " {
			return i
		}
	}

	return 187
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
