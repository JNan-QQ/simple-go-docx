package shared

const (
	_TWIP_PER_INCH = 1440
	_TWIP_PER_CM   = 566.929
	_TWIP_PER_MM   = 5669.29
	_TWIP_PER_PT   = 20
)

// Pt 英镑长度 1pt = 20twip
func Pt(pt float64) (twip int) {
	return int(pt * _TWIP_PER_PT)
}

// Cm 厘米 1cm = 566.929twip
func Cm(cm float64) (twip int) {
	return int(cm * _TWIP_PER_CM)
}

// Inch 英寸	1inch = 1440twip
func Inch(inch float64) (twip int) {
	return int(inch * _TWIP_PER_INCH)
}

// Mm 毫米	1mm = 5669.29twip
func Mm(mm float64) (twip int) {
	return int(mm * _TWIP_PER_MM)
}

// Wx 字体大小
//
//	word wps...显示的是磅数
//	example：
//		84号字 = 42pt = 42 * 20 twip = 84 * 10 twip
func Wx(wx int) (twip int) {
	return Pt(float64(wx / 2))
}

// Twip word中xml中的长度单位
func Twip(twip int) int {
	return twip
}
