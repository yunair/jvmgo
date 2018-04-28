package lang

import (
	"github.com/yunair/jvmgo/ch09/native"
	"github.com/yunair/jvmgo/ch09/msg"
	"github.com/yunair/jvmgo/ch09/rtda"
	"github.com/yunair/jvmgo/ch09/rtda/heap"
)

func init() {
	native.Register(msg.Class, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(msg.Class, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(msg.Class, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

// static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *rtda.Frame){
	nameObj := frame.LocalVars().GetThis()
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}

func getName0(frame *rtda.Frame){
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)
	frame.OperandStack().PushRef(nameObj)
}

func desiredAssertionStatus0(frame *rtda.Frame){
	//frame.OperandStack().PushBoolean(false)
	frame.OperandStack().PushInt(0)
}
