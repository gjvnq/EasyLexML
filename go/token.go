package easyLexML

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

func name2string(name xml.Name) string {
	if name.Space == "" {
		return name.Local
	}
	return name.Space + ":" + name.Local
}

func token2string(token xml.Token) string {
	switch tk := token.(type) {
	case xml.CharData:
		return fmt.Sprintf("[CharData] \"%s\"", strings.TrimSpace(string(tk)))
	case xml.Directive:
		return fmt.Sprintf("[Directive] \"%s\"", string(tk))
	case xml.Comment:
		return fmt.Sprintf("[Comment] \"%s\"", string(tk))
	case xml.ProcInst:
		return fmt.Sprintf("[ProcInst] Target:\"%s\" Inst:\"%s\"", tk.Target, string(tk.Inst))
	case xml.StartElement:
		ans := fmt.Sprintf("[StartElement] <")
		if tk.Name.Space != "" {
			ans += fmt.Sprintf("%s:", tk.Name.Space)
		}
		ans += fmt.Sprintf("%s", tk.Name.Local)
		for _, attr := range tk.Attr {
			ans += fmt.Sprintf(" ")
			if attr.Name.Space != "" {
				ans += fmt.Sprintf("%s:", attr.Name.Space)
			}
			ans += fmt.Sprintf("%s=\"%s\"", attr.Name.Local, attr.Value)
		}
		ans += fmt.Sprintf(">")
		return ans
	case xml.EndElement:
		ans := fmt.Sprintf("[EndElement] <")
		if tk.Name.Space != "" {
			ans += fmt.Sprintf("%s:", tk.Name.Space)
		}
		ans += fmt.Sprintf("%s/>", tk.Name.Local)
		return ans
	default:
		return fmt.Sprintf("%+v", token)
	}
}

func sanitize_char_data(token xml.Token) xml.Token {
	switch tk := token.(type) {
	case xml.CharData:
		tmp := strings.TrimSpace(string(tk))
		re_internal_whitespace := regexp.MustCompile(`[\s\p{Zs}]+`)
		ans := re_internal_whitespace.ReplaceAllString(tmp, " ")
		return xml.CharData(ans)
	default:
		return token
	}
}

func is_token_empty(token xml.Token) bool {
	switch tk := token.(type) {
	case xml.CharData:
		return len(strings.TrimSpace(string(tk))) == 0
	default:
		return false
	}
}

func token_set_attr(token *xml.StartElement, key string, value interface{}) {
	val_str := fmt.Sprintf("%q", value)
	val_str = strings.TrimPrefix(val_str, "\"")
	val_str = strings.TrimSuffix(val_str, "\"")

	for i := range token.Attr {
		attr := &token.Attr[i]
		if attr.Name.Local == key {
			attr.Value = val_str
			return
		}
	}
	token.Attr = append(token.Attr, xml.Attr{
		Name:  xml.Name{Space: "", Local: key},
		Value: val_str,
	})
}

func token_has_attr(token xml.StartElement, key string) bool {
	_, ans := token_get_attr(token, key)
	return ans
}

func token_get_attr(token xml.StartElement, key string) (string, bool) {
	for _, attr := range token.Attr {
		if name2string(attr.Name) == key {
			return attr.Value, true
		}
	}
	return "", false
}

func token_requires_p(token xml.Token) bool {
	tag := ""
	switch tk := token.(type) {
	case xml.StartElement:
		tag = tk.Name.Local
	case xml.EndElement:
		tag = tk.Name.Local
	default:
		return false
	}
	return tag == "Artigo" || tag == "Caput" || tag == "Recital" || tag == "Parágrafo" || tag == "Inciso" || tag == "Alínea" || tag == "Item"
}

func token_same_element(t1, t2 xml.Token) bool {
	tag1 := ""
	tag2 := ""
	switch tk := t1.(type) {
	case xml.StartElement:
		tag1 = name2string(tk.Name)
	case xml.EndElement:
		tag1 = name2string(tk.Name)
	default:
		return false
	}
	switch tk := t2.(type) {
	case xml.StartElement:
		tag2 = name2string(tk.Name)
	case xml.EndElement:
		tag2 = name2string(tk.Name)
	default:
		return false
	}

	Debugln("token_same_element", tag1, tag2)
	return tag1 == tag2
}
