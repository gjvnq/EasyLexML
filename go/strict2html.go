package easyLexML

import (
	"bytes"
	"html/template"
	"io"
	"strings"

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
	toc := base.SelectElement("toc")

	// Replace elements
	replace_with_html_elements(base)

	// Get title
	metadata := base.SelectElement("metadata")
	if metadata != nil {
		title := metadata.SelectElement("title")
		pageData.Title = title.OutputXML(false)
	}

	// Finish Corpus
	buf := new(bytes.Buffer)
	corpus.OutputXMLToWriter(buf, true, true)
	pageData.Corpus = template.HTML(buf.String())

	// Finish Metadata
	buf = new(bytes.Buffer)
	if toc != nil {
		toc.OutputXMLToWriter(buf, true, true)
		pageData.Toc = template.HTML(buf.String())
	}

	// Finish
	tmpl.Execute(output, pageData)
	return nil
}

func replace_with_html_elements(root *xmlquery.Node) {
	if root.Type != xmlquery.ElementNode {
		return
	}

	tag := root.Data
	old_tag := tag
	switch {
	case tag == "toc":
		tag = "section"
	case tag == "corpus":
		tag = "section"
	case tag == "note":
		tag = "section"
		root.SetAttr("class", "note")
	case tag == "sec" || tag == "sec-nn":
		tag = "section"
	case tag == "cls" || tag == "cls-nn":
		tag = "section"
	case tag == "sub" || tag == "sub-nn":
		tag = "section"
		root.SetAttr("class", "sub")
	case tag == "label":
		tag = "a"
		root.SetAttr("class", "label")
	}
	root.Data = tag
	// Ensure we don't lose the information of the original XML tag
	if old_tag != tag {
		root.SetAttr("data-tag", old_tag)
	}
	// Fix attributes
	prefix_non_html_attributes(root)
	// Recursive step
	for child := root.FirstChild; child != nil; child = child.NextSibling {
		replace_with_html_elements(child)
	}
}

func is_math_ml_element(tag string) bool {
	ans := false
	ans = ans || (tag == "math")
	ans = ans || (tag == "maction")
	ans = ans || (tag == "maligngroup")
	ans = ans || (tag == "malignmark")
	ans = ans || (tag == "menclose")
	ans = ans || (tag == "merror")
	ans = ans || (tag == "mfenced")
	ans = ans || (tag == "mfrac")
	ans = ans || (tag == "mglyph")
	ans = ans || (tag == "mi")
	ans = ans || (tag == "mlabeldtr")
	ans = ans || (tag == "mlongdiv")
	ans = ans || (tag == "mmultiscript")
	ans = ans || (tag == "mn")
	ans = ans || (tag == "mo")
	ans = ans || (tag == "mover")
	ans = ans || (tag == "mpadded")
	ans = ans || (tag == "mphantom")
	ans = ans || (tag == "mroot")
	ans = ans || (tag == "mrow")
	ans = ans || (tag == "ms")
	ans = ans || (tag == "mscarries")
	ans = ans || (tag == "mscarry")
	ans = ans || (tag == "msgroup")
	ans = ans || (tag == "mstack")
	ans = ans || (tag == "msline")
	ans = ans || (tag == "mspace")
	ans = ans || (tag == "msqrt")
	ans = ans || (tag == "msrow")
	ans = ans || (tag == "mstyle")
	ans = ans || (tag == "msub")
	ans = ans || (tag == "msup")
	ans = ans || (tag == "msubsup")
	ans = ans || (tag == "mtable")
	ans = ans || (tag == "mtd")
	ans = ans || (tag == "mtext")
	ans = ans || (tag == "mtr")
	ans = ans || (tag == "munder")
	ans = ans || (tag == "munderover")
	ans = ans || (tag == "semantics")
	ans = ans || (tag == "annotation")
	ans = ans || (tag == "annotation-xml")

	return ans
}

func is_valid_html_attribute(tag, attr string) bool {
	ans := false

	if strings.HasPrefix(attr, "data-") {
		return true
	}

	//General attributes
	ans = ans || (attr == "id")
	ans = ans || (attr == "class")
	ans = ans || (attr == "style")
	ans = ans || (attr == "display")
	ans = ans || (attr == "height")
	ans = ans || (attr == "width")

	if is_math_ml_element(tag) {
		return true
	}

	// Tag specific attributes
	switch tag {
	case "a":
		ans = ans || (attr == "href")
	}

	return ans
}

func prefix_non_html_attributes(node *xmlquery.Node) {
	if node.Type != xmlquery.ElementNode {
		return
	}
	for i := range node.Attr {
		attr := &node.Attr[i]
		if !is_valid_html_attribute(node.Data, attr.Name.Local) {
			attr.Name.Local = "data-" + attr.Name.Local
		}
	}
}
