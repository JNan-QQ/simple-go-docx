package styles

import (
	"encoding/xml"
)

// Fonts 字体类型
type Fonts struct {
	XMLName  xml.Name `xml:"w:rFonts,omitempty"`
	ASCII    string   `xml:"w:ascii,attr,omitempty"`
	EastAsia string   `xml:"w:eastAsia,attr,omitempty"`
	HAnsi    string   `xml:"w:hAnsi,attr,omitempty"`
	Cs       string   `xml:"w:cs,attr,omitempty"`
	Hint     string   `xml:"w:hint,attr,omitempty"`
}

// Spacing 行距
type Spacing struct {
	XMLName xml.Name `xml:"w:spacing,omitempty"`

	Val int `xml:"w:val,attr,omitempty"`

	BeforeLines int    `xml:"w:beforeLines,attr,omitempty"`
	Before      int    `xml:"w:before,attr,omitempty"`
	AfterLines  int    `xml:"w:afterLines,attr,omitempty"`
	After       int    `xml:"w:after,attr,omitempty"`
	Line        int    `xml:"w:line,attr,omitempty"`
	LineRule    string `xml:"w:lineRule,attr,omitempty"`
}

// Ind 缩进
type Ind struct {
	XMLName xml.Name `xml:"w:ind,omitempty"`

	LeftChars      int `xml:"w:leftChars,attr,omitempty"`
	Left           int `xml:"w:left,attr,omitempty"`
	FirstLineChars int `xml:"w:firstLineChars,attr,omitempty"`
	FirstLine      int `xml:"w:firstLine,attr,omitempty"`
	HangingChars   int `xml:"w:hangingChars,attr,omitempty"`
	Hanging        int `xml:"w:hanging,attr,omitempty"`
}

// Justification 对齐方式
//
//	w:jc 属性的取值可以是以下之一：
//		start：左对齐。
//		center：居中对齐。
//		end：右对齐。
//		both：两端对齐。
//		distribute：分散对齐。
type Justification struct {
	XMLName xml.Name `xml:"w:jc,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// TextAlignment ...
type TextAlignment struct {
	XMLName xml.Name `xml:"w:textAlignment,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// VertAlign 垂直对齐
type VertAlign struct {
	XMLName xml.Name `xml:"w:vertAlign,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// Shade 阴影
type Shade struct {
	XMLName       xml.Name `xml:"w:shd,omitempty"`
	Val           string   `xml:"w:val,attr,omitempty"`
	Color         string   `xml:"w:color,attr,omitempty"`
	Fill          string   `xml:"w:fill,attr,omitempty"`
	ThemeFill     string   `xml:"w:themeFill,attr,omitempty"`
	ThemeFillTint string   `xml:"w:themeFillTint,attr,omitempty"`
}

// RunStyle 包含文本块样式
type RunStyle struct {
	XMLName xml.Name `xml:"w:rStyle,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// Style contains styling for a paragraph
type Style struct {
	XMLName xml.Name `xml:"w:pStyle,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// Color 颜色
type Color struct {
	XMLName xml.Name `xml:"w:color,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// Size 大小
type Size struct {
	XMLName xml.Name `xml:"w:sz,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// SizeCs 复杂样式大小
type SizeCs struct {
	XMLName xml.Name `xml:"w:szCs,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// Bold 加粗
type Bold struct {
	XMLName xml.Name `xml:"w:b,omitempty"`
}

// BoldCs 复杂格式加粗
type BoldCs struct {
	XMLName xml.Name `xml:"w:bCs,omitempty"`
}

// Italic 斜体
type Italic struct {
	XMLName xml.Name `xml:"w:i,omitempty"`
}

// ItalicCs 复杂格式斜体
type ItalicCs struct {
	XMLName xml.Name `xml:"w:iCs,omitempty"`
}

// Underline 下划线
type Underline struct {
	XMLName xml.Name `xml:"w:u,omitempty"`
	Val     string   `xml:"w:val,attr,omitempty"`
}

// Highlight 高亮
type Highlight struct {
	XMLName xml.Name `xml:"w:highlight,omitempty"`
	Val     string   `xml:"w:val,attr,omitempty"`
}

// Kern ...
type Kern struct {
	XMLName xml.Name `xml:"w:kern,omitempty"`
	Val     int64    `xml:"w:val,attr"`
}

// AdjustRightInd ...
type AdjustRightInd struct {
	XMLName xml.Name `xml:"w:adjustRightInd,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// SnapToGrid ...
type SnapToGrid struct {
	XMLName xml.Name `xml:"w:snapToGrid,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// Kinsoku ...
type Kinsoku struct {
	XMLName xml.Name `xml:"w:kinsoku,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

// OverflowPunct ...
type OverflowPunct struct {
	XMLName xml.Name `xml:"w:overflowPunct,omitempty"`
	Val     int      `xml:"w:val,attr"`
}

type Name struct {
	XMLName xml.Name `xml:"w:name,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// w:qFormat
type Format struct {
	XMLName xml.Name `xml:"w:qFormat,omitempty"`
}
