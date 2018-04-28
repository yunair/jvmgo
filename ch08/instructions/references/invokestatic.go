package references

import (
	"github.com/yunair/jvmgo/ch08/instructions/base"
	"github.com/yunair/jvmgo/ch08/rtda"
	"github.com/yunair/jvmgo/ch08/rtda/heap"
	"github.com/yunair/jvmgo/ch08/msg"
)

// Invoke a class (static) method
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	class := resolvedMethod.Class()

	if !class.InitStarted(){
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !resolvedMethod.IsStatic() {
		panic(msg.IncompatibleClassChangeError)
	}
	base.InvokeMethod(frame, resolvedMethod)
}
