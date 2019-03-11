package easyLexML

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type tocTreeNode struct {
	Parent   *tocTreeNode
	Text     string
	Href     string
	Children []*tocTreeNode
}

func (this *tocTreeNode) AddEntry(text, href string) *tocTreeNode {
	if this.Children == nil {
		this.Children = make([]*tocTreeNode, 0)
	}
	ans := new(tocTreeNode)
	ans.Text = text
	ans.Href = href
	ans.Parent = this
	this.Children = append(this.Children, ans)
	return ans
}

func (this *tocTreeNode) String() string {
	ans := fmt.Sprintf("[%s](%s){", this.Text, this.Href)
	for i, child := range this.Children {
		if i != 0 {
			ans += " "
		}
		ans += child.String()
	}
	ans += "}"
	return ans
}

func (this *tocTreeNode) ToXML() string {
	buf := new(bytes.Buffer)
	encoder := xml.NewEncoder(buf)
	encoder.Indent("", "\t")
	this.ToXMLWithEncoder(encoder, "Table of Contents")
	encoder.Flush()
	return buf.String()
}

func (this *tocTreeNode) ToXMLWithEncoder(encoder *xml.Encoder, toc_label string) {
	if this.Parent == nil {
		panicIfErr(encoder.EncodeToken(
			xml.StartElement{
				Name: xml.Name{Local: "toc"},
				Attr: []xml.Attr{xml.Attr{
					Name:  xml.Name{Local: "id"},
					Value: "toc"}}}))
		tk_lbl := xml.StartElement{Name: xml.Name{Local: "label"}}
		panicIfErr(encoder.EncodeToken(tk_lbl))
		panicIfErr(encoder.EncodeToken(xml.CharData(toc_label)))
		panicIfErr(encoder.EncodeToken(tk_lbl.End()))
	}
	if len(this.Children) > 0 {
		panicIfErr(encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: "ul"}}))
	}
	for _, child := range this.Children {
		tk_li := xml.StartElement{Name: xml.Name{Local: "li"}}
		tk_p := xml.StartElement{Name: xml.Name{Local: "p"}}
		tk_a := xml.StartElement{
			Name: xml.Name{Local: "a"},
			Attr: []xml.Attr{xml.Attr{
				Name:  xml.Name{Local: "href"},
				Value: "#" + child.Href,
			}}}
		panicIfErr(encoder.EncodeToken(tk_li))
		panicIfErr(encoder.EncodeToken(tk_p))
		panicIfErr(encoder.EncodeToken(tk_a))
		panicIfErr(encoder.EncodeToken(xml.CharData(child.Text)))
		panicIfErr(encoder.EncodeToken(tk_a.End()))
		panicIfErr(encoder.EncodeToken(tk_p.End()))
		child.ToXMLWithEncoder(encoder, toc_label)
		panicIfErr(encoder.EncodeToken(tk_li.End()))
	}
	if len(this.Children) > 0 {
		panicIfErr(encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "ul"}}))
	}
	if this.Parent == nil {
		panicIfErr(encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "toc"}}))
	}
}
