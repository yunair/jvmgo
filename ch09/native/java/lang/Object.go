package lang

import (
	"github.com/yunair/jvmgo/ch09/native"
	"github.com/yunair/jvmgo/ch09/msg"
	"github.com/yunair/jvmgo/ch09/rtda"
)

func init() {
	native.Register(msg.Object, "getClass", "()Ljava/lang/Class;", getClass)
}

// public final native Class<?> getClass();
func getClass(frame *rtda.Frame){
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
