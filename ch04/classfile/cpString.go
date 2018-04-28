package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	cp    ConstantPool
	index uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.index = reader.readUint16()
}

func(self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.index)
}