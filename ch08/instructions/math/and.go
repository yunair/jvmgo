package math

import (
	"github.com/yunair/jvmgo/ch08/instructions/base"
	"github.com/yunair/jvmgo/ch08/rtda"
)

type IAND struct { base.NoOperandsInstruction }

func (*IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	bool1 := stack.PopInt()
	bool2 := stack.PopInt()
	stack.PushInt(bool1 & bool2)
}

type LAND struct { base.NoOperandsInstruction }
