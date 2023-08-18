package paragraph

import (
	"encoding/xml"
	"gitee.com/jn-qq/simple-go-docx/image"
	"gitee.com/jn-qq/simple-go-docx/run"
	"gitee.com/jn-qq/simple-go-docx/shared"
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

// AddDrawing 在段落中添加图片，可多张
//
//	name: 上传的图片名称,注意图片是否上传
//	width、height: 图片长宽
//	note：题注 默认下方居中（就是另一段落）
//	align: 图片对其方式 左(left)、右(right)、居中(center)
//	layout: 布局方式，默认为内嵌型
func (p *Paragraph) AddDrawing(name string, width, height shared.Twip, align string) *Paragraph {

	if align != "" {
		p.Style.Align(align)
	} else {
		p.Style.Align("center")
	}

	r := &run.Run{Image: &image.Draw{
		Inline: &image.Inline{
			Extent: &image.Extent{
				Cx: width.Emus().Int64(),
				Cy: width.Emus().Int64(),
			},
			Pic: &image.Pic{
				ImgId: &image.ImgId{RId: image.ImagesList[name].GetId()},
				ImgSize: &image.ImgSize{
					Cx: width.Emus().Int64(),
					Cy: width.Emus().Int64(),
				},
			},
		},
	}}
	p.Texts = append(p.Texts, r)
	return p
}
