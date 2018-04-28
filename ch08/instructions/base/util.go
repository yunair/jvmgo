package base

import (
	"github.com/yunair/jvmgo/ch08/msg"
	"github.com/yunair/jvmgo/ch08/rtda/heap"
)

func CheckIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic(msg.ArrayIndexOutOfBoundsException)
	}
}
func CheckNotNil(ref *heap.Object) {
	if ref == nil {
		panic(msg.NullPointerException)
	}
}
