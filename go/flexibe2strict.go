package easyLexML

import (
	"io"
	"strconv"
	"strings"

	"github.com/gjvnq/xmlquery"
)

func Draft2Strict(input io.Reader, output io.Writer) error {
	// Read XML
	root, err := xmlquery.Parse(input)
	base := root.SelectElement("EasyLexML")
	corpus := base.SelectElement("corpus")
	tocTitle := "Table of Contents"
	abstractTitle := "Abstract"
	// cursor := root
	if err != nil {
		return err
	}

	// Remove TOC
	node := base.SelectElement("toc")
	if node != nil {
		node.DeleteMe()
	}

	// Put text within <p> (and also add <label>)
	envelop_text(corpus)

	// Add id, lexid and labels
	ctx := new(context)
	ctx.SecLabel = "§ {num}"
	ctx.ClsLabel = "Cls. {num}"
	ctx.SubLabel = "{num})"
	ctx.NoteLabel = "Note {num} —"
	ctx.SecHeading = "§ {num}\\n{title}"
	ctx.ClsHeading = "Cls. {num}\\n{title}"
	ctx.SubHeading = "{num}\\n{title}"
	ctx.NoteHeading = "Note {num}\\n{title}"
	cls_counter := 0
	corpus.Info = ctx
	process_ids_and_labels(corpus, &cls_counter)
	corpus.SetAttr("id", "corpus")

	// Get TocTitle
	node = xmlquery.FindOne(root, "//set-meta[@TocTitle]")
	if node != nil {
		tocTitle = node.GetAttrWithDefault("TocTitle", tocTitle)
	}

	// Get AbstractTitle
	node = xmlquery.FindOne(root, "//set-meta[@AbstractTitle]")
	if node != nil {
		abstractTitle = node.GetAttrWithDefault("AbstractTitle", abstractTitle)
	}
	// Add <label> to <abstract>
	node = xmlquery.FindOne(root, "//abstract")
	if node != nil {
		node.SetAttr("title", abstractTitle)
		envelop_text(node)
	}
	node = xmlquery.FindOne(root, "//abstract/label")
	if node != nil {
		node.AddChild(new_node_text(abstractTitle))
	}

	// Remove all <set-meta>
	node = xmlquery.FindOne(root, "//set-meta")
	for node != nil {
		node.DeleteMe()
		node = xmlquery.FindOne(root, "//set-meta")
	}

	// Remove "unnecessary" attributes
	remove_draft_attr(base)

	generate_toc(base, tocTitle)

	// Output
	root.OutputXMLToWriter(output, true, false)

	return nil
}

func generate_toc(base *xmlquery.Node, toc_title string) {
	// Preapre
	toc_node := new_node_element("toc")
	toc_label := new_node_element("label")
	toc_label.AddChild(new_node_text(toc_title))
	toc_node.AddChild(toc_label)
	toc_ul := new_node_element("ul")
	toc_node.AddChild(toc_ul)

	// Add TOC after <abstract>
	abstract_node := base.SelectElement("abstract")
	if abstract_node != nil {
		abstract_node.AddAfter(toc_node)
	} else {
		// Add toc before <metadata>
		metadata_node := base.SelectElement("metadata")
		if metadata_node != nil {
			metadata_node.AddAfter(toc_node)
		} else {
			base.FirstChild.AddBefore(toc_node)
		}
	}

	toc_iterator_generator(toc_ul, base)
}

func toc_iterator_generator(toc_cursor, doc_cursor *xmlquery.Node) {
	// Only <sec> and <cls> get TOC entries
	if doc_cursor.Type != xmlquery.ElementNode {
		return
	}
	if doc_cursor.Data == "sec" || doc_cursor.Data == "cls" {
		// Get name
		label := doc_cursor.SelectElement("label")
		if label != nil {
			// Generate and add link
			li := new_node_element("li")
			link := new_node_element("a")
			link.SetAttr("href", "#"+doc_cursor.SelectAttr("id"))
			genTocEntry(label, link)
			li.AddChild(link)
			toc_cursor.AddChild(li)

			// Create a sub level
			if doc_cursor.Data == "sec" {
				ul := new_node_element("ul")
				li.AddChild(ul)
				toc_cursor = ul
			}
		}
	}

	// Recursive step
	for child := doc_cursor.FirstChild; child != nil; child = child.NextSibling {
		toc_iterator_generator(toc_cursor, child)
	}
}

func genTocEntry(label, toc *xmlquery.Node) {
	i := 0
	for label_child := label.FirstChild; label_child != nil; label_child = label_child.NextSibling {
		if label_child.Type == xmlquery.ElementNode && label_child.Data == "span" {
			if i > 0 {
				toc.AddChild(new_node_text(" - "))
			}
			genTocEntry(label_child, toc)
			i++
		} else {
			toc_child := new(xmlquery.Node)
			toc_child.Type = label_child.Type
			toc_child.Data = label_child.Data
			genTocEntry(label_child, toc_child)
			toc.AddChild(toc_child)
		}
	}
}

func process_ids_and_labels(node *xmlquery.Node, cls_counter *int) {
	if node.Type != xmlquery.ElementNode {
		return
	}

	ctx := node.Info.(*context)
	// Set lexid and id if they areno not set yet
	lexid, _ := node.GetAttr("lexid")
	id, _ := node.GetAttr("id")
	if tag_has_lexid(node.Data) && lexid == "" {
		node.SetAttr("lexid", gen_lexid(node))
	}
	if tag_has_lexid(node.Data) && id == "" {
		lexid, _ = node.GetAttr("lexid")
		node.SetAttr("id", lexid+"_v1")
	}

	// Update variables
	ctx.Update(node)
	update_cls_counter(node, cls_counter)

	// Fill labels
	if node.Data == "label" {
		gen_label(node, *cls_counter)
	}

	// Take care of the children
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ctx.Update(child)
		child.Info = ctx.Copy()
		process_ids_and_labels(child, cls_counter)
	}
}

func gen_lexid(node *xmlquery.Node) string {
	if node.Type != xmlquery.ElementNode {
		return ""
	}

	list := make([]string, 0)
	cursor := node
	for cursor != nil && cursor.Data != "corpus" {
		lexid_part := cursor.Data + strconv.Itoa(cursor.NthChildOfElem()+1)
		list = append(list, lexid_part)
		cursor = cursor.Parent
	}

	ans := ""
	for i := len(list) - 1; i >= 0; i-- {
		if len(ans) != 0 {
			ans += "_"
		}
		ans += list[i]
	}

	return ans
}

func remove_draft_attr(root *xmlquery.Node) {
	if root.Type != xmlquery.ElementNode {
		return
	}

	root.DelAttr("label-style")
	root.DelAttr("ref")

	for child := root.FirstChild; child != nil; child = child.NextSibling {
		remove_draft_attr(child)
	}
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

	// Re-add children
	var p_node *xmlquery.Node
	has_label := false
	state := 0 // 0 - Looking for start node | 1 - Looking for end node

	// Special case: labels that are headings (sections and EU-style articles)
	_, ok_title := root.GetAttr("title")
	if root.Data == "sec" || ok_title {
		lbl_node := new_node_element("label")
		root.AddChild(lbl_node)
		has_label = true
	}

	for _, child := range children {
		// Ignore empty children
		if len(child.Data) == 0 {
			continue
		}

		switch state {
		case 0:
			if child.Type == xmlquery.TextNode {
				p_node = new_node_element("p")
				root.AddChild(p_node)
				if !has_label {
					lbl_node := new_node_element("label")
					p_node.AddChild(lbl_node)
					has_label = true
				}
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
		return tag_has_label(tag) || tag == "set-meta"
	}
	return node.Type == xmlquery.TextNode
}

func new_node_element(name string) *xmlquery.Node {
	ans := new(xmlquery.Node)
	ans.Type = xmlquery.ElementNode
	ans.Data = name

	return ans
}

func new_node_text(text string) *xmlquery.Node {
	ans := new(xmlquery.Node)
	ans.Type = xmlquery.TextNode
	ans.Data = text

	return ans
}
