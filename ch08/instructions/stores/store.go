package stores

import (
	"github.com/yunair/jvmgo/ch08/instructions/base"
	"github.com/yunair/jvmgo/ch08/rtda"
)

type FSTORE struct { base.Index8Instruction }

func (self *FSTORE) Execute(frame *rtda.Frame) {
	//inLStore(frame, self.Index)
}

type DSTORE struct { base.Index8Instruction }

func (self *DSTORE) Execute(frame *rtda.Frame) {
	//inLStore(frame, self.Index)
}

type ASTORE struct { base.Index8Instruction }

func (self *ASTORE) Execute(frame *rtda.Frame) {
	inAStore(frame, self.Index)
}

type ASTORE_0 struct { base.NoOperandsInstruction }

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	inAStore(frame, 0)
}

type ASTORE_1 struct { base.NoOperandsInstruction }

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	inAStore(frame, 1)
}

type ASTORE_2 struct { base.NoOperandsInstruction }

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	inAStore(frame, 2)
}

type ASTORE_3 struct { base.NoOperandsInstruction }

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	inAStore(frame, 3)
}


func inAStore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
