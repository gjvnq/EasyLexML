package easyLexML

import (
	"bytes"
	"encoding/xml"
	"io"
)

func Draft2Strict(input io.Reader, output io.Writer) error {
	// Prepare stuff
	label_config := labelConfig{
		Sec: "Sec. {}",
		Cls: "Cls. {}",
		Sub: "¶ {}",
		Toc: "Table of Contents",
	}

	cls_counter := 0
	next_label := "???"
	label_pending := false
	root := new(xmlTreeNode)
	cursor := root
	toc_root := new(tocTreeNode)
	toc_cursor := toc_root
	make_toc := true

	// Start reading XML (step 1)
	decoder := xml.NewDecoder(input)
	step1_output := new(bytes.Buffer)
	encoder := xml.NewEncoder(step1_output)
	encoder.Indent("", "\t")
	for {
		token, err := decoder.RawToken()
		if err == io.EOF {
			break
		}
		panicIfErr(err)
		token = sanitize_char_data(token)

		Debugf("\n\n")
		Debugln("[tree_path]", cursor.PathToHere())
		Debugln("<<<", token2string(token))

		switch tk := token.(type) {
		case xml.StartElement:
			tag := name2string(tk.Name)

			switch tag {
			case "label-signs":
				if sec, ok := token_get_attr(tk, "sec"); ok {
					label_config.Sec = sec
				}
				if cls, ok := token_get_attr(tk, "cls"); ok {
					label_config.Cls = cls
				}
				if sub, ok := token_get_attr(tk, "sub"); ok {
					label_config.Sub = sub
				}
				if toc, ok := token_get_attr(tk, "toc"); ok {
					label_config.Toc = toc
				}
			case "cls":
				cls_counter++
			case "corpus":
				token_set_attr(&tk, "id", "corpus")
			case "reset-cls-counter":
				cls_counter = 0
				continue
			case "option":
				key, _ := token_get_attr(tk, "key")
				switch key {
				case "no-toc":
					make_toc = false
					continue
				}
			}

			// Finish any pending </p>
			if tag_has_label(tag) && cursor.PathToHere().PeekTag() == "p" {
				tk_p := xml.StartElement{Name: xml.Name{Local: "p"}}.End()
				cursor = cursor.Parent
				Debugln(">>>", token2string(tk_p))
				panicIfErr(encoder.EncodeToken(tk_p))
			}
			cursor = cursor.AddChild(tk)

			if cursor.PathToHere().Has("toc") {
				// the table of contents will be auto generated latter
				continue
			}
			if cursor.PathToHere().Has("metadata") {
				// ignore metadata for now
				Debugln(">>>", token2string(tk))
				panicIfErr(encoder.EncodeToken(tk))
				continue
			}

			// Set id and lexid for elements like <cls>, <sec>, <sub>
			if tag_has_id_in_stack(tag) {
				lexid := cursor.PathToHere().LexId()
				token_set_attr(&tk, "lexid", lexid)
				token_set_attr(&tk, "id", lexid+"_v1")
			}

			// Add <label>
			if label_pending {
				add_label(encoder, next_label)
				label_pending = false
			}
			if tag_has_label(tag) {
				label_pending = true
				next_label = cursor.PathToHere().Label(cls_counter, label_config)
				if next_label == "" {
					label_pending = false
				}
			}

			// Add TOC entry
			if tag == "sec" {
				id, _ := token_get_attr(tk, "id")
				toc_cursor = toc_cursor.AddEntry(next_label, id)
			}

			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		case xml.CharData:
			cursor.AddChild(tk)
			// Ignore whitespaces
			if is_token_empty(tk) {
				continue
			}

			if cursor.PathToHere().Has("toc") {
				continue
			}
			// In the corpus, all text MUST be inside paragraphs
			if cursor.PathToHere().Has("corpus") && !cursor.PathToHere().Has("p") {
				// Add <p>
				tk_p := xml.StartElement{Name: xml.Name{Local: "p"}}
				cursor = cursor.AddChild(tk_p)
				lexid := cursor.PathToHere().LexId()
				token_set_attr(&tk_p, "lexid", lexid)
				token_set_attr(&tk_p, "id", lexid+"_v1")
				Debugln(">>>", token2string(tk_p))
				panicIfErr(encoder.EncodeToken(tk_p))
				// Add <label>
				if label_pending {
					add_label(encoder, next_label)
					label_pending = false
				}
			}

			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		case xml.EndElement:
			tag := name2string(tk.Name)
			if cursor.PathToHere().Has("toc") {
				// the table of contents will be auto generated latter
				cursor = cursor.Parent
				continue
			}
			if tag == "reset-cls-counter" || tag == "option" {
				continue
			}
			if tag == "sec" {
				toc_cursor = toc_cursor.Parent
			}

			// End all tags we might have opened
			for tag != cursor.PathToHere().PeekTag() {
				Debugln(tag, cursor.PathToHere().PeekTag())
				tk2 := cursor.Token.(xml.StartElement).End()
				cursor = cursor.Parent
				Debugln(">>>", token2string(tk2))
				panicIfErr(encoder.EncodeToken(tk2))
				Debugln("[path]", cursor.PathToHere())
			}
			cursor = cursor.Parent
			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		default:
			if cursor.PathToHere().Has("toc") {
				continue
			}
			panicIfErr(encoder.EncodeToken(tk))
		}
	}

	// Generate TOC & finalize
	Debugf("\n\n=====STEP 2=====\n\n")
	encoder.Flush()
	decoder = xml.NewDecoder(step1_output)
	encoder = xml.NewEncoder(output)
	encoder.Indent("", "\t")
	Debugln(toc_root.ToXML())
	root = new(xmlTreeNode)
	cursor = root
	printed_toc := false
	for {
		token, err := decoder.RawToken()
		if err == io.EOF {
			break
		}
		panicIfErr(err)
		token = sanitize_char_data(token)

		Debugf("\n\n")
		Debugln("[tree_path]", cursor.PathToHere())
		Debugln("<<<", token2string(token))

		encoder.Flush()
		switch tk := token.(type) {
		case xml.StartElement:
			tag := name2string(tk.Name)
			cursor = cursor.AddChild(tk)
			if make_toc && !printed_toc && (tag == "toc" || tag == "corpus") {
				toc_root.ToXMLWithEncoder(encoder, label_config.Toc)
			}
			if cursor.PathToHere().Has("toc") {
				// the table of contents will be auto generated latter
				continue
			}
			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		case xml.CharData:
			// Ignore whitespaces
			if is_token_empty(tk) {
				continue
			}
			if cursor.PathToHere().Has("toc") {
				continue
			}
			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		case xml.EndElement:
			// tag := name2string(tk.Name)
			if cursor.PathToHere().Has("toc") {
				// the table of contents will be auto generated latter
				cursor = cursor.Parent
				continue
			}
			cursor = cursor.Parent
			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		default:
			panicIfErr(encoder.EncodeToken(tk))
		}
	}
	encoder.Flush()

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
