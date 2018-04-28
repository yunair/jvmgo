package heap

import (
	"github.com/yunair/jvmgo/ch08/classpath"
	"fmt"
	"github.com/yunair/jvmgo/ch08/classfile"
	"github.com/yunair/jvmgo/ch08/msg"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/

type ClassLoader struct {
	cp          *classpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class //loader classes
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		verboseFlag: verboseFlag,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}

	if name[0] == '['{
		return self.loadArrayClass(name)
	}

	return self.loadNonArrayClass(name)
}
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	if self.verboseFlag{
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func link(class *Class) {
	verify(class)
	prepare(class)
}
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}

}
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.slotId
	if cpIndex > 0 {
		val := cp.GetConstant(cpIndex)

		switch field.descriptor {
		case "Z", "B", "C", "S", "I":
			vars.SetInt(slotId, val.(int32))
		case "J":
			vars.SetLong(slotId, val.(int64))
		case "F":
			vars.SetFloat(slotId, val.(float32))
		case "D":
			vars.SetDouble(slotId, val.(float64))
		case "Ljava/lang/String;":
			goStr := val.(string)
			jStr := JString(class.loader, goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}
func verify(class *Class) {

}

func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, //TODO
		name: name,
		loader: self,
		initStarted: true,
		superClass: self.LoadClass(msg.Object),
		interfaces: []*Class{
			self.LoadClass(msg.Cloneable),
			self.LoadClass(msg.Serializable),
		},
	}

	self.classMap[name] = class
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}

	return data, entry
}
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func resolveInterfaces(class *Class) {
	interfacesCount := len(class.interfaceNames)
	if interfacesCount > 0 {
		class.interfaces = make([]*Class, interfacesCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}

	return newClass(cf)
}
