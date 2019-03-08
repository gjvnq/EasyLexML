package easyLexML

import (
	"encoding/xml"
	"fmt"
)

type elementStack struct {
	Data []xml.StartElement
}

func newElementStack() *elementStack {
	ans := new(elementStack)
	ans.Data = make([]xml.StartElement, 0)
	return ans
}

func (this *elementStack) Push(elem xml.StartElement) {
	this.Data = append(this.Data, elem.Copy())
}

func (this *elementStack) Peek() xml.StartElement {
	return this.Data[this.Len()-1]
}

func (this *elementStack) Pop() xml.StartElement {
	ans := this.Data[this.Len()-1]
	this.Data = this.Data[:this.Len()-1]
	return ans
}

func (this *elementStack) PeekTag() string {
	ans := this.Peek()
	return name2string(ans.Name)
}

func (this *elementStack) Has(tag string) bool {
	for _, elem := range this.Data {
		if name2string(elem.Name) == tag {
			return true
		}
	}
	return false
}

func (this *elementStack) Len() int {
	return len(this.Data)
}

func (this *elementStack) String() string {
	ans := ""
	for i, elem := range this.Data {
		if i != 0 {
			ans += " "
		}
		ans += "<" + name2string(elem.Name)
		for _, attr := range elem.Attr {
			ans += " "
			ans += fmt.Sprintf("%s=%q", name2string(attr.Name), attr.Value)
		}

		ans += ">"
	}
	return ans
}
