package gbonsai

type Color = int

const (
	Yellow     Color = 0
	Brown      Color = 1
	LightGreen Color = 2
	DarkGreen  Color = 3
)

type ColorPair struct {
	bg Color
	fg Color
}

var pairs = make(map[int]ColorPair, 16)

/*func ColorPair(fg int) int {
	//return 1 << n
	return 0
}*/

func Pair(i int) ColorPair {
	if len(pairs) > i {
		return pairs[i]
	} else {
		return ColorPair{bg: 0, fg: 0}
	}
}

func InitColors() {
	for i := range 16 {
		pairs[i] = ColorPair{
			bg: i, fg: i,
		}
	}
}
