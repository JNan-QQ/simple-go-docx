package styles

import (
	"encoding/xml"
)

// fonts 字体类型
type fonts struct {
	XMLName  xml.Name `xml:"w:rFonts,omitempty"`
	ASCII    string   `xml:"w:ascii,attr,omitempty"`
	EastAsia string   `xml:"w:eastAsia,attr,omitempty"`
	HAnsi    string   `xml:"w:hAnsi,attr,omitempty"`
	Cs       string   `xml:"w:cs,attr,omitempty"`
	Hint     string   `xml:"w:hint,attr,omitempty"`
}

// spacing 行距
type spacing struct {
	XMLName xml.Name `xml:"w:spacing,omitempty"`

	// 文本字符间距调整
	Val int `xml:"w:val,attr,omitempty"`

	// 段落
	Before            int `xml:"w:before,attr,omitempty"`
	BeforeLines       int `xml:"w:beforeLines,attr,omitempty"`
	BeforeAutospacing int `xml:"w:beforeAutospacing,attr,omitempty"`

	After            int `xml:"w:after,attr,omitempty"`
	AfterLines       int `xml:"w:afterLines,attr,omitempty"`
	AfterAutospacing int `xml:"w:afterAutospacing,attr,omitempty"`

	Line     int    `xml:"w:line,attr,omitempty"`
	LineRule string `xml:"w:lineRule,attr,omitempty"`
}

// ind 缩进
type ind struct {
	XMLName xml.Name `xml:"w:ind,omitempty"`

	// 左缩进
	LeftChars int `xml:"w:leftChars,attr,omitempty"`
	Left      int `xml:"w:left,attr,omitempty"`
	// 右缩进
	RightChars int `xml:"w:rightChars,attr,omitempty"`
	Right      int `xml:"w:right,attr,omitempty"`
	// 首行缩进
	FirstLineChars int `xml:"w:firstLineChars,attr,omitempty"`
	FirstLine      int `xml:"w:firstLine,attr,omitempty"`
	// 悬挂缩进
	HangingChars int `xml:"w:hangingChars,attr,omitempty"`
	Hanging      int `xml:"w:hanging,attr,omitempty"`
}

// justification 对齐方式
//
//	w:jc 属性的取值可以是以下之一：
//		start：左对齐。
//		center：居中对齐。
//		end：右对齐。
//		both：两端对齐。
//		distribute：分散对齐。
type justification struct {
	XMLName xml.Name `xml:"w:jc,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// textAlignment 文本对其方式
type textAlignment struct {
	XMLName xml.Name `xml:"w:textAlignment,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// vertAlign 垂直对齐
type vertAlign struct {
	XMLName xml.Name `xml:"w:vertAlign,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// lang 垂直对齐
type lang struct {
	XMLName  xml.Name `xml:"w:lang,omitempty"`
	Val      string   `xml:"w:val,attr"`
	EastAsia string   `xml:"w:eastAsia,attr,omitempty"`
	Bidi     string   `xml:"w:bidi,attr"`
}

// shade 阴影纹理
type shade struct {
	XMLName       xml.Name `xml:"w:shd,omitempty"`
	Val           string   `xml:"w:val,attr,omitempty"`
	Color         string   `xml:"w:color,attr,omitempty"`
	Fill          string   `xml:"w:fill,attr,omitempty"`
	ThemeFill     string   `xml:"w:themeFill,attr,omitempty"`
	ThemeFillTint string   `xml:"w:themeFillTint,attr,omitempty"`
}

// rStyle 包含文本块样式
type rStyle struct {
	XMLName xml.Name `xml:"w:rStyle,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// pStyle 段落指定styleId
type pStyle struct {
	XMLName xml.Name `xml:"w:pStyle,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// colors 颜色
type colors struct {
	XMLName xml.Name `xml:"w:color,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// fontSize 大小
type fontSize struct {
	XMLName xml.Name `xml:"w:sz,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// fontSizeCs 复杂样式大小
type fontSizeCs struct {
	XMLName xml.Name `xml:"w:szCs,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// bold 加粗
type bold struct {
	XMLName xml.Name `xml:"w:b,omitempty"`
}

// boldCs 复杂格式加粗
type boldCs struct {
	XMLName xml.Name `xml:"w:bCs,omitempty"`
}

// italic 斜体
type italic struct {
	XMLName xml.Name `xml:"w:i,omitempty"`
}

// italicCs 复杂格式斜体
type italicCs struct {
	XMLName xml.Name `xml:"w:iCs,omitempty"`
}

// underline 下划线
type underline struct {
	XMLName xml.Name `xml:"w:u,omitempty"`
	Val     string   `xml:"w:val,attr,omitempty"`
}

// highlight 高亮
type highlight struct {
	XMLName xml.Name `xml:"w:highlight,omitempty"`
	Val     string   `xml:"w:val,attr,omitempty"`
}

// kern 字体字距
type kern struct {
	XMLName xml.Name `xml:"w:kern,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// adjustRightInd 使用文档网格时自动调整右缩进
type adjustRightInd struct {
	XMLName xml.Name `xml:"w:adjustRightInd,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// snapToGrid 与不可见的网格线对齐
type snapToGrid struct {
	XMLName xml.Name `xml:"w:snapToGrid,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// overflowPunct 允许标点符号越界
type overflowPunct struct {
	XMLName xml.Name `xml:"w:overflowPunct,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// widowControl 允许标点符号越界
type widowControl struct {
	XMLName xml.Name `xml:"w:widowControl,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// keepNext
type keepNext struct {
	XMLName xml.Name `xml:"w:keepNext,omitempty"`
}

// keepLines
type keepLines struct {
	XMLName xml.Name `xml:"w:keepLines,omitempty"`
}

// keepLines
type outlineLvl struct {
	XMLName xml.Name `xml:"w:outlineLvl,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// StyleName 自定义格式名称
type StyleName struct {
	XMLName xml.Name `xml:"w:name,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// w:qFormat
type Format struct {
	XMLName xml.Name `xml:"w:qFormat,omitempty"`
}

type uiPriority struct {
	XMLName xml.Name `xml:"w:uiPriority,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// 样式基于。。。
type basedOn struct {
	XMLName xml.Name `xml:"w:basedOn,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// 后文样式
type next struct {
	XMLName xml.Name `xml:"w:next,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

type link struct {
	XMLName xml.Name `xml:"w:link,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

type semiHidden struct {
	XMLName xml.Name `xml:"w:semiHidden,omitempty"`
}

type unhideWhenUsed struct {
	XMLName xml.Name `xml:"w:unhideWhenUsed,omitempty"`
}
