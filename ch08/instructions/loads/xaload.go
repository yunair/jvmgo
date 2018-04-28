package loads

import (
	"github.com/yunair/jvmgo/ch08/instructions/base"
	"github.com/yunair/jvmgo/ch08/rtda"
)

type AALOAD struct{ base.NoOperandsInstruction }

func (*AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	base.CheckNotNil(arrRef)
	refs := arrRef.Refs()
	base.CheckIndex(len(refs), index)
	stack.PushRef(refs[index])
}

type BALOAD struct{ base.NoOperandsInstruction }
type CALOAD struct{ base.NoOperandsInstruction }
type DALOAD struct{ base.NoOperandsInstruction }
type FALOAD struct{ base.NoOperandsInstruction }
type IALOAD struct{ base.NoOperandsInstruction }

func (*IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	base.CheckNotNil(arrRef)
	refs := arrRef.Ints()
	base.CheckIndex(len(refs), index)
	stack.PushInt(refs[index])
}

type LALOAD struct{ base.NoOperandsInstruction }

func (*LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	base.CheckNotNil(arrRef)
	refs := arrRef.Longs()
	base.CheckIndex(len(refs), index)
	stack.PushLong(refs[index])
}

type SALOAD struct{ base.NoOperandsInstruction }
