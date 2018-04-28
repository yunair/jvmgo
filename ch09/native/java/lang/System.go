package lang

import (
	"github.com/yunair/jvmgo/ch09/native"
	"github.com/yunair/jvmgo/ch09/msg"
	"github.com/yunair/jvmgo/ch09/rtda"
	"github.com/yunair/jvmgo/ch09/rtda/heap"
)

func init() {
	native.Register(msg.System, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtda.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if src == nil || dest == nil {
		panic(msg.NullPointerException)
	}

	if !checkArrayCopy(src, dest) {
		panic(msg.ArrayStoreException)
	}

	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic(msg.IndexOutOfBoundsException)
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}
func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}

	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}
