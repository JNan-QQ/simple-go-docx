package image

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"gitee.com/jn-qq/go-tools/data"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
)

// Image 图片信息
//
//	Image.Local: 本地图片路径， Image.Online: 在线图片链接， Image.Bytes: 图片文件二进制数据切片。三者任选其一
//	Name: 图片名称，文档中唯一，便于复用图片
type Image struct {
	Local  string
	Online string
	Bytes  []byte
	rId    string
	Name   string
	size   []int
}

// ImagesList 图片数据缓存
var ImagesList = map[string]*Image{} // 图片集合

// Down 获取图片二进制数据。 compress 是否对图片进行压缩已达到减少文件体积的目的。
func (i *Image) Down(quality int) error {
	if i.Local != "" {
		if _, err := os.Stat(i.Local); err != nil {
			return fmt.Errorf("本地文件不存在！")
		}
		// 读取本地文件
		if file, err := os.ReadFile(i.Local); err != nil {
			return err
		} else {
			i.Bytes = file
		}

	} else if i.Online != "" {
		// 获取网络图片
		response, err := http.Get(i.Online)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		// 解析图片内容
		if i.Bytes, err = io.ReadAll(response.Body); err != nil {
			return err
		}

	} else if i.Bytes == nil {
		return fmt.Errorf("文件对象为空！")
	}

	// 通过前512字节判断图片类型
	n := 512
	if len(i.Bytes) < 512 {
		n = len(i.Bytes)
	}
	if !data.Contains([]string{"image/jpeg", "image/png"}, http.DetectContentType(i.Bytes[:n])) {
		return fmt.Errorf("未知图片类型")
	}

	// 获取原始图片宽高
	im, _, _ := image.DecodeConfig(bytes.NewReader(i.Bytes))
	i.size = []int{im.Width, im.Height}

	// 压缩图片质量
	if quality != 100 {
		buf := bytes.Buffer{}
		img, _, _ := image.Decode(bytes.NewReader(i.Bytes))
		err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 40})
		if err == nil {
			i.Bytes = buf.Bytes()
		}
	}

	return nil
}

// SetId 设置映射关系id
func (i *Image) SetId(id string) {
	i.rId = id
}

// GetId 获取映射关系id
func (i *Image) GetId() string {
	return i.rId
}

// Draw 画布
type Draw struct {
	XMLName xml.Name `xml:"w:drawing"`
	Inline  *Inline  // 内嵌图片
	Anchor  *Anchor  // 环绕图片
}

// Inline 默认内嵌类型图片布局
type Inline struct {
	XMLName xml.Name `xml:"wp:inline"`
	Extent  *Extent
	Pic     *Pic `xml:"a:graphic>a:graphicData>pic:pic"`
}

// Anchor 环绕类型图片布局 未实现
type Anchor struct {
	XMLName xml.Name `xml:"wp:anchor"`
}

// Extent 画布大小 1twip = 635emus
type Extent struct {
	XMLName xml.Name `xml:"wp:extent"`
	Cx      int64    `xml:"cx,attr"`
	Cy      int64    `xml:"cy,attr"`
}

// Pic 图片对象
type Pic struct {
	XMLName xml.Name `xml:"pic:pic"`
	ImgId   *ImgId   `xml:"pic:blipFill>a:blip"`
	ImgSize *ImgSize `xml:"pic:spPr>a:xfrm>a:ext"`
}

// ImgId 图片映射id
type ImgId struct {
	XMLName xml.Name `xml:"a:blip"`
	RId     string   `xml:"r:embed,attr"`
}

// ImgSize 图片大小 1twip = 635emus
type ImgSize struct {
	XMLName xml.Name `xml:"a:ext"`
	Cx      int64    `xml:"cx,attr"`
	Cy      int64    `xml:"cy,attr"`
}
