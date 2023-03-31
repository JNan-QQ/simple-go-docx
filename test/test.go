package main

import (
	"gitee.com/jn-qq/simple-go-docx"
	"path/filepath"
	"runtime"
)

func main() {
	// 创建docx文档对象
	document := docx.NewDocx()

	//添加段落
	p1 := document.AddParagraph()
	//设置段落格式
	p1.IndFirst().Head(1)
	// 添加文本
	r1 := p1.AddRun("测试所有字体格式")
	// 设置文本格式
	r1.Size(10).Color("FF0000").Font("楷体").
		Shade("clear", "auto", "E7E6E6").
		Bold().Italic().Underline("wave").
		Highlight("yellow")

	// 添加文本
	p1.AddRun("段落新增文本1")
	p1.AddRun("段落新增文本2")

	// 简单连写添加
	document.AddParagraph().Head(2).XSpce(2).AddRun("段落2文本").Color("FF0000").Font("楷体")

	// 自定义样式
	// 声明样式对象
	cs := document.NewCustomStyle("自定义段落样式", "paragraph")
	// 设置具体样式
	cs.CreatePPR().XSpce(2).IndFirst()
	cs.CreateRPR().SetFont("楷体", "楷体", "楷体", "楷体", "eastAsia").SetColor(255, 0, 0).SetSize(10)
	// 添加段落指定段落样式 Head 中的参数要-1
	document.AddParagraph().Head(cs.Id - 1).AddRun("自定义段落样式")

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
