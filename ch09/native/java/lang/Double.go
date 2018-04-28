package lang

import (
	"github.com/yunair/jvmgo/ch09/native"
	"github.com/yunair/jvmgo/ch09/msg"
	"github.com/yunair/jvmgo/ch09/rtda"
	"math"
)

func init() {
	native.Register(msg.Double, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}