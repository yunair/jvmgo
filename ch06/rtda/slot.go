package rtda

import "github.com/yunair/jvmgo/ch06/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
