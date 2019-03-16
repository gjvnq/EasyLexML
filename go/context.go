package easyLexML

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gjvnq/xmlquery"
	"github.com/jinzhu/copier"
)

type context struct {
	SecLabel   string
	ClsLabel   string
	SubLabel   string
	ClsCounter int
}

func (this *context) Update(node *xmlquery.Node) error {
	var err error

	if node.Data == "cls" {
		this.ClsCounter++
		return nil
	}

	if node.Data != "set-meta" {
		return nil
	}

	for _, attr := range node.Attr {
		tag := name2string(attr.Name)
		switch tag {
		case "ClsCounter":
			this.ClsCounter, err = strconv.Atoi(attr.Value)
		case "SecLabel":
			this.SecLabel = attr.Value
		case "ClsLabel":
			this.ClsLabel = attr.Value
		case "SubLabel":
			this.SubLabel = attr.Value
		}
		panicIfErr(err)
	}

	return err
}

func (this *context) String() string {
	return fmt.Sprintf("context{%d %q %q %q}",
		this.ClsCounter,
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

func (this *context) GetLabel(node *xmlquery.Node) string {
	// Find teh "real" parent
	parent := node.Parent
	for !tag_has_label(parent.Data) {
		parent = parent.Parent
	}

	tag := parent.Data
	ans := ""
	switch tag {
	case "sec":
		ans = this.SecLabel
	case "cls":
		ans = strings.Replace(this.ClsLabel, "{num}", strconv.Itoa(this.ClsCounter), 0)
	case "sub":
		ans = this.SubLabel
	}
	ans = strings.Replace(ans, "{num}", strconv.Itoa(parent.NthChildOfElem()), 0)

	return ans
}

func tag_has_label(tag string) bool {
	return tag == "sec" || tag == "cls" || tag == "sub"
}
