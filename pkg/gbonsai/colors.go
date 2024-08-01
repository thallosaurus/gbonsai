package gbonsai

import (
	"fmt"
	"log"
)

type Color = int

type ColorPair struct {
	bg Color
	fg Color
}

var pairs []ColorPair

/*func ColorPair(fg int) int {
	//return 1 << n
	return 0
}*/

func Pair(i int) ColorPair {
	if i < len(pairs) {
		return pairs[i]
	} else {
		return ColorPair{bg: White, fg: 0}
	}
}

func InitColors() {
	fmt.Println("Init Colors")
	bg := Transparent

	//init_pair(0, White, -1)

	for i := range 16 {
		//init_pair(i, i, bg)
		p := ColorPair{
			bg, i,
		}
		pairs = append(pairs, p)
	}

	//fmt.Println(pairs)

	/*init_pair(8, 7, bg) // gray will look white
	init_pair(9, 1, bg)
	init_pair(10, 2, bg)
	init_pair(11, 3, bg)
	init_pair(12, 4, bg)
	init_pair(13, 5, bg)
	init_pair(14, 6, bg)
	init_pair(15, 7, bg)*/
}

func DecodeColorHtml(c Color) string {

	switch c {
	//	case Transparent:
	//		return "transparent"

	case Black:
		return "black"

	case Maroon:
		return "#800000"

	case Green:
		return "#008000"

	case Olive:
		return "#808000"

	case Navy:
		return "#000080"

	case Purple:
		return "#800080"

	case Teal:
		return "#008080"

	case Silver:
		return "#c0c0c0"

	case Grey:
		return "#808080"

	case Red:
		return "#ff0000"

	case Lime:
		return "#00ff00"

	case Yellow:
		return "#ffff00"

	case Blue:
		return "#0000ff"

	case Fuchsia:
		return "#ff00ff"

	case Aqua:
		return "#00ffff"

	case White:
		return "#ffffff"

	default:
		log.Printf("Unknown Color: %d", c)
		return "transparent"
	}

}

const (
	Yellow Color = 11
	Lime   Color = 10
	Olive  Color = 3
	Green  Color = 2

	Transparent Color = -1
	Black       Color = 0
	Navy        Color = 1
	Teal        Color = 6
	Maroon      Color = 4
	Fuchsia     Color = 5
	White       Color = 7
	Grey        Color = 8
	Blue        Color = 9
	Aqua        Color = 14
	Red         Color = 12
	Purple      Color = 13
	Silver      Color = 15
)
