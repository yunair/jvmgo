package rtda

import "github.com/yunair/jvmgo/ch08/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
