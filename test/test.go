package main

import (
	"gitee.com/jn-qq/simple-go-docx"
	"gitee.com/jn-qq/simple-go-docx/shared"
	"gitee.com/jn-qq/simple-go-docx/styles"
	"path/filepath"
	"runtime"
)

func main() {
	// 创建docx文档对象
	document := docx.NewDocx()

	// 修改默认样式
	defaultStyle := document.GetStyle("Normal")
	defaultStyle.TextStyle.SetFont("楷体")

	//添加段落
	p1 := document.AddParagraph()
	//设置段落格式
	p1.Style.IndFirst()
	// 添加文本
	r1 := p1.AddText("测试所有字体格式")
	// 设置文本格式
	r1.Style.SetSize(shared.Pt(10)).SetColor("FF0000").SetFont("楷体").
		SetBold().SetItalic().SetUnderLine("wave").
		HighlightColor(shared.ColorLib.Yellow)

	// 添加文本
	p1.AddText("段落新增文本1")
	p1.AddText("段落新增文本2")

	// 简单连写添加
	p2 := document.AddParagraph()
	p2.Style.SetHead(2).XLineSpce(2)
	p2.AddText("段落2文本").Style.SetColor("FF0000").SetFont("楷体")

	// 自定义样式
	style := styles.NewCustomStyle("自定义样式1", "paragraph")
	// 设置具体样式
	style.ParagraphStyle.IndFirst().XLineSpce(2)
	style.TextStyle.SetFont("楷体").SetSize(shared.Pt(20)).SetColor(shared.ColorLib.Blue)
	// 添加声明样式 获取id
	sid := document.AddCustomStyle(&style)

	// 添加段落指定段落样式
	p3 := document.AddParagraph()
	p3.Style.SetHead(sid)
	p3.AddText("自定义段落样式")

	// 添加字符样式
	//cs := docx.NewCustomStyle("自定义段落样式", "character")
	//添加属性。。。

	_, path, _, _ := runtime.Caller(0)
	path, _ = filepath.Split(path)

	err := document.Save(filepath.Join(path, "test.docx"))
	if err != nil {
		panic(err)
	}

}
