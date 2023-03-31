package styles

import "encoding/xml"

type ParagraphProperties struct {
	XMLName        xml.Name `xml:"w:pPr,omitempty"`
	Spacing        *Spacing
	Ind            *Ind
	Justification  *Justification
	Shade          *Shade
	Kern           *Kern
	Style          *Style
	TextAlignment  *TextAlignment
	AdjustRightInd *AdjustRightInd
	SnapToGrid     *SnapToGrid
	Kinsoku        *Kinsoku
	OverflowPunct  *OverflowPunct
}

// IndFirst 首行缩进
func (p *ParagraphProperties) IndFirst() *ParagraphProperties {
	p.Ind = &Ind{
		FirstLineChars: 200,
		FirstLine:      420,
	}
	return p
}

// IndCustom 自定义缩进
func (p *ParagraphProperties) IndCustom(ind Ind) *ParagraphProperties {
	p.Ind = &ind
	return p
}

// LineSpce 设置行间距
func (p *ParagraphProperties) LineSpce(before, after int) *ParagraphProperties {
	p.Spacing = &Spacing{
		Before:      before * 313,
		BeforeLines: before * 100,
		After:       after * 313,
		AfterLines:  after * 100,
	}
	return p
}

// XSpce 设置x倍行距
func (p *ParagraphProperties) XSpce(x float64) *ParagraphProperties {
	p.Spacing = &Spacing{
		Line:     int(240 * x),
		LineRule: "auto",
	}
	return p
}

// Align 对齐方式
// left：左对齐。center：居中对齐。 right：右对齐。both：两端对齐。distribute：分散对齐。
func (p *ParagraphProperties) Align(align string) *ParagraphProperties {
	p.Justification = &Justification{Val: align}
	p.TextAlignment = &TextAlignment{Val: "auto"}
	return p
}
