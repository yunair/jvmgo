package references

import (
	"github.com/yunair/jvmgo/ch07/instructions/base"
	"github.com/yunair/jvmgo/ch07/rtda"
	"github.com/yunair/jvmgo/ch07/rtda/heap"
	"github.com/yunair/jvmgo/ch07/msg"
)

type INVOKE_INTERFACE struct {
	index uint
	//count uint8    // 方法参数所需的slot数
	//zero uint8    // 向后兼容
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.ByteCodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8()
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic(msg.IncompatibleClassChangeError)
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1) // 弹出this引用

	if ref == nil {
		panic(msg.NullPointerException)
	}

	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic(msg.IncompatibleClassChangeError)
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic(msg.AbstractMethodError)
	}

	if !methodToBeInvoked.IsPublic() {
		panic(msg.IllegalAccessError)
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}



