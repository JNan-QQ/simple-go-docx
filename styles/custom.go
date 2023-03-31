package styles

import (
	"encoding/xml"
)

type CustomStyle struct {
	XMLName   xml.Name `xml:"w:style"`
	Type      string   `xml:"w:type,attr"`
	Id        int64    `xml:"w:styleId,attr"`
	Flg       string   `xml:"w:customStyle,attr"`
	StyleName *Name
	Format    *Format
	PPR       *ParagraphProperties
	RPR       *RunProperties
}

func (c *CustomStyle) CreatePPR() *ParagraphProperties {
	c.PPR = &ParagraphProperties{}
	return c.PPR
}

func (c *CustomStyle) CreateRPR() *RunProperties {
	c.RPR = &RunProperties{}
	return c.RPR
}
