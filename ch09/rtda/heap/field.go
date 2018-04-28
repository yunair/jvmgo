package heap

import "github.com/yunair/jvmgo/ch09/classfile"

type Field struct {
	ClassMember
	slotId uint
	constValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		field := &Field{}
		field.class = class
		field.copyMemberInfo(cfField)
		field.copyAttributes(cfField)
		fields[i] = field
	}
	return fields
}

func (self *Field) Class() *Class {
	return self.class
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) IsStatic() bool {
	return self.accessFlags & ACC_STATIC != 0
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
func (self *Field) copyAttributes(info *classfile.MemberInfo) {
	if valAttr := info.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}
func (self *Field) Descriptor() string {
	return self.descriptor
}
func (self *Field) SlotId() uint {
	return self.slotId
}
