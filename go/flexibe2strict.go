package easyLexML

import (
	"encoding/xml"
	"io"
)

func Draft2Strict(input io.Reader, output io.Writer) error {
	// Prepare stuff
	// known_ids := make(map[string]bool)
	// cls_counter := 0
	elem_stack := newElementStack()
	counter_stack := newCounterStack()
	// sec_sign := "§{}"
	// cls_sign := "Art. {}"
	// sub_sign := "{}."

	// Start reading XML
	next_label := "???"
	label_pending := false
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

		Debugf("\n\n\n")
		Debugln("[elem_stack]", elem_stack)
		Debugln("[counter_stack]", counter_stack)
		Debugln("<<<", token2string(token))

		switch tk := token.(type) {
		case xml.StartElement:
			tag := tk.Name.Local

			if tag_has_label(tag) && elem_stack.PeekTag() == "p" {
				tk_p := xml.StartElement{Name: xml.Name{Local: "p"}}.End()
				elem_stack.Pop()
				counter_stack.Pop()
				Debugln(">>>", token2string(tk_p))
				panicIfErr(encoder.EncodeToken(tk_p))
			}

			elem_stack.Push(tk)
			if elem_stack.Has("toc") {
				// the table of contents will be auto generated latter
				continue
			}
			if elem_stack.Has("metadata") {
				// ignore metadata for now
				Debugln(">>>", token2string(tk))
				panicIfErr(encoder.EncodeToken(tk))
				continue
			}

			// Set id and lexid for elements like <cls>, <sec>, <sub>
			if tag_has_id_in_stack(tag) {
				counter_stack.PushOrUpdate(tk)
				lexid := counter_stack.LexId()
				token_set_attr(&tk, "lexid", lexid)
				token_set_attr(&tk, "id", lexid+"_v1")
			}
			if tag_has_label(tag) {
				label_pending = true
				next_label = counter_stack.Label(tk)
			}
			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		case xml.CharData:
			// Ignore whitespaces
			if is_token_empty(tk) {
				continue
			}

			if elem_stack.Has("toc") {
				continue
			}
			// In the corpus, all text MUST be inside paragraphs
			if elem_stack.Has("corpus") && !elem_stack.Has("p") {
				// Add <p>
				tk_p := xml.StartElement{Name: xml.Name{Local: "p"}}
				elem_stack.Push(tk_p)
				counter_stack.PushOrUpdate(tk_p)
				lexid := counter_stack.LexId()
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
			tag := tk.Name.Local
			if elem_stack.Has("toc") {
				// the table of contents will be auto generated latter
				elem_stack.Pop()
				continue
			}
			if !elem_stack.Has("metadata") && tag_has_id_in_stack(tag) {
				counter_stack.Pop()
			}

			// End all tags we might have opened
			for tag != elem_stack.PeekTag() {
				tk2 := elem_stack.Pop().End()
				if tag_has_id_in_stack(name2string(tk2.Name)) {
					counter_stack.Pop()
				}
				Debugln(">>>", token2string(tk2))
				panicIfErr(encoder.EncodeToken(tk2))
				Debugln("[elem_stack]", elem_stack)
			}
			elem_stack.Pop()

			// switch tag {
			// case "corpus":
			// 	in_corpus = false
			// }

			Debugln(">>>", token2string(tk))
			panicIfErr(encoder.EncodeToken(tk))
		default:
			if elem_stack.Has("toc") {
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
