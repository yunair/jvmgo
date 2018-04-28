package base

import "github.com/yunair/jvmgo/ch08/rtda"

type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *ByteCodeReader)

	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {}

func (self *NoOperandsInstruction)FetchOperands(reader *ByteCodeReader)  {
	// do nothing
}

// 跳转指令
type BranchInstruction struct {
	Offset int  // 跳转偏移量
}

func (self *BranchInstruction)FetchOperands(reader *ByteCodeReader)  {
	self.Offset = int(reader.ReadInt16())
}


type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction)FetchOperands(reader *ByteCodeReader)  {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction)FetchOperands(reader *ByteCodeReader)  {
	self.Index = uint(reader.ReadUint16())
}