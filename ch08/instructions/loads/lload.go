package loads

import (
	"github.com/yunair/jvmgo/ch08/instructions/base"
	"github.com/yunair/jvmgo/ch08/rtda"
)

type LLOAD struct{ base.Index8Instruction }

func (self *LLOAD) Execute(frame *rtda.Frame) {
	inLLoad(frame, uint(self.Index))
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	inLLoad(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	inLLoad(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	inLLoad(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	inLLoad(frame, 3)
}


func inLLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
