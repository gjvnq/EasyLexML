package easyLexML

import (
	"encoding/xml"
	"strconv"
)

type counterStackFrame struct {
	Main int
	Rev  int
}

func (this *counterStackFrame) String() string {
	return counter2string(this.Main, this.Rev)
}

func (this *counterStackFrame) LexId(elem string) string {
	ans := elem + strconv.Itoa(this.Main)
	if this.Rev != 0 {
		ans += "-" + strconv.Itoa(this.Rev)
	}
	return ans
}

type counterStack struct {
	ClsCounter counterStackFrame
	Frames     []*counterStackFrame
	Elems      []string
}

func newCounterStack() *counterStack {
	ans := new(counterStack)
	ans.Frames = make([]*counterStackFrame, 1)
	ans.Frames[0] = new(counterStackFrame)
	ans.Elems = make([]string, 0)
	ans.ClsCounter.Main = 0
	return ans
}

func (this *counterStack) Len() int {
	return len(this.Frames)
}

func (this *counterStack) Peek() *counterStackFrame {
	return this.Frames[this.Len()-1]
}

func (this *counterStack) PrePeek() *counterStackFrame {
	return this.Frames[this.Len()-2]
}

func (this *counterStack) Push(elem string, frame *counterStackFrame) {
	this.Frames = append(this.Frames, frame)
	this.Elems = append(this.Elems, elem)
}

func (this *counterStack) PushOrUpdate(token xml.StartElement) {
	tag := name2string(token.Name)
	if tag == "EasyLexML" || tag == "corpus" {
		return
	}
	if tag == "cls" {
		this.ClsCounter.Main++
	}
	if tag == "sec" {
		val, _ := token_get_attr(token, "major")
		if val == "true" {
			this.ClsCounter.Main = 0
			this.ClsCounter.Rev = 0
			Debugln("reset counter")
		}
	}

	if this.Len() >= 1 {
		this.Peek().Main++
	}
	this.Push(tag, &counterStackFrame{
		Main: 0,
		Rev:  0,
	})
}

func (this *counterStack) Pop() *counterStackFrame {
	ans := this.Peek()
	this.Frames = this.Frames[:this.Len()-1]
	this.Elems = this.Elems[:this.Len()-1]
	return ans
}

func (this *counterStack) LexId() string {
	ans := ""
	last_elem := ""
	Debugln(this)
	for i, _ := range this.Elems {
		elem := this.Elems[i]
		frame := this.Frames[i]
		Debugln(ans)
		if last_elem != this.Elems[i] {
			if i != 0 {
				ans += "_"
			}
			last_elem = elem
			ans += frame.LexId(elem)
		} else {
			ans += "."
			ans += frame.LexId("")
		}
		Debugln(ans)
	}
	return ans
}

func (this *counterStack) Label(tk xml.StartElement) string {
	if name2string(tk.Name) == "cls" {
		return this.ClsCounter.String()
	}
	if this.Len() == 0 {
		return ""
	} else {
		return this.PrePeek().String()
	}
}

func tag_has_id_in_stack(tag string) bool {
	return !(tag == "EasyLexML" || tag == "corpus" || tag == "toc" || tag == "metadata")
}

func tag_has_label(tag string) bool {
	return tag == "sec" || tag == "cls" || tag == "sub"
}
