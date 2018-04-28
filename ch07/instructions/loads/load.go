package loads

import (
	"github.com/yunair/jvmgo/ch07/instructions/base"
	"github.com/yunair/jvmgo/ch07/rtda"
)

type FLOAD struct {
	base.Index8Instruction
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	inILoad(frame, uint(self.Index))
}

type DLOAD struct {
	base.Index8Instruction
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	inILoad(frame, uint(self.Index))
}


