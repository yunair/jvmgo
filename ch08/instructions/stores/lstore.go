package stores

import (
	"github.com/yunair/jvmgo/ch08/instructions/base"
	"github.com/yunair/jvmgo/ch08/rtda"
)

type LSTORE struct{ base.Index8Instruction }

func (self *LSTORE) Execute(frame *rtda.Frame) {
	inLStore(frame, self.Index)
}

type LSTORE_0 struct{ base.NoOperandsInstruction }

func (*LSTORE_0) Execute(frame *rtda.Frame) {
	inLStore(frame, 0)
}

type LSTORE_1 struct{ base.NoOperandsInstruction }

func (*LSTORE_1) Execute(frame *rtda.Frame) {
	inLStore(frame, 1)
}

type LSTORE_2 struct{ base.NoOperandsInstruction }

func (*LSTORE_2) Execute(frame *rtda.Frame) {
	inLStore(frame, 2)
}

type LSTORE_3 struct{ base.NoOperandsInstruction }

func (*LSTORE_3) Execute(frame *rtda.Frame) {
	inLStore(frame, 3)
}

func inLStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
