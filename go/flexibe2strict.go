package easyLexML

import (
	"encoding/xml"
	"io"
	"strings"

	"github.com/gjvnq/xmlquery"
)

func Draft2Strict(input io.Reader, output io.Writer) error {
	// Read XML
	root, err := xmlquery.Parse(input)
	base := root.SelectElement("EasyLexML")
	// cursor := root
	if err != nil {
		return err
	}

	// Remove TOC
	node := base.SelectElement("toc")
	node.DeleteMe()

	// Put text within <p> (and also add <label>)
	cursor := base.SelectElement("corpus")
	envelop_text(cursor)

	// Add id and lexid

	// Remove all <set-meta>
	node = xmlquery.FindOne(root, "//set-meta")
	for node != nil {
		node.DeleteMe()
		node = xmlquery.FindOne(root, "//set-meta")
	}

	// Output
	root.OutputXMLToWriter(output, true, false)

	return nil
}

func add_label(encoder *xml.Encoder, next_label string) {
	tk_lbl := xml.StartElement{Name: xml.Name{Local: "label"}}
	Debugln(">>>", token2string(tk_lbl))
	panicIfErr(encoder.EncodeToken(tk_lbl))

	tk_txt := xml.CharData(next_label)
	Debugln(">>>", token2string(tk_txt))
	panicIfErr(encoder.EncodeToken(tk_txt))

	Debugln(">>>", token2string(tk_lbl.End()))
	panicIfErr(encoder.EncodeToken(tk_lbl.End()))

	tk_space := xml.CharData(" ")
	Debugln(">>>", token2string(tk_space))
	panicIfErr(encoder.EncodeToken(tk_space))
}

func envelop_text(root *xmlquery.Node) {
	if !requires_p(root) {
		for child := root.FirstChild; child != nil; child = child.NextSibling {
			envelop_text(child)
		}
		return
	}

	// Get children
	children := make([]*xmlquery.Node, 0)
	for child := root.FirstChild; child != nil; child = child.NextSibling {
		children = append(children, child)
	}
	// Remove references
	root.FirstChild = nil
	root.LastChild = nil
	for _, child := range children {
		child.Parent = nil
		child.NextSibling = nil
		child.PrevSibling = nil
		// Also remove double whitespace
		if child.Type == xmlquery.TextNode {
			if len(strings.TrimSpace(child.Data)) == 0 {
				child.Data = remove_double_whitespace(child.Data)
			}
		}
	}

	// Readd children
	var p_node *xmlquery.Node
	state := 0 // 0 - Looking for start node | 1 - Looking for end node
	for _, child := range children {
		// Ignore empty children
		if len(child.Data) == 0 {
			continue
		}

		switch state {
		case 0:
			if child.Type == xmlquery.TextNode {
				p_node = new_node_element("p")
				lbl_node := new_node_element("label")
				root.AddChild(p_node)
				p_node.AddChild(lbl_node)
				p_node.AddChild(child)
				state = 1
			} else {
				root.AddChild(child)
			}
		case 1:
			if !requires_p(child) || child.Type == xmlquery.TextNode {
				p_node.AddChild(child)
			} else {
				state = 0
				p_node = nil
				root.AddChild(child)
			}
		}
	}

	// Recursion
	for child := root.FirstChild; child != nil; child = child.NextSibling {
		envelop_text(child)
	}
}

func requires_p(node *xmlquery.Node) bool {
	if node.Type == xmlquery.ElementNode {
		tag := node.Data
		return tag == "sec" || tag == "cls" || tag == "sub"
	}
	return node.Type == xmlquery.TextNode
}

func new_node_element(name string) *xmlquery.Node {
	ans := new(xmlquery.Node)
	ans.Type = xmlquery.ElementNode
	ans.Data = name

	return ans
}
