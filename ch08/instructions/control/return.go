package control

import (
	"github.com/yunair/jvmgo/ch08/instructions/base"
	"github.com/yunair/jvmgo/ch08/rtda"
)

type RETURN struct{ base.NoOperandsInstruction } // Return void from method

func (*RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct{ base.NoOperandsInstruction } // Return ref from method

func (*ARETURN) Execute(frame *rtda.Frame) {
	inReturn(frame, func(pop *rtda.OperandStack, push *rtda.OperandStack) {
		push.PushRef(pop.PopRef())
	})
}

type DRETURN struct{ base.NoOperandsInstruction } // Return double from method

func (*DRETURN) Execute(frame *rtda.Frame) {
	inReturn(frame, func(pop *rtda.OperandStack, push *rtda.OperandStack) {
		push.PushDouble(pop.PopDouble())
	})
}

type FRETURN struct{ base.NoOperandsInstruction } // Return float from method

func (*FRETURN) Execute(frame *rtda.Frame) {
	inReturn(frame, func(pop *rtda.OperandStack, push *rtda.OperandStack) {
		push.PushFloat(pop.PopFloat())
	})
}

type IRETURN struct{ base.NoOperandsInstruction } // Return int from method

func (*IRETURN) Execute(frame *rtda.Frame) {
	inReturn(frame, func(pop *rtda.OperandStack, push *rtda.OperandStack) {
		push.PushInt(pop.PopInt())
	})
}

type LRETURN struct{ base.NoOperandsInstruction } // Return long from method

func (*LRETURN) Execute(frame *rtda.Frame) {
	inReturn(frame, func(pop *rtda.OperandStack, push *rtda.OperandStack) {
		push.PushLong(pop.PopLong())
	})
}

func inReturn(frame *rtda.Frame, operate func(*rtda.OperandStack, *rtda.OperandStack)){
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	operate(currentFrame.OperandStack(), invokerFrame.OperandStack())
}
