package shared

const (
	_TWIP_PER_INCH = 1440
	_TWIP_PER_CM   = 567
	_TWIP_PER_MM   = 5670
	_TWIP_PER_PT   = 20
	_TWIP_PER_Line = 313
	_TWIP_PER_Char = 100
)

// Twip xml长度单位
type Twip int

// Pt 英镑长度 1pt = 20twip
func Pt(pt float64) (twip Twip) {
	return Twip(pt * _TWIP_PER_PT)
}

// Cm 厘米 1cm = 566.929twip
func Cm(cm float64) (twip Twip) {
	return Twip(cm * _TWIP_PER_CM)
}

// Inch 英寸	1inch = 1440twip
func Inch(inch float64) (twip Twip) {
	return Twip(inch * _TWIP_PER_INCH)
}

// Mm 毫米	1mm = 5669.29twip
func Mm(mm float64) (twip Twip) {
	return Twip(mm * _TWIP_PER_MM)
}

// Wx 字号
//
//	word wps...显示的是磅数
//	example：
//		84号字 = 42pt = 42 * 20 twip = 84 * 10 twip
func Wx(wx int) (twip Twip) {
	return Twip(float64(wx / 2))
}

// Ln 行数 1ln = 313 twip
func Ln(ln float64) (twip Twip) {
	return Twip(ln * _TWIP_PER_Line)
}

// Char 字符
func Char(char int) (twip Twip) {
	return Twip(char * _TWIP_PER_Char)
}
