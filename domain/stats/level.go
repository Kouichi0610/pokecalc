// レベル(1-100)
package stats

type Level uint

func NewLevel(l uint) Level {
	if l < 1 {
		l = 1
	}
	if l > 100 {
		l = 100
	}
	return Level(l)
}
