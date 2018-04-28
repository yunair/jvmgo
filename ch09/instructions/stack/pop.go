package stack

import (
	"github.com/yunair/jvmgo/ch09/instructions/base"
	"github.com/yunair/jvmgo/ch09/rtda"
)
// 弹出栈顶占用一个操作数栈位置的变量
type POP struct { base.NoOperandsInstruction }

func (*POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

type POP2 struct { base.NoOperandsInstruction }

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

