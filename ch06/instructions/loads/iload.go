package loads

import (
	"github.com/yunair/jvmgo/ch06/instructions/base"
	"github.com/yunair/jvmgo/ch06/rtda"
)

type ILOAD struct {
	base.Index8Instruction
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	inILoad(frame, uint(self.Index))
}

type ILOAD_0 struct { base.NoOperandsInstruction }
type ILOAD_1 struct { base.NoOperandsInstruction }

func (*ILOAD_1) Execute(frame *rtda.Frame) {
	inILoad(frame, 1)
}

type ILOAD_2 struct { base.NoOperandsInstruction }
type ILOAD_3 struct { base.NoOperandsInstruction }

func inILoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}