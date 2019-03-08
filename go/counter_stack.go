package easyLexML

type counterStackFrame struct {
	Main int
	Rev  int
}

func (this counterStackFrame) String() string {
	return counter2string(this.Main, this.Rev)
}

type counterStack struct {
	Frames []*counterStackFrame
}

func newCounterStack() *counterStack {
	ans := new(counterStack)
	ans.Frames = make([]*counterStackFrame, 0)
	return ans
}

func (this *counterStack) Len() int {
	return len(this.Frames)
}

func (this *counterStack) Peek() *counterStackFrame {
	return this.Frames[this.Len()-1]
}

func (this *counterStack) Push(frame *counterStackFrame) {
	this.Frames = append(this.Frames, frame)
}

func (this *counterStack) Pop() *counterStackFrame {
	ans := this.Peek()
	this.Frames = this.Frames[:this.Len()-1]
	return ans
}

func tag_has_id_in_stack(tag string) bool {
	return !(tag == "corpus" || tag == "toc" || tag == "metadata")
}

func tag_has_label(tag string) bool {
	return tag == "sec" || tag == "cls" || tag == "sub"
}
