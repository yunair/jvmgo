package heap

import "github.com/yunair/jvmgo/ch08/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) MaxLocals() uint {
	return m.maxLocals
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}
func (self *Method) Name() string {
	return self.name
}

func (self *Method) Class() *Class {
	return self.class
}
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}
func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D"{
			self.argSlotCount++
		}
	}
	// 加入隐含的this
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		method := &Method{}
		method.class = class
		method.copyMemberInfo(cfMethod)
		method.copyAttributes(cfMethod)
		method.calcArgSlotCount()
		methods[i] = method
	}
	return methods
}
