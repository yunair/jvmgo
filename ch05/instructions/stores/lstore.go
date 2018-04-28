package stores

import (
	"github.com/yunair/jvmgo/ch05/instructions/base"
	"github.com/yunair/jvmgo/ch05/rtda"
)

type LSTORE struct { base.Index8Instruction }

func (self *LSTORE) Execute(frame *rtda.Frame) {
	inLStore(frame, self.Index)
}

type LLOAD_0 struct { base.NoOperandsInstruction }
type LLOAD_1 struct { base.NoOperandsInstruction }
type LLOAD_2 struct { base.NoOperandsInstruction }
func (*LLOAD_2) Execute(frame *rtda.Frame) {
	inLStore(frame, 2)
}

type LLOAD_3 struct { base.NoOperandsInstruction }

func inLStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}