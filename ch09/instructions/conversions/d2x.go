package conversions

import (
	"github.com/yunair/jvmgo/ch09/instructions/base"
	"github.com/yunair/jvmgo/ch09/rtda"
)

// 把 double变量强制转化为其它类型


type D2F struct { base.NoOperandsInstruction }
type D2I struct { base.NoOperandsInstruction }

func (*D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushInt(int32(d))
}

type D2L struct { base.NoOperandsInstruction }