package references

import (
	"github.com/yunair/jvmgo/ch09/instructions/base"
	"github.com/yunair/jvmgo/ch09/rtda"
	"github.com/yunair/jvmgo/ch09/rtda/heap"
	"github.com/yunair/jvmgo/ch09/msg"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.Name() == "<init>" && resolvedClass != resolvedMethod.Class() {
		panic(msg.NoSuchMethodError)
	}
	if resolvedMethod.IsStatic() {
		panic(msg.IncompatibleClassChangeError)
	}
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic(msg.NullPointerException)
	}

	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic(msg.IllegalAccessError)
	}

	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
			resolvedMethod.Name() != "<init>" {
				methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
					methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic(msg.AbstractMethodError)
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
