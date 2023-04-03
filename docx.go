package docx

import (
	"archive/zip"
	"encoding/xml"
	"gitee.com/jn-qq/simple-go-docx/document"
	"gitee.com/jn-qq/simple-go-docx/paragraph"
	"gitee.com/jn-qq/simple-go-docx/styles"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

type Docx struct {
	Document document.Document // word/document.xml 文档内容

	docRelation document.Relationships // word/_rels/document.xml.rels 各文件之间的关系

	styleId int64

	templateType     string
	workDir          string
	templateFileList []string

	defaultStyle map[string]*styles.CustomStyle
}

// NewDocx 生成一个新的空 A4 docx 文件，我们可以对其进行操作和保存
func NewDocx() *Docx {
	docx := &Docx{
		Document: document.Document{
			XMLName: xml.Name{Space: "w"},
			XmlW:    xmlnsW,
			XmlR:    xmlnsR,
			XmlWp:   xmlnsWp,
			XmlWps:  xmlnsWps,
			XmlWpc:  xmlnsWpc,
			XmlWpg:  xmlnsWpg,
			Body:    document.Body{Items: make([]interface{}, 0, 64)},
		},
		docRelation: document.Relationships{
			Xmlns: xmlnsRel,
			Relationship: []document.Relationship{
				{
					ID:     "rId1",
					Type:   relationshipId1,
					Target: "styles.xml",
				},
				{
					ID:     "rId2",
					Type:   relationshipId2,
					Target: "theme/theme1.xml",
				},
				{
					ID:     "rId3",
					Type:   relationshipId3,
					Target: "fontTable.xml",
				},
			},
		},
		styleId:      7,
		templateType: "A4",
		templateFileList: []string{
			"_rels/.rels",
			"docProps/app.xml",
			"docProps/core.xml",
			"word/theme/theme1.xml",
			"word/fontTable.xml",
			"word/styles.xml",
			"[Content_Types].xml",
		},
		defaultStyle: new(styles.CustomStyle).DefaultStyle(),
	}

	return docx
}

// AddParagraph 添加段落,可对该段落设置样式
func (d *Docx) AddParagraph() *paragraph.Paragraph {
	p := &paragraph.Paragraph{
		Texts: make([]interface{}, 0, 64),
		Style: &styles.ParagraphProperties{},
	}
	d.Document.Body.Items = append(d.Document.Body.Items, p)

	// 获取当前文件路径
	_, path, _, _ := runtime.Caller(0)
	d.workDir, _ = filepath.Split(path)

	return p
}

// Save 保存docx文档
//
//	savePath：文档保存路径，仅支持.docx文档
func (d *Docx) Save(savePath string) error {
	if _, err := os.Stat(savePath); err == nil {
		// 删除旧文件
		if err := os.Remove(savePath); err != nil {
			return err
		}
	}
	// 创建文件
	docx, _ := os.Create(savePath)
	defer func(docx *os.File) {
		if err := docx.Close(); err != nil {
			panic(err.Error())
		}
	}(docx)
	// 创建zip写入对象
	zipWriter := zip.NewWriter(docx)
	defer func(zipWriter *zip.Writer) {
		if err := zipWriter.Close(); err != nil {
			panic(err.Error())
		}
	}(zipWriter)

	// 添加模板文件
	for _, path := range d.templateFileList {
		file, err := os.ReadFile(filepath.Join(d.workDir, "templates", d.templateType, path))
		if err != nil {
			return err
		}
		if path == "word/styles.xml" {
			file = d.replaceNode(file)
		}
		d.addFileToZip(zipWriter, path, file)
	}

	// 添加 word/_rels/document.xml.rels
	if marshal, err := xml.Marshal(d.docRelation); err != nil {
		return err
	} else {
		d.addFileToZip(zipWriter, "word/_rels/document.xml.rels", marshal)
	}

	// 添加 word/document.xml
	if marshal, err := xml.Marshal(d.Document); err != nil {
		return err
	} else {
		d.addFileToZip(zipWriter, "word/document.xml", marshal)
	}

	return nil
}

// AddCustomStyle 添加自定义样式
//
//	styleType 样式类型，可选：character|paragraph|tab|...
//
//	example:
//		// 创建自定义样式对象
//		style := styles.NewCustomStyle("自定义样式1", "paragraph")
//		style.ParagraphStyle.IndFirst().XLineSpce(2)
//		style.TextStyle.SetFont("楷体").SetSize(shared.Pt(20)).SetColor(shared.ColorLib.Blue)
//		// 添加声明样式 获取id
//		sid := document.AddCustomStyle(style)
//
//		//添加段落指定段落样式 Head 中的参数要-1
//		p := document.AddParagraph()
//		p.Style.SetHead(sid - 1)
//		p.AddText("自定义段落样式")
func (d *Docx) AddCustomStyle(style *styles.CustomStyle) int64 {
	style.Id = d.styleId
	defer func() { d.styleId += 1 }()
	d.defaultStyle[style.StyleName.Val] = style

	return d.styleId
}

// GetStyle 获取已存在的样式
func (d *Docx) GetStyle(styleName string, styleType ...string) *styles.CustomStyle {
	if style, ok := d.defaultStyle[styleName]; ok {
		return style
	} else {
		panic("样式不存在请创建！")
	}
}

// 生成样式id
func (d *Docx) createStyleId() (id int64) {
	id = d.styleId
	d.styleId += 1
	return
}

// 向zip对象中添加文件
func (d *Docx) addFileToZip(zipWriter *zip.Writer, zipFilePath string, streamOrPath any) {

	//创建压缩文件对象
	f, err := zipWriter.Create(zipFilePath)
	if err != nil {
		panic(err.Error())
	}

	// 写入
	switch t := streamOrPath.(type) {
	case string:
		p, _ := os.ReadFile(t)
		if _, err = f.Write(p); err != nil {
			panic(err.Error())
		}
	case []byte:
		if _, err = f.Write(t); err != nil {
			panic(err.Error())
		}
	}
}

// 替换xml-node节点
func (d *Docx) replaceNode(zipFile []byte) []byte {

	var customStyle []byte
	// 读取自定义样式
	for _, cs := range d.defaultStyle {
		if marshal, err := xml.Marshal(cs); err == nil {
			customStyle = append(customStyle, marshal...)
		}
	}

	//定义一个正则表达式re 匹配<w:body></w:body>
	re := regexp.MustCompile(`<w:customStyle>customStyle</w:customStyle>`)

	// 替换
	body := re.ReplaceAll(zipFile, customStyle)

	return body
}

//nolint:revive,style-check
const (
	xmlnsW   = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	xmlnsR   = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
	xmlnsWp  = `http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing`
	xmlnsWps = `http://schemas.microsoft.com/office/word/2010/wordprocessingShape`
	xmlnsWpc = `http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas`
	xmlnsWpg = `http://schemas.microsoft.com/office/word/2010/wordprocessingGroup`

	xmlnsRel = `http://schemas.openxmlformats.org/package/2006/relationships`

	relationshipId1 = `http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles`
	relationshipId2 = `http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme`
	relationshipId3 = `http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable`
)
