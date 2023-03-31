package styles

import (
	"encoding/xml"
	"strconv"
)

type RunProperties struct {
	XMLName   xml.Name `xml:"w:rPr,omitempty"`
	Fonts     *Fonts
	BCs       *BoldCs
	Bold      *Bold
	ICs       *ItalicCs
	Italic    *Italic
	Highlight *Highlight
	Color     *Color
	Size      *Size
	SizeCs    *SizeCs
	Spacing   *Spacing
	RunStyle  *RunStyle
	Style     *Style
	Shade     *Shade
	Kern      *Kern
	Underline *Underline
	VertAlign *VertAlign
}

// SetFont sets the font of the run
//
//	font []string = []string{"楷体"} 或
//		[]string{ASCII，EastAsia，HAnsi，Cs，Hint}
func (r *RunProperties) SetFont(font ...string) *RunProperties {
	if len(font) == 1 {
		r.Fonts = &Fonts{
			ASCII:    font[0],
			EastAsia: font[0],
			HAnsi:    font[0],
			Cs:       font[0],
			Hint:     "eastAsia",
		}
	} else {
		r.Fonts = &Fonts{
			ASCII:    font[0],
			EastAsia: font[1],
			HAnsi:    font[2],
			Cs:       font[3],
			Hint:     font[4],
		}
	}
	return r
}

// SetBold Bold ...
func (r *RunProperties) SetBold() *RunProperties {
	r.Bold = &Bold{}
	r.BCs = &BoldCs{}
	return r
}

// SetColor allows to set run color
//
//	color: [3]uint8{0-255,0-255,0-255} , 000000-FFFFFF
func (r *RunProperties) SetColor(color ...any) *RunProperties {
	var rgb RGB
	if len(color) == 1 {
		rgb = RGB{Hex: color[0].(string)}
	} else if len(color) == 3 {
		rgb = RGB{RGB: [3]uint8{uint8(color[0].(int)), uint8(color[1].(int)), uint8(color[2].(int))}}
	} else {
		panic("无效颜色")
	}
	r.Color = &Color{
		Val: rgb.Color(),
	}
	return r
}

// SetSize 字号
func (r *RunProperties) SetSize(size float64) *RunProperties {
	r.Size = &Size{
		Val: strconv.FormatFloat(size*2, 'G', -1, 64),
	}
	return r
}
