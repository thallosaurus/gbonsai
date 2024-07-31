package gbonsai

import (
	"math/rand"
	"slices"
)

func Run(conf Config) (string, string) {
	//TODO improve precision

	//conf := NewConfig()

	counters := Counters{}

	obj := NCObjects{
		treeBuf: NewTwoDimStringBuf(conf.max_x, conf.max_y),
		baseBuf: NewTwoDimStringBuf(conf.max_x, conf.max_y),
	}

	conf.leaves = slices.Insert(conf.leaves, 0, "&")

	//for {
	growTree(&conf, &obj, &counters)

	return obj.treeBuf.String(), obj.baseBuf.String()
	//}
}

func drawBase(buf TwoDimStringBuf, baseType int) {
	switch baseType {
	case 1:
		buf.Wprintw(":")
	}
}

func growTree(conf *Config, objects *NCObjects, counters *Counters) {
	counters.shoots = 0
	counters.branches = 0
	counters.shootCounter = conf.rng.Int()

	branch(conf, objects, counters, conf.max_x/2, conf.max_y-1, Trunk, conf.lifeStart)
}

func branch(conf *Config, objects *NCObjects, counters *Counters, x int, y int, t BranchType, life int) {
	counters.branches++
	dx := 0
	dy := 0
	age := 0
	shootCooldown := conf.multiplier

	for life > 0 {
		life -= 1
		age = conf.lifeStart - life

		setDeltas(t, conf, life, age, conf.multiplier, &dx, &dy)

		max_y := conf.max_y
		if dy > 0 && y > (max_y-2) {
			dy--
		}

		if life < 3 {
			branch(conf, objects, counters, x, y, Dead, life)
		} else if t == Dying && life < (conf.multiplier+2) {
			branch(conf, objects, counters, x, y, Dying, life)
		} else if (t == ShootLeft || t == ShootRight) && life < (conf.multiplier+2) {
			branch(conf, objects, counters, x, y, Dying, life)
		} else if t == Trunk && (((conf.rng.Int() % 3) == 0) || life%conf.multiplier == 0) {
			if conf.rng.Int()%8 == 0 && life > 7 {
				shootCooldown = conf.multiplier * 2
				branch(conf, objects, counters, x, y, Trunk, life+conf.rng.Int()%5-2)
			} else if shootCooldown <= 0 {
				shootCooldown = conf.multiplier * 2
				shootLife := life + conf.multiplier

				counters.shoots += 1
				counters.shootCounter += 1
				branch(conf, objects, counters, x, y, counters.shootCounter%2+1, shootLife)
			}
		}

		shootCooldown--

		x += dx
		y += dy

		branchStr := chooseString(conf, t, life, dx, dy)

		// print
		objects.treeBuf.Mvwprintw(x, y, branchStr)

	}
}

func chooseString(conf *Config, t BranchType, life int, dx int, dy int) string {
	branchStr := "?"

	if life < 4 {
		t = Dying
	}

	switch t {
	case Trunk:
		if dy == 0 {
			branchStr = "/~"
		} else if dx < 0 {
			branchStr = "\\|"
		} else if dx == 0 {
			branchStr = "/|\\"
		} else if dx > 0 {
			branchStr = "|/"
		}

	case ShootLeft:
		if dy > 0 {
			branchStr = "\\"
		} else if dy == 0 {
			branchStr = "\\_"
		} else if dx < 0 {
			branchStr = "\\|"
		} else if dx == 0 {
			branchStr = "/|"
		} else if dx > 0 {
			branchStr = "/"
		}
	case Dying:
		fallthrough
	case Dead:
		branchStr = conf.leaves[conf.rng.Int()%conf.leavesSize]
	}

	return branchStr
}

func setDeltas(t BranchType, conf *Config, life int, age int, multiplier int, returnDx *int, returnDy *int) {
	dx := 0
	dy := 0
	var dice int

	switch t {
	case Trunk:

		if age <= 2 || life < 4 {
			dy = 0
			dx = (conf.rng.Int() % 3) - 1
		} else if age < (multiplier * 3) {
			if age%(multiplier/2) == 0 {
				dy = -1
			} else {
				dy = 0
			}

			roll(conf.rng, &dice, 10)
			if dice >= 0 && dice <= 0 {
				dx = -2
			} else if dice >= 1 && dice <= 3 {
				dx = -1
			} else if dice >= 4 && dice <= 5 {
				dx = 0
			} else if dice >= 6 && dice <= 8 {
				dx = 1
			} else if dice >= 9 && dice <= 9 {
				dx = 2
			}
		} else {
			roll(conf.rng, &dice, 10)
			if dice > 2 {
				dy = -1
			} else {
				dy = 0
			}
			dx = conf.rng.Int()%3 - 1
		}

	case ShootLeft:
		roll(conf.rng, &dice, 10)
		if dice >= 0 && dice <= 1 {
			dy = -1
		} else if dice >= 2 && dice <= 7 {
			dy = 0
		} else if dice >= 8 && dice <= 9 {
			dy = 1
		}

		roll(conf.rng, &dice, 10)
		if dice >= 0 && dice <= 1 {
			dx = -2
		} else if dice >= 2 && dice <= 5 {
			dx = -1
		} else if dice >= 6 && dice <= 8 {
			dx = 0
		} else if dice >= 9 && dice <= 9 {
			dx = 1
		}
	case ShootRight:
		roll(conf.rng, &dice, 10)
		if dice >= 0 && dice <= 1 {
			dy = -1
		} else if dice >= 2 && dice <= 5 {
			dy = 0
		} else if dice >= 8 && dice <= 9 {
			dy = 1
		}

		roll(conf.rng, &dice, 10)
		if dice >= 0 && dice <= 1 {
			dx = 2
		} else if dice >= 2 && dice <= 5 {
			dx = 1
		} else if dice >= 6 && dice <= 8 {
			dx = 0
		} else if dice >= 9 && dice <= 9 {
			dx = -1
		}

	case Dying:
		roll(conf.rng, &dice, 10)
		if dice >= 0 && dice <= 1 {
			dy = -1
		} else if dice >= 2 && dice <= 8 {
			dy = 0
		} else if dice >= 9 && dice <= 9 {
			dy = 1
		}

		roll(conf.rng, &dice, 15)
		if dice >= 0 && dice <= 0 {
			dx = -3
		} else if dice >= 1 && dice <= 2 {
			dx = -2
		} else if dice >= 3 && dice <= 5 {
			dx = -1
		} else if dice >= 6 && dice <= 8 {
			dx = 0
		} else if dice >= 9 && dice <= 11 {
			dx = 1
		} else if dice >= 12 && dice <= 13 {
			dx = 2
		} else if dice >= 14 && dice <= 14 {
			dx = 3
		}

	case Dead:
		roll(conf.rng, &dice, 10)
		if dice >= 0 && dice <= 2 {
			dy = -1
		} else if dice >= 3 && dice <= 6 {
			dy = 0
		} else if dice >= 7 && dice <= 9 {
			dy = 1
		}
		dx = (conf.rng.Int() % 3) - 1
	}
	*returnDx = dx
	*returnDy = dy
}

func roll(rng *rand.Rand, dice *int, mod int) {
	*dice = rng.Int() % mod
}

type Config struct {
	live       int
	infinite   int
	lifeStart  int
	seed       int64
	multiplier int
	max_x      int
	max_y      int
	leaves     []string
	leavesSize int
	rng        *rand.Rand
}

func NewConfig(w, h int, seed int64, life int) Config {
	return Config{
		live:       0,
		infinite:   0,
		lifeStart:  life,
		seed:       int64(seed),
		max_x:      w,
		max_y:      h,
		leaves:     make([]string, 64),
		leavesSize: 1,
		multiplier: 5,
		rng:        rand.New(rand.NewSource(int64(seed))),
	}
}

type Counters struct {
	branches     int
	shoots       int
	shootCounter int
}

type NCObjects struct {
	baseBuf TwoDimStringBuf
	treeBuf TwoDimStringBuf
}

type BranchType = int

const (
	Trunk      BranchType = 0
	ShootLeft  BranchType = 1
	ShootRight BranchType = 2
	Dying      BranchType = 3
	Dead       BranchType = 4
)
