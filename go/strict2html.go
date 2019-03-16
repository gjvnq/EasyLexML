package easyLexML

import (
	"errors"
	"html/template"
	"io"
)

type htmlPage struct {
	Title  string
	Corpus template.HTML
	Toc    template.HTML
}

func Strict2HTML(input io.Reader, output io.Writer) error {
	return errors.New("not implemented")
}
