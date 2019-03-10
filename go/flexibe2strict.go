package easyLexML

import (
	"encoding/xml"
	"fmt"
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

		Debugln("[elem_stack]", elem_stack)
		Debugln(token2string(token))

		switch tk := token.(type) {
		case xml.StartElement:
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

			tag := tk.Name.Local
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
				Debugln(">>>", token2string(tk_p))
				panicIfErr(encoder.EncodeToken(tk_p))
				// Add <label>
				add_label(encoder, next_label)
				label_pending = false
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
	fmt.Println(label_pending)

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
