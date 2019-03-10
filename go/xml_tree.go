package easyLexML

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type labelConfig struct {
	Sec string
	Cls string
	Sub string
}

type xmlTreePath []*xmlTreeNode

type xmlTreeNode struct {
	Token          xml.Token
	Parent         *xmlTreeNode
	Children       []*xmlTreeNode
	NthChild       int
	NthChildOfElem int
}

func (this *xmlTreeNode) AddChild(token xml.Token) *xmlTreeNode {
	if this.Children == nil {
		this.Children = make([]*xmlTreeNode, 0)
	}
	new_child := new(xmlTreeNode)
	new_child.Parent = this
	new_child.Token = xml.CopyToken(token)
	new_child.NthChild = len(this.Children) + 1
	new_child.NthChildOfElem = 1

	// Find NthChildOfElem
	for i := len(this.Children) - 1; i >= 0; i-- {
		child := this.Children[i]
		if token_same_element(new_child.Token, child.Token) {
			new_child.NthChildOfElem = child.NthChildOfElem + 1
		}
	}

	this.Children = append(this.Children, new_child)

	return new_child
}

func (this *xmlTreeNode) PathToHere() xmlTreePath {
	buf := make([]*xmlTreeNode, 0)
	buf = xmlTreePath(this.pathToHere(buf))
	// Reverse list
	my_len := len(buf)
	ans := make([]*xmlTreeNode, my_len)
	for i, _ := range buf {
		ans[my_len-i-1] = buf[i]
	}
	return xmlTreePath(ans)
}

func (this *xmlTreeNode) pathToHere(buf []*xmlTreeNode) []*xmlTreeNode {
	if this.Token == nil {
		return buf
	}
	buf = append(buf, this)
	if this.Parent == nil {
		return buf
	} else {
		return this.Parent.pathToHere(buf)
	}
}

func (this xmlTreePath) Peek() xml.Token {
	return this[len(this)-1].Token
}

func (this xmlTreePath) PeekTag() string {
	return name2string(this.Peek().(xml.StartElement).Name)
}

func (this xmlTreePath) Has(tag string) bool {
	for _, node := range this {
		switch tk := node.Token.(type) {
		case xml.StartElement:
			if name2string(tk.Name) == tag {
				return true
			}
		}
	}
	return false
}

func (this xmlTreePath) String() string {
	ans := ""
	for i, node := range this {
		if i != 0 {
			ans += " "
		}
		switch tk := node.Token.(type) {
		case xml.StartElement:
			ans += fmt.Sprintf("<%s> (%d|%d)", name2string(tk.Name), node.NthChild, node.NthChildOfElem)
		case xml.CharData:
			ans += fmt.Sprintf("%q", tk)
		default:
			ans += fmt.Sprintf("%+v", tk)
		}
	}
	return ans
}

func (this xmlTreePath) LexId() string {
	ans := ""
	for _, node := range this {
		switch tk := node.Token.(type) {
		case xml.StartElement:
			tag := name2string(tk.Name)
			if tag == "EasyLexML" || tag == "corpus" {
				continue
			}
			if ans != "" {
				ans += "_"
			}
			ans += fmt.Sprintf("%s%d", name2string(tk.Name), node.NthChildOfElem)
		}
	}
	return ans
}

func (this xmlTreePath) Label(cls_counter int, cfg labelConfig) string {
	ans := ""
	last_tag := ""
	title := ""
	for i, node := range this {
		if i < 2 {
			// Ignore the <EasyLexML> and <corpus> part
			continue
		}
		tk, _ := node.Token.(xml.StartElement)
		if ans != "" {
			ans += "."
		}
		tag := name2string(tk.Name)
		last_tag = tag
		ans += strconv.Itoa(node.NthChildOfElem)
		if tag == "cls" {
			ans = ""
			// Is this label for a <cls>?
			if i == len(this)-1 {
				ans = strconv.Itoa(cls_counter)
				break
			}
		}
		if i == len(this)-1 {
			title, _ = token_get_attr(tk, "title")
		}
	}
	if title != "" {
		title = " " + title
	}
	switch last_tag {
	case "cls":
		return strings.Replace(cfg.Cls, "{}", ans, -1) + title
	case "sec":
		return strings.Replace(cfg.Sec, "{}", ans, -1) + title
	case "sub":
		return strings.Replace(cfg.Sub, "{}", ans, -1) + title
	default:
		return ans
	}
}
