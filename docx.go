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

	customStyle []*styles.CustomStyle
}

// NewDocx 生成一个新的空 A4 docx 文件，我们可以对其进行操作和保存
func NewDocx() *Docx {
	docx := &Docx{
		Document: document.Document{
			XMLName: xml.Name{Space: "w"},
			XmlW:    document.XMLNS_W,
			XmlR:    document.XMLNS_R,
			XmlWp:   document.XMLNS_WP,
			XmlWps:  document.XMLNS_WPS,
			XmlWpc:  document.XMLNS_WPC,
			XmlWpg:  document.XMLNS_WPG,
			Body:    document.Body{Items: make([]interface{}, 0, 64)},
		},
		docRelation: document.Relationships{
			Xmlns: document.XMLNS_REL,
			Relationship: []document.Relationship{
				{
					ID:     "rId1",
					Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles`,
					Target: "styles.xml",
				},
				{
					ID:     "rId2",
					Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme`,
					Target: "theme/theme1.xml",
				},
				{
					ID:     "rId3",
					Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable`,
					Target: "fontTable.xml",
				},
			},
		},
		styleId:      10,
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
	}

	return docx
}

// AddParagraph 添加段落
func (d *Docx) AddParagraph() *paragraph.Paragraph {
	p := &paragraph.Paragraph{
		Items:      make([]interface{}, 0, 64),
		Properties: &paragraph.Properties{},
	}
	d.Document.Body.Items = append(d.Document.Body.Items, p)

	// 获取当前文件路径
	_, path, _, _ := runtime.Caller(0)
	d.workDir, _ = filepath.Split(path)

	return p
}

func (d *Docx) Save(savePath string) error {
	if _, err := os.Stat(savePath); err == nil {
		// 删除旧文件
		if err := os.Remove(savePath); err != nil {
			return err
		}
	}
	// 创建文件
	docx, _ := os.Create(savePath)
	// 创建zip写入对象
	zipWriter := zip.NewWriter(docx)
	defer zipWriter.Close()

	// 添加模板文件
	for _, path := range d.templateFileList {
		file, err := os.ReadFile(filepath.Join(d.workDir, "templates", d.templateType, path))
		if err != nil {
			return err
		}
		if path == "word/styles.xml" && len(d.customStyle) != 0 {
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

// NewCustomStyle 自定义样式
//
//	styleName 样式名称
//	styleType 样式类型，可选：character|paragraph|tab|...
func (d *Docx) NewCustomStyle(styleName string, styleType string) *styles.CustomStyle {
	customStyle := &styles.CustomStyle{
		Type:      styleType,
		Id:        d.createStyleId(),
		Flg:       "1",
		StyleName: &styles.Name{Val: styleName},
		Format:    &styles.Format{},
	}

	d.customStyle = append(d.customStyle, customStyle)

	return customStyle
}

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

func (d *Docx) replaceNode(zipFile []byte) []byte {

	var customStyle []byte
	// 读取自定义样式
	for _, cs := range d.customStyle {
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
