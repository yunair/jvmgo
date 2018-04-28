package math

import (
	"github.com/yunair/jvmgo/ch09/instructions/base"
	"github.com/yunair/jvmgo/ch09/rtda"
)

type ISHL struct{ base.NoOperandsInstruction } // int左移

func (*ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}


type ISHR struct{ base.NoOperandsInstruction } // int右移
type IUSHR struct{ base.NoOperandsInstruction } // int逻辑右移

func (*IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}


type LSHL struct{ base.NoOperandsInstruction } // long左移
type LSHR struct{ base.NoOperandsInstruction } // long右移
func (*LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint(v2) & 0x3f
	result := v1 >> s
	stack.PushInt(result)
}
type LUSHR struct{ base.NoOperandsInstruction } //long逻辑右移
