package constants

import (
	"github.com/yunair/jvmgo/ch06/instructions/base"
	"github.com/yunair/jvmgo/ch06/rtda"
)

type LDC struct{ base.Index8Instruction }

func (self *LDC) Execute(frame *rtda.Frame) {
	inLdc(frame, self.Index)
}

type LDC_W struct{ base.Index16Instruction }

func (self *LDC_W) Execute(frame *rtda.Frame) {
	inLdc(frame, self.Index)
}

func inLdc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type){
	case int32: stack.PushInt(c.(int32))
	case float32: stack.PushFloat(c.(float32))
		//case string:
		// case *heap.ClassRef:
	default: panic("todo: ldc!")
	}
}

type LDC2_W struct{ base.Index16Instruction }

// 加载long double
func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)
	switch c.(type){
	case int64: stack.PushLong(c.(int64))
	case float64: stack.PushDouble(c.(float64))
	default: panic("todo: ldc!")
	}
}


