package heap

import "github.com/yunair/jvmgo/ch06/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (m *MemberRef) Descriptor() string {
	return m.descriptor
}

func (m *MemberRef) Name() string {
	return m.name
}

func newMethodRef(cp *ConstantPool, classInfo *classfile.ConstantMethodrefInfo) *MemberRef {
	ref := &MemberRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&classInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}
