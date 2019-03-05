package easyLexML

type IdStackFrame struct {
	Main int
	Rev  int
}

func (this IdStackFrame) String() string {
	return counter2string(this.Main, this.Rev)
}

type IdStack struct {
	Frames []*IdStackFrame
}

func NewIdStack() *IdStack {
	ans := new(IdStack)
	ans.Frames = make([]*IdStackFrame, 0)
	return ans
}

func (this *IdStack) Len() int {
	return len(this.Frames)
}

func (this *IdStack) Peek() *IdStackFrame {
	return this.Frames[this.Len()-1]
}

func (this *IdStack) Push(frame *IdStackFrame) {
	this.Frames = append(this.Frames, frame)
}

func (this *IdStack) Pop() *IdStackFrame {
	ans := this.Peek()
	this.Frames = this.Frames[:this.Len()-1]
	return ans
}
