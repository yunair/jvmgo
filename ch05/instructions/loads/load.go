package loads

import (
	"github.com/yunair/jvmgo/ch05/instructions/base"
	"github.com/yunair/jvmgo/ch05/rtda"
)

type LLOAD struct {
	base.Index8Instruction
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	inILoad(frame, uint(self.Index))
}

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

type ALOAD struct {
	base.Index8Instruction
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	inILoad(frame, uint(self.Index))
}
