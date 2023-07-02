package wind

type Wind int

const (
	East  Wind = 1
	South Wind = 2
	West  Wind = 3
	North Wind = 4
)

func (w Wind) String() string {
	switch w {
	case East:
		return "东"
	case South:
		return "南"
	case West:
		return "西"
	case North:
		return "北"
	}
	return ""
}

func (w Wind) Next() Wind {
	switch w {
	case East:
		return South
	case South:
		return West
	case West:
		return North
	case North:
		return East
	}
	return East
}

func (w Wind) Prev() Wind {
	switch w {
	case East:
		return North
	case South:
		return East
	case West:
		return South
	case North:
		return West
	}
	return East
}
