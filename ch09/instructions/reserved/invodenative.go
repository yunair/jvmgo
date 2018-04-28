package reserved

import (
	"github.com/yunair/jvmgo/ch09/instructions/base"
	"github.com/yunair/jvmgo/ch09/rtda"
	"github.com/yunair/jvmgo/ch09/native"
	_ "github.com/yunair/jvmgo/ch09/native/java/lang"
	"github.com/yunair/jvmgo/ch09/msg"

)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (*INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic(msg.UnsatisfiedLinkError + ": " + methodInfo)
	}
	nativeMethod(frame)
}

