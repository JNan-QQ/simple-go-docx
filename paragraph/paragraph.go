package paragraph

import (
	"encoding/xml"
	"gitee.com/jn-qq/simple-go-docx/run"
	"gitee.com/jn-qq/simple-go-docx/styles"
)

// Paragraph <w:p>
type Paragraph struct {
	XMLName xml.Name `xml:"w:p,omitempty"`

	Style *styles.ParagraphProperties // 样式 w:pPr

	Texts []interface{}
}

// AddText 添加文本块
func (p *Paragraph) AddText(text string) *run.Run {
	r := &run.Run{
		Text:  text,
		Style: &styles.RunProperties{},
	}
	p.Texts = append(p.Texts, r)
	return r
}
