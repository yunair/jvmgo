package references

import (
	"github.com/yunair/jvmgo/ch09/instructions/base"
	"github.com/yunair/jvmgo/ch09/rtda"
	"github.com/yunair/jvmgo/ch09/msg"
)

type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic(msg.NullPointerException)
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}


