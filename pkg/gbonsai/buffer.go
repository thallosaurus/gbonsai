package gbonsai

import (
	"bytes"
	"fmt"
	"strings"
)

type StrBuf = map[int][]byte

type BufAttr = map[int]BufAttrEntry
type BufAttrEntry = struct {
	color int
}

type TwoDimStringBuf struct {
	width  int
	height int
	vec    StrBuf
	attrs  BufAttr

	last_color int
}

func NewTwoDimStringBuf(w, h int) TwoDimStringBuf {
	// Allocate memory for the row buffers
	vc := make(StrBuf, h)
	// Initialize children for each row
	for i := range h {
		s_buf := make([]byte, w)
		for j := range s_buf {
			s_buf[j] = []byte(" ")[0]
		}
		vc[i] = s_buf
	}

	// initialize attributes
	attrs := make(BufAttr, w*h)
	for i := range h * w {
		attrs[i] = BufAttrEntry{
			color: 0,
		}
	}

	return TwoDimStringBuf{
		width:      w,
		height:     h,
		vec:        vc,
		attrs:      attrs,
		last_color: 0,
	}
}

/*func (t TwoDimStringBuf) Read() byte {

	return byte(0)
}*/

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

func (t TwoDimStringBuf) Mvwprintw(x, y int, s string) {
	//index := xy_to_index(0, x, y)
	//t.vec[index] = slices.Insert(t.vec[index])

	if y < t.height {
		//row := []byte(t.vec[y])
		st := []byte(s)

		for i, v := range st {
			if x+i < len(t.vec[y]) && x+i > 0 {
				t.vec[y][x+i] = v
				t.attrs[xy_to_index(t.width, x, y)] = BufAttrEntry{
					color: t.last_color,
				}
			}
		}
	}

}

func (t TwoDimStringBuf) Wprintw(s string) {

}

func xy_to_index(w, x, y int) int {
	return y*w + x
}

type BufferEntry struct {
	size int
	data string
}
