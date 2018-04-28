package constants

import (
	"github.com/yunair/jvmgo/ch09/instructions/base"
	"github.com/yunair/jvmgo/ch09/rtda"
)

type ACONST_NULL struct{ base.NoOperandsInstruction }

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

type DCONST_0 struct{ base.NoOperandsInstruction }

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	inIConst(frame, -1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	inIConst(frame, 0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	inIConst(frame, 1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

func (*ICONST_2) Execute(frame *rtda.Frame) {
	inIConst(frame, 2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

func (*ICONST_3) Execute(frame *rtda.Frame) {
	inIConst(frame, 3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

func (*ICONST_4) Execute(frame *rtda.Frame) {
	inIConst(frame, 4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

func (*ICONST_5) Execute(frame *rtda.Frame) {
	inIConst(frame, 5)
}

type LCONST_0 struct{ base.NoOperandsInstruction }



func (*LCONST_0) Execute(frame *rtda.Frame) {
	inLConst(frame, 0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

func (*LCONST_1) Execute(frame *rtda.Frame) {
	inLConst(frame, 1)
}

func inLConst(frame *rtda.Frame, val int64) {
	frame.OperandStack().PushLong(val)
}


func inIConst(frame *rtda.Frame, val int32) {
	frame.OperandStack().PushInt(val)
}