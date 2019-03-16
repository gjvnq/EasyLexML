package easyLexML

import (
	"strconv"

	"github.com/gjvnq/xmlquery"
)

type variables struct {
	TocTitle   string
	SecLabel   string
	ClsLabel   string
	SubLabel   string
	ClsCounter int
}

func (this *variables) Update(node *xmlquery.Node) error {
	var err error

	for _, attr := range node.Attr {
		tag := name2string(attr.Name)
		switch tag {
		case "ClsCounter":
			this.ClsCounter, err = strconv.Atoi(attr.Value)
		case "TocTitle":
			this.TocTitle = attr.Value
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
