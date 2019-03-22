package easyLexML

import (
	"bytes"
	"html/template"
	"io"

	"github.com/gjvnq/xmlquery"
)

type htmlPage struct {
	Title         string
	Corpus        template.HTML
	Abstract      template.HTML
	AbstractTitle string
	Toc           template.HTML
}

func Strict2HTML(input io.Reader, output io.Writer) error {
	// Prepare the template
	pageData := htmlPage{}
	tmpl_raw, err := Asset("res/standalone.html")
	panicIfErr(err)
	tmpl := template.New("standalone")
	tmpl, err = tmpl.Parse(string(tmpl_raw))
	panicIfErr(err)

	// Read XML
	root, err := xmlquery.Parse(input)
	if err != nil {
		return err
	}
	base := root.SelectElement("EasyLexML")
	corpus := base.SelectElement("corpus")

	// Replace elements
	replace_with_html_elements(base)

	// Get title
	metadata := base.SelectElement("metadata")
	if metadata != nil {
		title := metadata.SelectElement("title")
		pageData.Title = title.OutputXML(false)
	}

	// Finish
	buf := new(bytes.Buffer)
	corpus.OutputXMLToWriter(buf, true, true)
	pageData.Corpus = template.HTML(buf.String())
	tmpl.Execute(output, pageData)

	return nil
}

func replace_with_html_elements(root *xmlquery.Node) {
	if root.Type != xmlquery.ElementNode {
		return
	}

	tag := root.Data
	switch {
	case tag == "sec" || tag == "sec-nn":
		tag = "section"
	case tag == "cls" || tag == "cls-nn":
		tag = "section"
	case tag == "sub" || tag == "sub-nn":
		tag = "section"
	case tag == "label":
		tag = "a"
		root.SetAttr("class", "label")
	}
	root.SetAttr("data-tag", tag)
	root.Data = tag
	for child := root.FirstChild; child != nil; child = child.NextSibling {
		replace_with_html_elements(child)
	}
}
