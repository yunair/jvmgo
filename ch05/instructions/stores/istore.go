package stores

import (
	"github.com/yunair/jvmgo/ch05/instructions/base"
	"github.com/yunair/jvmgo/ch05/rtda"
)

// Store int into local variable
type ISTORE struct{ base.Index8Instruction }

func (self *ISTORE) Execute(frame *rtda.Frame) {
	inIstore(frame, uint(self.Index))
}

type ISTORE_0 struct{ base.NoOperandsInstruction }

func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	inIstore(frame, 0)
}

type ISTORE_1 struct{ base.NoOperandsInstruction }

func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	inIstore(frame, 1)
}

type ISTORE_2 struct{ base.NoOperandsInstruction }

func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	inIstore(frame, 2)
}

type ISTORE_3 struct{ base.NoOperandsInstruction }

func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	inIstore(frame, 3)
}

func inIstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
