package easyLexML

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gjvnq/xmlquery"
	"github.com/jinzhu/copier"
)

type context struct {
	SecLabel  string
	ClsLabel  string
	SubLabel  string
	NoteLabel string
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
		}
	}
}

func (this *context) String() string {
	return fmt.Sprintf("context{%q %q %q}",
		this.SecLabel,
		this.ClsLabel,
		this.SubLabel,
	)
}

func (this *context) Copy() *context {
	ans := new(context)
	copier.Copy(ans, this)
	return ans
}

func gen_label(node *xmlquery.Node, cls_counter int) (string, string) {
	// Find teh "real" parent
	parent := node.Parent
	for !tag_has_label(parent.Data) {
		parent = parent.Parent
	}

	tag := parent.Data
	ctx := parent.Info.(*context)
	ans := ""
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
	ans = parent.GetAttrWithDefault("label-style", ans)

	num := strconv.Itoa(cls_counter)
	if tag != "cls" {
		num = strconv.Itoa(parent.NthChildOfElem() + 1)
	}
	ans = strings.Replace(ans, "{num}", num, -1)
	ans = strings.Replace(ans, "{id}", parent.GetAttrWithDefault("id", ""), -1)
	ans = strings.Replace(ans, "{lexid}", parent.GetAttrWithDefault("lexid", ""), -1)

	return ans, parent.GetAttrWithDefault("id", "")
}

func tag_has_label(tag string) bool {
	return tag == "sec" || tag == "cls" || tag == "sub" || tag == "note"
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
	return tag == "sec" || tag == "cls" || tag == "sub" || tag == "note" || tag == "p"
}
