package heap

import (
	"strings"
	"github.com/yunair/jvmgo/ch09/classfile"
	"github.com/yunair/jvmgo/ch09/msg"
)

type Class struct {
	accessFlags       uint16
	name              string // this class name
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object // java.lang.Class instance
}

func (c *Class) JClass() *Object {
	return c.jClass
}

func (c *Class) Loader() *ClassLoader {
	return c.loader
}

func (c *Class) InitStarted() bool {
	return c.initStarted
}

func (c *Class) StartInit() {
	c.initStarted = true
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) SuperClass() *Class {
	return c.superClass
}

func (c *Class) Methods() []*Method {
	return c.methods
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) IsPublic() bool {
	return self.accessFlags&ACC_PUBLIC != 0
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}
func (self *Class) NewObject() *Object {
	return newObject(self)
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, m := range self.Methods() {
		if m.Name() == name && m.Descriptor() == descriptor {
			return m
		}
	}
	return nil
}
func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}
func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}

func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}
func toClassName(descriptor string) string {
	if descriptor[0] == '[' { // array
		return descriptor
	}
	if descriptor[0] == 'L' { // object
		// remove start 'L' and end ';'
		return descriptor[1:len(descriptor)-1]
	}

	for className, d := range primitiveTypes {
		if d == descriptor { // primitiveTypes
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}

func (self *Class) isJlObject() bool {
	return self.name == msg.Object
}
func (self *Class) isJlCloneable() bool {
	return self.name == msg.Cloneable
}
func (self *Class) isJioSerializable() bool {
	return self.name == msg.Serializable
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}
func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}
func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}
