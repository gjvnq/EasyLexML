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
	// counter_stack := newCounterStack()
	// sec_sign := "§{}"
	// cls_sign := "Art. {}"
	// sub_sign := "{}."

	// Start reading XML
	last_label := "???"
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

			tag := tk.Name.Local
			switch {
			case tag == "cls":
				last_label = "??cls??"
				label_pending = true
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
				add_label(encoder, last_label)
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

	return nil
}

func add_label(encoder *xml.Encoder, last_label string) {
	tk_lbl := xml.StartElement{Name: xml.Name{Local: "label"}}
	Debugln(">>>", token2string(tk_lbl))
	panicIfErr(encoder.EncodeToken(tk_lbl))

	tk_txt := xml.CharData(last_label)
	Debugln(">>>", token2string(tk_txt))
	panicIfErr(encoder.EncodeToken(tk_txt))

	Debugln(">>>", token2string(tk_lbl.End()))
	panicIfErr(encoder.EncodeToken(tk_lbl.End()))

	tk_space := xml.CharData(" ")
	Debugln(">>>", token2string(tk_space))
	panicIfErr(encoder.EncodeToken(tk_space))
}
