package easyLexML

import (
	"encoding/xml"
	"io"
)

func Draft2Strict(input io.Reader, output io.Writer) error {
	// Prepare stuff
	label_config := labelConfig{
		Sec: "Sec. {}",
		Cls: "Cls. {}",
		Sub: "¶ {}",
	}

	cls_counter := 0
	next_label := "???"
	label_pending := false
	root := new(xmlTreeNode)
	cursor := root

	// Start reading XML
	decoder := xml.NewDecoder(input)
	encoder := xml.NewEncoder(output)
	encoder.Indent("", "\t")
	defer encoder.Flush()
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
			case "cls":
				cls_counter++
			case "sec":
				major, ok := token_get_attr(tk, "major")
				if ok && major == "true" {
					cls_counter = 0
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

			// switch tag {
			// case "corpus":
			// 	in_corpus = false
			// }

			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		default:
			if cursor.PathToHere().Has("toc") {
				continue
			}
			panicIfErr(encoder.EncodeToken(tk))
		}
	}
	// Finalize

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
