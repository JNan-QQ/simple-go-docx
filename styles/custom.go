package styles

import (
	"encoding/xml"
	"gitee.com/jn-qq/simple-go-docx/shared"
	"strings"
)

// CustomStyle 自定义样式
type CustomStyle struct {
	XMLName        xml.Name `xml:"w:style"`
	Type           string   `xml:"w:type,attr,omitempty"`
	Id             int64    `xml:"w:styleId,attr,omitempty"`
	DefaultId      int64    `xml:"w:default,attr,omitempty"`
	Flg            string   `xml:"w:customStyle,attr,omitempty"`
	uiPriority     *uiPriority
	StyleName      *styleName
	Format         *format
	basedOn        *basedOn
	next           *next
	link           *link
	semiHidden     *semiHidden
	unhideWhenUsed *unhideWhenUsed
	ParagraphStyle *ParagraphProperties
	TextStyle      *RunProperties
}

// NewCustomStyle 创建自定义样式
//
//	styleType 样式类型，可选：character|paragraph|tab|...
func NewCustomStyle(name, styleType string) CustomStyle {
	return CustomStyle{
		Type:           styleType,
		Flg:            "1",
		StyleName:      &styleName{Val: name},
		Format:         &format{},
		ParagraphStyle: &ParagraphProperties{},
		TextStyle:      &RunProperties{},
	}
}

// DefaultStyle 添加默认段落样式
//
//	系统调用
func (c *CustomStyle) DefaultStyle() map[string]*CustomStyle {
	var defaultStyle = map[string]*CustomStyle{
		"Normal": {
			Type:       "paragraph",
			Id:         1,
			DefaultId:  1,
			StyleName:  &styleName{Val: "Normal"},
			Format:     &format{},
			uiPriority: &uiPriority{},
			ParagraphStyle: &ParagraphProperties{
				WidowControl:  &widowControl{},
				Justification: &justification{Val: shared.AlignArguments.Both},
			},
			TextStyle: &RunProperties{
				Fonts: &fonts{
					ASCII:    "Times DefaultStyle Roman",
					EastAsia: "宋体",
					HAnsi:    "Times DefaultStyle Roman",
					Cs:       "Times DefaultStyle Roman",
					Hint:     "Times DefaultStyle Roman",
				},
				Kern:   &kern{Val: 2},
				Size:   &fontSize{Val: 21},
				SizeCs: &fontSizeCs{Val: 24},
				Lang: &lang{
					Val:      "en-US",
					EastAsia: "zh-CN",
					Bidi:     "ar-SA",
				},
			},
		},
		"heading 1": {
			Type:       "paragraph",
			Id:         2,
			StyleName:  &styleName{Val: "heading 1"},
			basedOn:    &basedOn{Val: 1},
			next:       &next{Val: 1},
			Format:     &format{},
			uiPriority: &uiPriority{},
			ParagraphStyle: &ParagraphProperties{
				keepNext:  &keepNext{},
				keepLines: &keepLines{},
				Spacing: &spacing{
					Before:      340,
					BeforeLines: 0,
					BeforeAuto:  0,
					After:       330,
					AfterLines:  0,
					AfterAuto:   0,
					Line:        576,
					LineRule:    "auto",
				},
				outlineLvl: &outlineLvl{},
			},
			TextStyle: &RunProperties{
				Bold: &bold{},
				Kern: &kern{Val: 44},
				Size: &fontSize{Val: 44},
			},
		},
		"heading 2": {
			Type:           "paragraph",
			Id:             3,
			StyleName:      &styleName{Val: "heading 2"},
			basedOn:        &basedOn{Val: 1},
			next:           &next{Val: 1},
			link:           &link{Val: 10},
			semiHidden:     &semiHidden{},
			unhideWhenUsed: &unhideWhenUsed{},
			Format:         &format{},
			uiPriority:     &uiPriority{Val: 9},
			ParagraphStyle: &ParagraphProperties{
				keepNext:  &keepNext{},
				keepLines: &keepLines{},
				Spacing: &spacing{
					Before:      260,
					BeforeLines: 0,
					BeforeAuto:  0,
					After:       260,
					AfterLines:  0,
					AfterAuto:   0,
					Line:        413,
					LineRule:    "auto",
				},
				outlineLvl: &outlineLvl{Val: 1},
			},
			TextStyle: &RunProperties{
				Fonts: &fonts{
					ASCII:    "Arial",
					HAnsi:    "Arial",
					EastAsia: "黑体",
				},
				Bold: &bold{},
				Size: &fontSize{Val: 32},
			},
		},
		"heading 3": {
			Type:           "paragraph",
			Id:             4,
			StyleName:      &styleName{Val: "heading 3"},
			basedOn:        &basedOn{Val: 1},
			next:           &next{Val: 1},
			link:           &link{Val: 11},
			semiHidden:     &semiHidden{},
			unhideWhenUsed: &unhideWhenUsed{},
			Format:         &format{},
			uiPriority:     &uiPriority{Val: 9},
			ParagraphStyle: &ParagraphProperties{
				keepNext:  &keepNext{},
				keepLines: &keepLines{},
				Spacing: &spacing{
					Before:      260,
					BeforeLines: 0,
					BeforeAuto:  0,
					After:       260,
					AfterLines:  0,
					AfterAuto:   0,
					Line:        413,
					LineRule:    "auto",
				},
				outlineLvl: &outlineLvl{Val: 2},
			},
			TextStyle: &RunProperties{
				Bold: &bold{},
				Size: &fontSize{Val: 32},
			},
		},
		"heading 4": {
			Type:           "paragraph",
			Id:             5,
			StyleName:      &styleName{Val: "heading 4"},
			basedOn:        &basedOn{Val: 1},
			next:           &next{Val: 1},
			link:           &link{Val: 12},
			semiHidden:     &semiHidden{},
			unhideWhenUsed: &unhideWhenUsed{},
			Format:         &format{},
			uiPriority:     &uiPriority{Val: 9},
			ParagraphStyle: &ParagraphProperties{
				keepNext:  &keepNext{},
				keepLines: &keepLines{},
				Spacing: &spacing{
					Before:      280,
					BeforeLines: 0,
					BeforeAuto:  0,
					After:       290,
					AfterLines:  0,
					AfterAuto:   0,
					Line:        372,
					LineRule:    "auto",
				},
				outlineLvl: &outlineLvl{Val: 3},
			},
			TextStyle: &RunProperties{
				Fonts: &fonts{
					ASCII:    "Arial",
					HAnsi:    "Arial",
					EastAsia: "黑体",
				},
				Bold: &bold{},
				Size: &fontSize{Val: 28},
			},
		},
		"character": {
			Type:           "character",
			Id:             6,
			DefaultId:      1,
			StyleName:      &styleName{Val: "Default Paragraph Font"},
			semiHidden:     &semiHidden{},
			unhideWhenUsed: &unhideWhenUsed{},
			Format:         &format{},
			uiPriority:     &uiPriority{Val: 1},
		},
	}

	return defaultStyle
}

// RunProperties 文本样式
type RunProperties struct {
	XMLName   xml.Name `xml:"w:rPr,omitempty"`
	Fonts     *fonts
	BCs       *boldCs
	Bold      *bold
	ICs       *italicCs
	Italic    *italic
	Highlight *highlight
	Color     *colors
	Size      *fontSize
	SizeCs    *fontSizeCs
	Spacing   *spacing
	RStyle    *rStyle
	Shade     *shade
	Kern      *kern
	Underline *underline
	VertAlign *vertAlign
	Lang      *lang
}

// SetFont 设置文本字体
//
//	font：可以输入单个字体样式 或 依次属于ASCII，EastAsia，HAnsi，Cs，Hint样式
func (r *RunProperties) SetFont(font ...string) *RunProperties {
	if len(font) == 1 {
		r.Fonts = &fonts{
			ASCII:    font[0],
			EastAsia: font[0],
			HAnsi:    font[0],
			Cs:       font[0],
			Hint:     "eastAsia",
		}
	} else {
		r.Fonts = &fonts{
			ASCII:    font[0],
			EastAsia: font[1],
			HAnsi:    font[2],
			Cs:       font[3],
			Hint:     font[4],
		}
	}
	return r
}

// SetBold 设置粗体
func (r *RunProperties) SetBold() *RunProperties {
	r.Bold = &bold{}
	r.BCs = &boldCs{}
	return r
}

// SetColor 设置字体颜色
//
//	color: shared.RGB , 000000-FFFFFF, shared.ColorLib.Red
func (r *RunProperties) SetColor(color any) *RunProperties {
	var hexString string

	switch t := color.(type) {
	case string:
		// 输入16进制颜色值
		hexString = strings.ToUpper(strings.ReplaceAll(t, "#", ""))
		if len(hexString) != 6 || hexString < "000000" || hexString > "FFFFFF" {
			panic("不支持的颜色格式")
		}
	case shared.RGB:
		hexString = t.RGBColor()
	case [2]string:
		hexString = t[0]
	default:
		hexString = "FF0000"
	}
	r.Color = &colors{
		Val: hexString,
	}

	return r
}

// SetSize 字号
//
//	size: shared.Twip 可选：
//		shared.Cm() | shared.Mm() | shared.Pt() | shared.Inch() | shared.Wx()
func (r *RunProperties) SetSize(size shared.Twip) *RunProperties {
	r.Size = &fontSize{
		Val: size / 10,
	}
	return r
}

// HighlightColor 高亮显示
//
//	color: 可输入标准颜色英文字符或直接在调用颜色库 shared.ColorLib 中的颜色
func (r *RunProperties) HighlightColor(color any) *RunProperties {
	var hexString string

	switch t := color.(type) {
	case string:
		// 输入16进制颜色值
		hexString = t
	case [2]string:
		hexString = t[1]
	default:
		hexString = "white"
	}
	r.Highlight = &highlight{Val: hexString}

	return r
}

// SetItalic 设置斜体
func (r *RunProperties) SetItalic() *RunProperties {
	r.Italic = &italic{}
	r.ICs = &italicCs{}
	return r
}

// SetSpace 字符间距调整
//
//	size: twip = shared.Cm() | shared.Mm() | shared.Pt() | shared.Inch() | shared.Wx()
func (r *RunProperties) SetSpace(size shared.Twip) *RunProperties {
	r.Spacing = &spacing{Val: size}
	return r
}

// SetUnderLine 设置下划线
//
//	single	单下划线
//	words	仅为非空格字符添加下划线
//	double  双下划线
//	thick	粗下划线
//	dotted	点下划线
//	dottedHeavy	粗点下划线
//	dash	虚线下划线
//	dashedHeavy	粗虚线下划线
//	dashLong	长虚线下划线
//	dashLongHeavy	粗长虚线下划线
//	dotDash		虚线点下划线
//	dashDotHeavy	粗虚线点下划线
//	dotDotDash	破折号-点-点下划线
//	dashDotDotHeavy	粗破折号-点-点下划线
//	wave	波浪下划线
//	wavyHeavy	粗波浪下划线
//	wavyDouble	双重波浪下划线
//	none	无
func (r *RunProperties) SetUnderLine(name string) *RunProperties {
	r.Underline = &underline{Val: name}
	return r
}

// ParagraphProperties 段落样式
type ParagraphProperties struct {
	XMLName        xml.Name `xml:"w:pPr,omitempty"`
	Spacing        *spacing
	Ind            *ind
	Justification  *justification
	Shade          *shade
	Kern           *kern
	PStyle         *pStyle
	TextAlignment  *textAlignment
	AdjustRightInd *adjustRightInd
	SnapToGrid     *snapToGrid
	OverflowPunct  *overflowPunct
	WidowControl   *widowControl
	keepNext       *keepNext
	keepLines      *keepLines
	outlineLvl     *outlineLvl
	RPR            *RunProperties
}

// SetHead 设置标题级别
//
//	styleId：自定义样式或默认样式id，默认如下：
//		Normal   	：1
//		heading 1	：2
//		heading 2	：3
//		heading 3	：4
//		heading 4	：5
//		character	：6
func (p *ParagraphProperties) SetHead(styleId int64) *ParagraphProperties {
	p.PStyle = &pStyle{Val: styleId}
	return p
}

// IndFirst 首行缩进,默认2字符
func (p *ParagraphProperties) IndFirst() *ParagraphProperties {
	p.Ind = &ind{
		FirstLineChars: 2 * 100,
	}
	return p
}

// IndHang 悬挂缩进，默认2字符
func (p *ParagraphProperties) IndHang() *ParagraphProperties {
	p.Ind = &ind{
		HangingChars: 2 * 100,
	}
	return p
}

// IndCustom 自定义缩进
//
//	leftChars：左缩进字符数量 shared.Char()
//	rightChars：右缩进字符数量 shared.Char()
func (p *ParagraphProperties) IndCustom(leftChars, rightChars shared.Twip) *ParagraphProperties {
	p.Ind = &ind{
		LeftChars:  leftChars,
		RightChars: rightChars,
	}
	return p
}

// LineSpace 设置行间距
//
//	before：段前间距行数 shared.line()
//	after：段后间距行数 shared.line()
func (p *ParagraphProperties) LineSpace(before, after shared.Twip) *ParagraphProperties {
	p.Spacing = &spacing{
		BeforeLines: before,
		AfterLines:  after,
	}
	return p
}

// XLineSpace 设置x倍行距
func (p *ParagraphProperties) XLineSpace(x float64) *ParagraphProperties {
	p.Spacing = &spacing{
		Line:     int(240 * x),
		LineRule: "auto",
	}
	return p
}

// Align 段落对齐方式
//
//	left：左对齐。center：居中对齐。 right：右对齐。both：两端对齐。distribute：分散对齐。
//	shared.AlignArguments...
func (p *ParagraphProperties) Align(align string) *ParagraphProperties {
	p.Justification = &justification{Val: align}
	p.TextAlignment = &textAlignment{Val: "auto"}
	return p
}
