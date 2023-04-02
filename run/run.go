package run

import (
	"encoding/xml"
	"gitee.com/jn-qq/simple-go-docx/styles"
)

type Run struct {
	XMLName xml.Name `xml:"w:r,omitempty"`
	Space   string   `xml:"xml:space,attr,omitempty"`

	Style *styles.RunProperties `xml:"w:rPr,omitempty"`

	InstrText string `xml:"w:instrText,omitempty"`

	Text string `xml:"w:t"`
}

// SetText 设置文本
func (r *Run) SetText(text string) *Run {
	r.Text = text
	return r
}

//// colors allows to set run color
//// [3]uint8{0-255,0-255,0-255} , 000000-FFFFFF
//func (r *Run) colors(color ...any) *Run {
//	var rgb styles.RGB
//	if len(color) == 1 {
//		rgb = styles.RGB{Hex: color[0].(string)}
//	} else if len(color) == 3 {
//		rgb = styles.RGB{RGB: [3]uint8{uint8(color[0].(int)), uint8(color[1].(int)), uint8(color[2].(int))}}
//	} else {
//		panic("无效颜色")
//	}
//	r.Style.colors = &styles.colors{
//		Val: rgb.colors(),
//	}
//	return r
//}
//
//// fontSize 字号
////func (r *Run) fontSize(size float64) *Run {
////	r.Style.fontSize = &styles.fontSize{
////		Val: strconv.FormatFloat(size*2, 'G', -1, 64),
////	}
////	return r
////}
//
//// shade allows to set run shade
//func (r *Run) shade(val, color, fill string) *Run {
//	r.Style.shade = &styles.shade{
//		Val:   val,
//		colors: color,
//		Fill:  fill,
//	}
//	return r
//}
//
//// bold ...
//func (r *Run) bold() *Run {
//	r.Style.bold = &styles.bold{}
//	r.Style.BCs = &styles.boldCs{}
//	return r
//}
//
//// italic ...
//func (r *Run) italic() *Run {
//	r.Style.italic = &styles.italic{}
//	r.Style.ICs = &styles.italicCs{}
//	return r
//}
//
//// underline 可选如下：
////
////	dash - a dashed line
////	dashDotDotHeavy - a series of thick dash, dot, dot characters
////	dashDotHeavy - a series of thick dash, dot characters
////	dashedHeavy - a series of thick dashes
////	dashLong - a series of long dashed characters
////	dashLongHeavy - a series of thick, long, dashed characters
////	dotDash - a series of dash, dot characters
////	dotDotDash - a series of dash, dot, dot characters
////	dotted - a series of dot characters
////	dottedHeavy - a series of thick dot characters
////	double - two lines
////	none - no underline
////	single - a single line
////	thick - a single think line
////	wave - a single wavy line
////	wavyDouble - a pair of wavy lines
////	wavyHeavy - a single thick wavy line
////	words - a single line beneath all non-space characters
//func (r *Run) underline(val string) *Run {
//	r.Style.underline = &styles.underline{Val: val}
//	return r
//}
//
//// highlight ...
//func (r *Run) highlight(val string) *Run {
//	r.Style.highlight = &styles.highlight{Val: val}
//	return r
//}
//
//// Font sets the font of the run
//// font []string = []string{"楷体"} 或 []string{ASCII，EastAsia，HAnsi，Cs，Hint}
//func (r *Run) Font(font ...string) *Run {
//	if len(font) == 1 {
//		r.Style.fonts = &styles.fonts{
//			ASCII:    font[0],
//			EastAsia: font[0],
//			HAnsi:    font[0],
//			Cs:       font[0],
//			Hint:     "eastAsia",
//		}
//	} else {
//		r.Style.fonts = &styles.fonts{
//			ASCII:    font[0],
//			EastAsia: font[1],
//			HAnsi:    font[2],
//			Cs:       font[3],
//			Hint:     font[4],
//		}
//	}
//
//	return r
//}
