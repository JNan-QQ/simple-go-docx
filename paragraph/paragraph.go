package paragraph

import (
	"encoding/xml"
	"gitee.com/jn-qq/simple-go-docx/run"
	"gitee.com/jn-qq/simple-go-docx/styles"
	"strconv"
)

// Paragraph <w:p>
type Paragraph struct {
	XMLName xml.Name `xml:"w:p,omitempty"`

	Properties *Properties // 样式 w:pPr
	Items      []interface{}
}

// Properties <w:pPr> 段落样式
type Properties struct {
	styles.ParagraphProperties

	RunProperties *run.Properties
}

// AddRun 添加文本块
func (p *Paragraph) AddRun(text string) *run.Run {
	r := &run.Run{
		Text:       text,
		Properties: &run.Properties{},
	}
	p.Items = append(p.Items, r)
	return r
}

// Head 设置标题
func (p *Paragraph) Head(lv int64) *Paragraph {
	p.Properties.Style = &styles.Style{Val: strconv.FormatInt(lv+1, 10)}
	return p
}

// IndFirst 首行缩进
func (p *Paragraph) IndFirst() *Paragraph {
	p.Properties.Ind = &styles.Ind{
		FirstLineChars: 200,
		FirstLine:      420,
	}
	return p
}

// IndCustom 自定义缩进
func (p *Paragraph) IndCustom(ind styles.Ind) *Paragraph {
	p.Properties.Ind = &ind
	return p
}

// LineSpce 设置行间距
func (p *Paragraph) LineSpce(before, after int) *Paragraph {
	p.Properties.Spacing = &styles.Spacing{
		Before:      before * 313,
		BeforeLines: before * 100,
		After:       after * 313,
		AfterLines:  after * 100,
	}
	return p
}

// XSpce 设置x倍行距
func (p *Paragraph) XSpce(x float64) *Paragraph {
	p.Properties.Spacing = &styles.Spacing{
		Line:     int(240 * x),
		LineRule: "auto",
	}
	return p
}

// Align 对齐方式
// left：左对齐。center：居中对齐。 right：右对齐。both：两端对齐。distribute：分散对齐。
func (p *Paragraph) Align(align string) *Paragraph {
	p.Properties.Justification = &styles.Justification{Val: align}
	p.Properties.TextAlignment = &styles.TextAlignment{Val: "auto"}
	return p
}
