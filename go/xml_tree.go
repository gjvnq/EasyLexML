package easyLexML

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type LabelConfig struct {
	Sec string
	Cls string
	Sub string
	Toc string
}

type xmlTreePath []*XMLTreeNode

type XMLTreeNode struct {
	Token          xml.Token
	Parent         *XMLTreeNode
	Children       []*XMLTreeNode
	LabelConfig 	LabelConfig
}

func (this *XMLTreeNode) Encode(out io.Writer) error {
	return this.encode(out, 0)
}

func (this *XMLTreeNode) encode(out io.Writer, depth int) error {
	var err error
	buf := new(bytes.Buffer)
	has_to_end_tag := false
	end_tag := new(bytes.Buffer)

	// Ensure proper indentation
	for i:=0; i < depth; i++ {
		buf.WriteString("\t")
		end_tag.WriteString("\t")
	}

	// Encode this token (unless we are root)
	if this.Parent != nil {
		switch tk := this.Token.(type) {
			case xml.StartElement:
				buf.WriteString("<")
				buf.WriteString(name2string(tk.Name))
				for _, attr := range tk.Attr {
					buf.WriteString(" ")
					buf.WriteString(name2string(attr.Name))
					buf.WriteString(`="`)
					xml.EscapeText(buf, attr.Value)
					buf.WriteString(`"`)
				}
				// Make self closing tag if possible
				if len(this.Children) == 0 {
					buf_str += "/>"
					buf.WriteString(`/>`)
				} else {
					buf.WriteString(`>`)
					has_to_end_tag = true
					end_tag.WriteString(`<`)
					end_tag.WriteString(name2string(tk.Name))
					end_tag.WriteString(`/>`)
				}
			case xml.CharData:
				err = xml.EscapeText(buf, []byte(tk))
				if err != nil {
					return err
				}
			case xml.Comment:
				if bytes.Contains(tk, `-->`) {
					return errors.New("xml: EncodeToken of Comment containing --> marker")
				}
				buf.WriteString(`<!--`)
				buf.Write([]byte(tk))
				buf.WriteString(`-->`)
			case xml.ProcInst:
				buf.WriteString(`<!`)
				buf.Write([]byte(tk.Target))
				if len(tk.Inst) {
					buf.WriteString(` `)
					buf.Write([]byte(tk.Inst))
				}
				buf.WriteString(`?>`)
			case xml.Directive:
				buf.WriteString(`<!`)
				buf.Write([]byte(tk))
				buf.WriteString(`>`)
			default:
				return errors.New("xml: EncodeToken of invalid token type")
			}
		}

		// Output
		buf.WriteString("\n")
		err = out.Write(buf)
		if err != nil {
			return err
		}
	}

	// Encode children
	depth++
	for _, child := this.Children {
		err = child.encode(out, depth)
		if err != nil {
			return err
		}
	}
	depth--

	// Finalize tag if needed (also ensure we are not root)
	end_tag.WriteString("\n")
	if has_to_end_tag && this.Parent != nil {
		err = out.Write([]byte(end_tag))
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *XMLTreeNode) NthChild() int {
	for i, child := range this.Parent.Children {
		if child == this {
			return i
		}
	}
	panic("token not contained in parent")
}

func (this *XMLTreeNode) NthChildOfElem() int {
	ans := 0
	for _, child := range this.Parent.Children {
		if token_same_element(this, child) {
			ans++
		}
	}
	return ans
}

func (this *XMLTreeNode) new_child(token xml.Token) *XMLTreeNode {
	if this.Children == nil {
		this.Children = make([]*XMLTreeNode, 0)
	}
	new_child := new(XMLTreeNode)
	new_child.Parent = this
	new_child.Token = xml.CopyToken(token)
}

func (this *XMLTreeNode) Tag(token xml.Token) *XMLTreeNode {
	switch tk := this.Token.(type) {
	case xml.StartElement:
		return name2string(tk.Name)
	default:
		return ""
	}
}

func (this *XMLTreeNode) AddChild(token xml.Token) *XMLTreeNode {
	new_child := this.new_child(token)
	this.Children = append(this.Children, new_child)
	return new_child
}

func (this *XMLTreeNode) Replace(token xml.Token) {
	this.Token = token
}

func (this *XMLTreeNode) insert_at(index int, token xml.Token) *XMLTreeNode {
	new_child := this.new_child(token)
	this.Children = this.Children[0:len(this.Children)+1]
	copy(slice[index+1:], slice[index:])
	this.Children[index] = new_child
	return new_child
}

func (this *XMLTreeNode) InsertBefore(token xml.Token) *XMLTreeNode {
	return this.Parent.insert_at(this.NthChild(), token)
}

func (this *XMLTreeNode) InsertAfter(token xml.Token) *XMLTreeNode {
	return this.Parent.insert_at(this.NthChild()+1, token)
}

func (this *XMLTreeNode) PathToHere() xmlTreePath {
	buf := make([]*XMLTreeNode, 0)
	buf = xmlTreePath(this.pathToHere(buf))
	// Reverse list
	my_len := len(buf)
	ans := make([]*XMLTreeNode, my_len)
	for i, _ := range buf {
		ans[my_len-i-1] = buf[i]
	}
	return xmlTreePath(ans)
}

func (this *XMLTreeNode) pathToHere(buf []*XMLTreeNode) []*XMLTreeNode {
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

func (this xmlTreePath) PeekN(n int) xml.Token {
	return this[len(this)-n].Token
}

func (this xmlTreePath) PeekTagN(n int) string {
	return name2string(this.PeekN(n).(xml.StartElement).Name)
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

func (this xmlTreePath) Label(cls_counter int, cfg LabelConfig) string {
	ans := ""
	last_tag := ""
	title := ""
	var last_token xml.StartElement
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
		last_token = tk
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

	if style, ok := token_get_attr(last_token, "label-style"); ok {
		return strings.Replace(style, "{}", ans, -1) + title
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

func StreamXML2Tree(input io.Reader, trim bool) (error, *XMLTreeNode) {
	// Prepare
	root := new(XMLTreeNode)
	cursor := root
	decoder := xml.NewDecoder(input)

	// Parse
	for {
		token, err := decoder.RawToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err, root
		}
		if trim {
			token = sanitize_char_data(token)
		}

		switch tk := token.(type) {
		case xml.StartElement:
			cursor = cursor.AddChild(tk)
		case xml.EndElement:
			cursor = cursor.Parent
		default:
			cursor.AddChild(tk)
		}
	}
	return root
}