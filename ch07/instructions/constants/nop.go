package constants

import (
	"github.com/yunair/jvmgo/ch07/instructions/base"
	"github.com/yunair/jvmgo/ch07/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}


func (self *NOP) Execute(frame *rtda.Frame) {

}
