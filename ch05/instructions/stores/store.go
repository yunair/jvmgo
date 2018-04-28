package stores

import (
	"github.com/yunair/jvmgo/ch05/instructions/base"
	"github.com/yunair/jvmgo/ch05/rtda"
)

type FSTORE struct { base.Index8Instruction }

func (self *FSTORE) Execute(frame *rtda.Frame) {
	inLStore(frame, self.Index)
}

type DSTORE struct { base.Index8Instruction }

func (self *DSTORE) Execute(frame *rtda.Frame) {
	inLStore(frame, self.Index)
}

type ASTORE struct { base.Index8Instruction }

func (self *ASTORE) Execute(frame *rtda.Frame) {
	inLStore(frame, self.Index)
}

