package comparisons

import (
	"github.com/yunair/jvmgo/ch06/instructions/base"
	"github.com/yunair/jvmgo/ch06/rtda"
)

// float 比较，因为可能存在NaN，所以存在无法比较的结果
// 两个指令集的不同是对无法比较的情况处理的不同

type FMCPG struct { base.NoOperandsInstruction}

func (*FMCPG) Execute(frame *rtda.Frame) {
	fcmp(frame, true)
}

type FMCPL struct { base.NoOperandsInstruction}

func (*FMCPL) Execute(frame *rtda.Frame) {
	fcmp(frame, false)
}

func fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}

}
