package stores

import (
	"github.com/yunair/jvmgo/ch08/instructions/base"
	"github.com/yunair/jvmgo/ch08/rtda"
)

type AASTORE struct{ base.NoOperandsInstruction }
type BASTORE struct{ base.NoOperandsInstruction }
type CASTORE struct{ base.NoOperandsInstruction }
type DASTORE struct{ base.NoOperandsInstruction }
type IASTORE struct{ base.NoOperandsInstruction }

func (*IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	base.CheckNotNil(arrRef)
	ints := arrRef.Ints()
	base.CheckIndex(len(ints), index)
	ints[index] = int32(val)
}

type LASTORE struct{ base.NoOperandsInstruction }

func (*LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	base.CheckNotNil(arrRef)
	longs := arrRef.Longs()
	base.CheckIndex(len(longs), index)
	longs[index] = int64(val)
}

type FASTORE struct{ base.NoOperandsInstruction }
type SASTORE struct{ base.NoOperandsInstruction }
