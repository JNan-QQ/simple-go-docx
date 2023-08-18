package paragraph

import (
	docx "gitee.com/jn-qq/simple-go-docx"
	"gitee.com/jn-qq/simple-go-docx/image"
	"gitee.com/jn-qq/simple-go-docx/shared"
	"gitee.com/jn-qq/simple-go-docx/styles"
)

func ExampleParagraph_AddText() {
	p := &Paragraph{
		Texts: make([]interface{}, 0, 64),
		Style: &styles.ParagraphProperties{},
	}

	// 设置段落格式
	p.Style.
		Align("center"). // 居中
		XLineSpace(1.5). // 1.5倍行距
		IndFirst()       // 首行缩进2字符

	// 添加文字
	r1 := p.AddText("文本1")
	r1.Style.SetSize(shared.Pt(20)).SetColor(shared.ColorLib.Blue)
	p.AddText("文本2").Style.SetSize(shared.Wx(10)).SetBold().SetFont("黑体")
}

func ExampleParagraph_AddDrawing() {
	document := docx.NewDocx()

	// 上传图片
	if err := document.UploadImages(100, []*image.Image{}...); err != nil {
		panic(err)
	}

	// 映射图片
	document.AddParagraph().AddDrawing("...", shared.Cm(10), shared.Cm(5), "right")
}
