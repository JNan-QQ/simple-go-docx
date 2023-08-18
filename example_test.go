package docx

import (
	"gitee.com/jn-qq/simple-go-docx/image"
	"gitee.com/jn-qq/simple-go-docx/shared"
	"gitee.com/jn-qq/simple-go-docx/styles"
)

func ExampleDocx_AddParagraph() {
	// 创建docx文档对象
	document := NewDocx()
	defer document.Save("test.docx")

	// 声明段落 可添加格式文本...
	p1 := document.AddParagraph()
	p1.AddText("测试段落。")

}

func ExampleDocx_UploadImages() {
	// 创建docx文档对象
	document := NewDocx()
	defer document.Save("test.docx")

	// 上传图片，支持本地、网络及数据流,图片质量50
	if err := document.UploadImages(
		50,
		&image.Image{
			Online: "https://tse4-mm.cn.bing.net/th/id/OIP-C.4UlvcR0AB1Oh_iXwP7szowHaGI",
			Name:   "image1",
		},
		&image.Image{
			Local: ".../test.png",
			Name:  "image2",
		},
		&image.Image{
			Bytes: []byte{},
			Name:  "image3",
		},
	); err != nil {
		panic(err)
	}

	// 使用 创建段落 -> 映射图片 -> 设置题注。。。

	// 图片②③同行居中显示
	p1 := document.AddParagraph()
	p1.AddDrawing("image2", shared.Cm(5), shared.Cm(5), shared.AlignArguments.Center)
	p1.AddDrawing("image3", shared.Cm(5), shared.Cm(5), shared.AlignArguments.Center)

	// 题注
	t := document.AddParagraph()
	t.Style.Align(shared.AlignArguments.Center)
	t.AddText("图1 两个同行显示图").Style.SetSize(shared.Wx(10))
}

func ExampleDocx_AddCustomStyle() {
	// 创建docx文档对象
	document := NewDocx()
	defer document.Save("test.docx")

	// 创建自定义样式对象
	style := styles.NewCustomStyle("ct1", "paragraph")

	// 设置段落文本格式
	style.ParagraphStyle.IndFirst().XLineSpace(2)
	style.TextStyle.SetFont("楷体").SetSize(shared.Pt(20)).SetColor(shared.ColorLib.Blue)

	// 添加声明样式到文档 获取系统分配id
	sid := document.AddCustomStyle(&style)

	//添加段落指定段落样式 Head 中的参数要-1
	p := document.AddParagraph()
	p.Style.SetHead(sid)
	p.AddText("自定义段落样式")
}

func ExampleDocx_GetStyle() {
	// 创建docx文档对象
	document := NewDocx()
	defer document.Save("test.docx")

	// 获取默认样式 Normal | heading 1-4 | character
	defaultStyle := document.GetStyle("Normal")
	defaultStyle.TextStyle.SetFont("楷体")

	// 声明段落 可添加格式文本...
	p1 := document.AddParagraph()
	p1.AddText("修改默认文本字体。")
}
