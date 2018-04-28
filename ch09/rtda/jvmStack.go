package rtda

type Stack struct {
	maxSize uint
	size uint
	top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) isEmpty() bool {
	return self.top == nil
}

func (self *Stack) pop() *Frame{
	if self.top == nil {
		panic("jvm stack is empty")
	}

	top := self.top
	self.top = top.lower
	top.lower = nil
	self.size--
	return top
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if self.top != nil {
		frame.lower = self.top
	}

	self.top = frame
	self.size++
}