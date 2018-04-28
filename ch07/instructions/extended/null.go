package extended

import (
	"github.com/yunair/jvmgo/ch07/instructions/base"
	"github.com/yunair/jvmgo/ch07/rtda"
)

type IFNULL struct{ base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

type IFNONNULL struct{ base.BranchInstruction }
