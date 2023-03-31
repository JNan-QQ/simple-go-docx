package styles

import (
	"strconv"
	"strings"
)

// RGB 颜色 对应 rgb与16进制
// RGB.RGB 大小为3的 uint8 数组
// RGB.Hex 对应颜色的16进制，不支持透明度
type RGB struct {
	RGB [3]uint8
	Hex string
}

// Color 获取颜色值
func (r *RGB) Color() string {
	var hex string
	if r.Hex != "" {
		// 去除#
		hex = strings.ReplaceAll(r.Hex, "#", "")
		if len(hex) != 6 || hex < "000000" || hex > "FFFFFF" {
			panic("不支持的颜色格式")
		}

	} else if r.RGB != [3]uint8{} {
		for _, u := range r.RGB {
			result := strconv.FormatUint(uint64(u), 16)
			if len(result) == 1 {
				result = "0" + result
			}
			hex += result
		}
	} else {
		panic("未输入颜色字符")
	}
	return hex
}
