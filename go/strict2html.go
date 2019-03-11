package easyLexML

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"io"
)

type htmlPage struct {
	Title  string
	Corpus template.HTML
	Toc    template.HTML
}

func Strict2HTML(input io.Reader, output io.Writer) error {
	root := new(xmlTreeNode)
	cursor := root
	data := htmlPage{}

	// Convert XML to HTML
	decoder := xml.NewDecoder(input)
	corpus_buffer := new(bytes.Buffer)
	corpus_encoder := xml.NewEncoder(corpus_buffer)
	corpus_encoder.Indent("\t\t", "\t")
	toc_buffer := new(bytes.Buffer)
	toc_encoder := xml.NewEncoder(toc_buffer)
	toc_encoder.Indent("\t\t", "\t")
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
		corpus_encoder.Flush()

		switch tk := token.(type) {
		case xml.StartElement:
			tag := name2string(tk.Name)
			cursor = cursor.AddChild(tk)
			path := cursor.PathToHere()
			switch {
			case path.Has("metadata"):
				continue
			case tag == "toc":
				continue
			case path.Has("toc"):
				if tag == "label" {
					tk.Name.Local = "h2"
				}
				Debugln("TOC: >>>", token2string(tk))
				panicIfErr(toc_encoder.EncodeToken(tk))
				continue
			case tag == "EasyLexML" || tag == "corpus":
				continue
			case tag == "cls" || tag == "sec" || tag == "sub":
				tk.Name.Local = "section"
				token_set_attr(&tk, "class", tag)
			case tag == "label":
				class, _ := token_get_attr(path.PeekN(2).(xml.StartElement), "class")
				if class == "sec" {
					tk.Name.Local = "h3"
				} else {
					tk.Name.Local = "span"
					token_set_attr(&tk, "class", "label")
				}
			}
			cursor.Replace(tk)
			Debugln(">>>", token2string(tk))
			panicIfErr(corpus_encoder.EncodeToken(tk))
		case xml.CharData:
			// Ignore whitespaces
			if is_token_empty(tk) {
				continue
			}
			path := cursor.PathToHere()
			if path.Has("toc") {
				Debugln("TOC: >>>", token2string(tk))
				panicIfErr(toc_encoder.EncodeToken(tk))
				continue
			}
			if path.Has("metadata") && path.PeekTag() == "title" {
				data.Title = string(tk)
			}
			if path.Has("metadata") || path.Has("toc") {
				continue
			}
			Debugln(">>>", token2string(tk))
			panicIfErr(corpus_encoder.EncodeToken(tk))
		case xml.EndElement:
			tag := name2string(tk.Name)
			path := cursor.PathToHere()
			cursor = cursor.Parent
			switch {
			case path.Has("metadata"):
				continue
			case tag == "toc":
				tk_hr := xml.StartElement{Name: xml.Name{Local: "hr"}}
				Debugln("TOC: >>>", token2string(tk_hr))
				panicIfErr(toc_encoder.EncodeToken(tk_hr))
				Debugln("TOC: >>>", token2string(tk_hr.End()))
				panicIfErr(toc_encoder.EncodeToken(tk_hr.End()))
				continue
			case path.Has("toc"):
				if tag == "label" {
					tk.Name.Local = "h2"
				}
				Debugln("TOC: >>>", token2string(tk))
				panicIfErr(toc_encoder.EncodeToken(tk))
				continue
			case tag == "EasyLexML" || tag == "corpus" || tag == "toc" || tag == "metadata":
				continue
			case tag == "cls" || tag == "sec" || tag == "sub":
				tk.Name.Local = "section"
			case tag == "label":
				class, _ := token_get_attr(path.PeekN(2).(xml.StartElement), "class")
				if class == "sec" {
					tk.Name.Local = "h3"
				} else {
					tk.Name.Local = "span"
				}
			}
			Debugln(">>>", token2string(tk))
			panicIfErr(corpus_encoder.EncodeToken(tk))
		}
	}
	toc_encoder.Flush()
	corpus_encoder.Flush()

	// Output
	data.Corpus = template.HTML(corpus_buffer.String())
	data.Toc = template.HTML(toc_buffer.String())
	tmpl := template.Must(template.ParseFiles("res/standalone.html"))
	tmpl.Execute(output, data)

	return nil
}
