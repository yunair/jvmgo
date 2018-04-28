package stack

import (
	"github.com/yunair/jvmgo/ch05/instructions/base"
	"github.com/yunair/jvmgo/ch05/rtda"
)

// dup 复制栈顶单个变量
type DUP struct { base.NoOperandsInstruction }

func (*DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

type DUP_X1 struct { base.NoOperandsInstruction }
type DUP_X2 struct { base.NoOperandsInstruction }
type DUP2 struct { base.NoOperandsInstruction }
type DUP2_X1 struct { base.NoOperandsInstruction }
type DUP2_X2 struct { base.NoOperandsInstruction }
