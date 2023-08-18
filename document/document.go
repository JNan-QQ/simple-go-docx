package document

import (
	"encoding/xml"
)

// Document docx对象
//
//	Xmlns 命名空间，Body 文档主体
type Document struct {
	XMLName xml.Name `xml:"w:document"`
	XmlW    string   `xml:"xmlns:w,attr"`             // cannot be unmarshalled in
	XmlR    string   `xml:"xmlns:r,attr,omitempty"`   // cannot be unmarshalled in
	XmlWp   string   `xml:"xmlns:wp,attr,omitempty"`  // cannot be unmarshalled in
	XmlWps  string   `xml:"xmlns:wps,attr,omitempty"` // cannot be unmarshalled in
	XmlWpc  string   `xml:"xmlns:wpc,attr,omitempty"` // cannot be unmarshalled in
	XmlWpg  string   `xml:"xmlns:wpg,attr,omitempty"` // cannot be unmarshalled in
	XmlPic  string   `xml:"xmlns:pic,attr,omitempty"` // cannot be unmarshalled in
	XmlA    string   `xml:"xmlns:a,attr,omitempty"`   // cannot be unmarshalled in

	Body Body `xml:"w:body"`
}

// Body <w:body>
// Body.Items 文本内容格式
type Body struct {
	Items []interface{}
}

// Relationships ...
type Relationships struct {
	Xmlns        string `xml:"xmlns,attr"`
	Relationship []Relationship
}

// Relationship ...
type Relationship struct {
	ID         string `xml:"Id,attr"`
	Type       string `xml:"Type,attr"`
	Target     string `xml:"Target,attr"`
	TargetMode string `xml:"TargetMode,attr,omitempty"`
}
