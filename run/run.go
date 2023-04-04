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