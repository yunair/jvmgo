package classfile


/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	cp    ConstantPool
	index uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.index = reader.readUint16()
}


func(self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.index)
}