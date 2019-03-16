package easyLexML

import (
	"strconv"
	"strings"

	"github.com/gjvnq/xmlquery"
	"github.com/jinzhu/copier"
)

type context struct {
	SecLabel    string
	ClsLabel    string
	SubLabel    string
	NoteLabel   string
	SecHeading  string
	ClsHeading  string
	SubHeading  string
	NoteHeading string
}

func (this *context) Update(node *xmlquery.Node) {
	if node.Data != "set-meta" {
		return
	}

	for _, attr := range node.Attr {
		tag := name2string(attr.Name)
		switch tag {
		case "SecLabel":
			this.SecLabel = attr.Value
		case "ClsLabel":
			this.ClsLabel = attr.Value
		case "SubLabel":
			this.SubLabel = attr.Value
		case "NoteLabel":
			this.NoteLabel = attr.Value
		case "SecHeading":
			this.SecHeading = attr.Value
		case "ClsHeading":
			this.ClsHeading = attr.Value
		case "SubHeading":
			this.SubHeading = attr.Value
		case "NoteHeading":
			this.NoteLabel = attr.Value
		}
	}
}

func (this *context) Copy() *context {
	ans := new(context)
	copier.Copy(ans, this)
	return ans
}

func gen_label(node *xmlquery.Node, cls_counter int) {
	// Find teh "real" parent
	parent := node.Parent
	for !tag_has_label(parent.Data) {
		parent = parent.Parent
	}
	title := parent.GetAttrWithDefault("title", "")
	tag := parent.Data
	ctx := parent.Info.(*context)
	ans := ""

	if title == "" {
		switch tag {
		case "sec":
			ans = ctx.SecLabel
		case "cls":
			ans = ctx.ClsLabel
		case "sub":
			ans = ctx.SubLabel
		case "note":
			ans = ctx.NoteLabel
		default:
			ans = "{lexid}"
		}
	} else {
		switch tag {
		case "sec":
			ans = ctx.SecHeading
		case "cls":
			ans = ctx.ClsHeading
		case "sub":
			ans = ctx.SubHeading
		case "note":
			ans = ctx.NoteHeading
		default:
			ans = "{lexid} - {title}"
		}
	}
	ans = parent.GetAttrWithDefault("label-style", ans)

	num := ""
	switch tag {
	case "cls":
		num = strconv.Itoa(cls_counter)
	case "sec":
		tmp := make([]string, 0)
		cursor := parent
		for cursor != nil {
			if cursor.Data == "sec" {
				tmp = append(tmp, strconv.Itoa(parent.NthChildOfElem()+1))
			}
			cursor = cursor.Parent
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			if len(num) > 0 {
				num += "."
			}
			num += tmp[i]
		}
	default:
		num = strconv.Itoa(parent.NthChildOfElem() + 1)
	}
	ans = strings.Replace(ans, "{num}", num, -1)
	ans = strings.Replace(ans, "{id}", parent.GetAttrWithDefault("id", ""), -1)
	ans = strings.Replace(ans, "{lexid}", parent.GetAttrWithDefault("lexid", ""), -1)
	ans = strings.Replace(ans, "{title}", title, -1)

	href := parent.GetAttrWithDefault("id", "")
	parts := strings.Split(ans, `\n`)
	for i, part := range parts {
		if i != 0 {
			node.AddChild(new_node_element("br"))
		}
		node.AddChild(new_node_text(part))
	}

	if href != "" {
		node.SetAttr("href", "#"+href)
	}
}

func tag_has_label(tag string) bool {
	return tag == "sec" || tag == "sec-nn" || tag == "cls" || tag == "sub" || tag == "note"
}

func update_cls_counter(node *xmlquery.Node, cls_counter *int) {
	var err error

	if node.Data == "set-meta" {
		val, ok := node.GetAttr("ClsCounter")
		if ok {
			*cls_counter, err = strconv.Atoi(val)
			panicIfErr(err)
		}
	}
	if node.Data == "cls" {
		*cls_counter++
	}
}

func tag_has_lexid(tag string) bool {
	return tag == "sec" || tag == "sec-nn" || tag == "cls" || tag == "sub" || tag == "note" || tag == "p"
}
