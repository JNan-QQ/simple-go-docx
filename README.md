# simple-go-docx

#### 介绍
golang 生成简单docx文档，纯文本样式

<br>

#### 包结构

~~~mermaid
graph LR
    A(simple-go-docx)--package-->B((docx))
    A--dir-->C{{document}}--package-->D((document))
    A--dir-->E{{paragraph}}--package-->F((paragraph))
    A--dir-->G{{run}}--package-->H((run))
    A--dir-->I{{styles}}--package-->J((styles))
    A--dir-->K[(templates)]
~~~


#### 安装

```go
go get -u gitee.com/jn-qq/simple-go-docx
```

<br>

#### 使用说明

`simple-go-docx/docx_test.go`

```
    // 创建docx文档对象
	document := NewDocx()

	// 修改默认样式
	defaultStyle := document.GetStyle("Normal")
	defaultStyle.TextStyle.SetFont("楷体")

	//添加段落 ①
	p1 := document.AddParagraph()
	//设置段落格式
	p1.Style.IndFirst()
	// 添加文本 ①
	r1 := p1.AddText("测试所有字体格式")
	// 设置文本格式
	r1.Style.SetSize(shared.Pt(10)).SetColor("FF0000").SetFont("楷体").
		SetBold().SetItalic().SetUnderLine("wave").
		HighlightColor(shared.ColorLib.Yellow)

	// 添加文本 ②
	p1.AddText("段落新增文本1")
	p1.AddText("段落新增文本2")

	// 简单连写添加
	p2 := document.AddParagraph()
	p2.Style.SetHead(2).XLineSpace(2)
	p2.AddText("段落2文本").Style.SetColor("FF0000").SetFont("楷体")

	// 自定义样式
	style := styles.NewCustomStyle("自定义样式1", "paragraph")
	// 设置具体样式
	style.ParagraphStyle.IndFirst().XLineSpace(2)
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

	// 图片相关操作
	// 上传图片
	if err := document.UploadImages(10, &image.Image{
		Online: "https://tse4-mm.cn.bing.net/th/id/OIP-C.4UlvcR0AB1Oh_iXwP7szowHaGI",
		Name:   "image1",
	}); err != nil {
		panic(err)
	}
	p4 := document.AddParagraph()
	p4.AddDrawing("image1", shared.Cm(5), shared.Cm(5), "")
	p4.AddDrawing("image1", shared.Cm(5), shared.Cm(5), "")

	_, path, _, _ := runtime.Caller(0)
	path, _ = filepath.Split(path)

	err := document.Save(filepath.Join(path, "test.docx"))
	if err != nil {
		panic(err)
	}

```

<br>


#### 参考

1. [go-docx](https://github.com/fumiama/go-docx)

2. [ooxml](https://www.datypic.com/sc/ooxml)

    

