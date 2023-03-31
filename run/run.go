package run

import (
	"encoding/xml"
	"gitee.com/jn-qq/simple-go-docx/styles"
	"strconv"
)

type Run struct {
	XMLName xml.Name `xml:"w:r,omitempty"`
	Space   string   `xml:"xml:space,attr,omitempty"`

	Properties *Properties `xml:"w:rPr,omitempty"`

	InstrText string `xml:"w:instrText,omitempty"`

	Text string `xml:"w:t"`
}

// Properties w:rPr 文本块样式
type Properties struct {
	styles.RunProperties
}

// Color allows to set run color
// [3]uint8{0-255,0-255,0-255} , 000000-FFFFFF
func (r *Run) Color(color ...any) *Run {
	var rgb styles.RGB
	if len(color) == 1 {
		rgb = styles.RGB{Hex: color[0].(string)}
	} else if len(color) == 3 {
		rgb = styles.RGB{RGB: [3]uint8{color[0].(uint8), color[1].(uint8), color[2].(uint8)}}
	} else {
		panic("无效颜色")
	}
	r.Properties.Color = &styles.Color{
		Val: rgb.Color(),
	}
	return r
}

// Size 字号
func (r *Run) Size(size float64) *Run {
	r.Properties.Size = &styles.Size{
		Val: strconv.FormatFloat(size*2, 'G', -1, 64),
	}
	return r
}

// Shade allows to set run shade
func (r *Run) Shade(val, color, fill string) *Run {
	r.Properties.Shade = &styles.Shade{
		Val:   val,
		Color: color,
		Fill:  fill,
	}
	return r
}

// Bold ...
func (r *Run) Bold() *Run {
	r.Properties.Bold = &styles.Bold{}
	r.Properties.BCs = &styles.BoldCs{}
	return r
}

// Italic ...
func (r *Run) Italic() *Run {
	r.Properties.Italic = &styles.Italic{}
	r.Properties.ICs = &styles.ItalicCs{}
	return r
}

// Underline 可选如下：
//
//	dash - a dashed line
//	dashDotDotHeavy - a series of thick dash, dot, dot characters
//	dashDotHeavy - a series of thick dash, dot characters
//	dashedHeavy - a series of thick dashes
//	dashLong - a series of long dashed characters
//	dashLongHeavy - a series of thick, long, dashed characters
//	dotDash - a series of dash, dot characters
//	dotDotDash - a series of dash, dot, dot characters
//	dotted - a series of dot characters
//	dottedHeavy - a series of thick dot characters
//	double - two lines
//	none - no underline
//	single - a single line
//	thick - a single think line
//	wave - a single wavy line
//	wavyDouble - a pair of wavy lines
//	wavyHeavy - a single thick wavy line
//	words - a single line beneath all non-space characters
func (r *Run) Underline(val string) *Run {
	r.Properties.Underline = &styles.Underline{Val: val}
	return r
}

// Highlight ...
func (r *Run) Highlight(val string) *Run {
	r.Properties.Highlight = &styles.Highlight{Val: val}
	return r
}

// Font sets the font of the run
// font []string = []string{"楷体"} 或 []string{ASCII，EastAsia，HAnsi，Cs，Hint}
func (r *Run) Font(font ...string) *Run {
	if len(font) == 1 {
		r.Properties.Fonts = &styles.Fonts{
			ASCII:    font[0],
			EastAsia: font[0],
			HAnsi:    font[0],
			Cs:       font[0],
			Hint:     "eastAsia",
		}
	} else {
		r.Properties.Fonts = &styles.Fonts{
			ASCII:    font[0],
			EastAsia: font[1],
			HAnsi:    font[2],
			Cs:       font[3],
			Hint:     font[4],
		}
	}

	return r
}
