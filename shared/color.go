package shared

import (
	"strconv"
)

type RGB struct {
	R, G, B uint8
}

// RGBColor 颜色格式化
func (c *RGB) RGBColor() (hexString string) {

	if r := strconv.FormatInt(int64(c.R), 16); len(r) == 1 {
		hexString += "0" + r
	} else {
		hexString += r
	}

	if g := strconv.FormatInt(int64(c.G), 16); len(g) == 1 {
		hexString += "0" + g
	} else {
		hexString += g
	}

	if b := strconv.FormatInt(int64(c.B), 16); len(b) == 1 {
		hexString += "0" + b
	} else {
		hexString += b
	}

	return
}

// ColorLib 默认颜色表
var ColorLib = struct {
	Aqua, Black, Blue, Fuchsia, Gray, Green, Lime, Maroon, Navy, Olive, Purple, Red, Silver, Teal, White, Yellow [2]string
}{
	Aqua:    [2]string{"00FFFF", "aqua"},
	Black:   [2]string{"000000", "black"},
	Blue:    [2]string{"0000FF", "blue"},
	Fuchsia: [2]string{"FFOOFF", "fuchsia"},
	Gray:    [2]string{"808080", "gray"},
	Green:   [2]string{"008000", "green"},
	Lime:    [2]string{"00FF00", "lime"},
	Maroon:  [2]string{"800000", "maroon"},
	Navy:    [2]string{"000080", "navy"},
	Olive:   [2]string{"808000", "olive"},
	Purple:  [2]string{"800080", "purple"},
	Red:     [2]string{"FF0000", "red"},
	Silver:  [2]string{"COCOCO", "sliver"},
	Teal:    [2]string{"008080", "teal"},
	White:   [2]string{"FFFFFF", "white"},
	Yellow:  [2]string{"FFFF00", "yellow"},
}
