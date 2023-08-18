package docx

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"gitee.com/jn-qq/simple-go-docx/document"
	"gitee.com/jn-qq/simple-go-docx/image"
	"gitee.com/jn-qq/simple-go-docx/paragraph"
	"gitee.com/jn-qq/simple-go-docx/styles"
	"os"
	"regexp"
)

type Docx struct {
	Document document.Document // word/document.xml 文档内容

	docRelation document.Relationships // word/_rels/document.xml.rels 各文件之间的关系

	styleId int64

	templateType     string
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
			XmlPic:  xmlnsPic,
			XmlA:    xmlnsA,
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
		styleId:          7,
		templateType:     "A4",
		templateFileList: templateFiles,
		defaultStyle:     new(styles.CustomStyle).DefaultStyle(),
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

	return p
}

// UploadImages 批量上传图片
//
//	q: 图片质量 1-100，默认100
func (d *Docx) UploadImages(quality int, imgList ...*image.Image) error {

	if quality < 1 || quality > 100 {
		quality = 100
	}

	for _, img := range imgList {
		if img.Name == "" {
			return fmt.Errorf("请为图片添加名称标识")
		}
		// 格式化图片二进制数据
		if err := img.Down(quality); err != nil {
			return err
		}
		// 设置关系id
		img.SetId(fmt.Sprintf("rId%d", len(d.docRelation.Relationship)+1))

		// 保存数据
		image.ImagesList[img.Name] = img

		// 写入关系映射
		d.docRelation.Relationship = append(
			d.docRelation.Relationship,
			document.Relationship{
				ID:     img.GetId(),
				Type:   imageRelationship,
				Target: "media/" + img.Name + ".jpeg",
			},
		)
	}
	return nil
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
		err := docx.Close()
		if err != nil {
			panic("docx文档未正常关闭！")
		}
	}(docx)

	// 创建zip写入对象
	zipWriter := zip.NewWriter(docx)
	defer func(zipWriter *zip.Writer) {
		if err := zipWriter.Close(); err != nil {
			panic("压缩包文件未正常关闭")
		}
	}(zipWriter)

	// 添加模板文件
	for _, path := range d.templateFileList {
		file, err := templateFS.ReadFile("templates/" + d.templateType + "/" + path)
		if err != nil {
			return err
		}
		if path == "word/styles.xml" {
			file = d.replaceNode(file)
		}

		if err := d.addFileToZip(zipWriter, path, file); err != nil {
			return err
		}
	}

	// 添加 word/_rels/document.xml.rels
	if marshal, err := xml.Marshal(d.docRelation); err != nil {
		return err
	} else {
		if err := d.addFileToZip(zipWriter, "word/_rels/document.xml.rels", marshal); err != nil {
			return err
		}
	}

	// 添加 word/document.xml
	if marshal, err := xml.Marshal(d.Document); err != nil {
		return err
	} else {
		if err := d.addFileToZip(zipWriter, "word/document.xml", marshal); err != nil {
			return err
		}
	}

	// 向 word/media 文件夹中添加图片
	for name, img := range image.ImagesList {
		if err := d.addFileToZip(zipWriter, "word/media/"+name+".jpeg", img.Bytes); err != nil {
			return err
		}
	}

	return nil
}

// AddCustomStyle 添加自定义样式
//
//	styleType 样式类型，可选：character|paragraph|tab|...
func (d *Docx) AddCustomStyle(style *styles.CustomStyle) int64 {
	style.Id = d.styleId
	defer func() { d.styleId += 1 }()
	d.defaultStyle[style.StyleName.Val] = style

	return d.styleId
}

// GetStyle 获取已存在的样式
func (d *Docx) GetStyle(styleName string) *styles.CustomStyle {
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
func (d *Docx) addFileToZip(zipWriter *zip.Writer, zipFilePath string, streamOrPath any) error {

	//创建压缩文件对象
	f, err := zipWriter.Create(zipFilePath)
	if err != nil {
		return err
	}

	// 写入
	switch t := streamOrPath.(type) {
	case string:
		p, _ := os.ReadFile(t)
		if _, err = f.Write(p); err != nil {
			return err
		}
	case []byte:
		if _, err = f.Write(t); err != nil {
			return err
		}
	}
	return nil
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
	xmlnsA   = `http://schemas.openxmlformats.org/drawingml/2006/main`
	xmlnsPic = `http://schemas.openxmlformats.org/drawingml/2006/picture`

	xmlnsRel = `http://schemas.openxmlformats.org/package/2006/relationships`

	relationshipId1 = `http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles`
	relationshipId2 = `http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme`
	relationshipId3 = `http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable`

	imageRelationship = `http://schemas.openxmlformats.org/officeDocument/2006/relationships/image`
)
