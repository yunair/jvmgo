package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"char":    "C",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"float":   "F",
	"double":  "D",
}


// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(name string) string {
	return "[" + toDescriptor(name)
}

// [XXX => [XXX
// int  => I
// XXX  => LXXX;
func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}

	if d, ok := primitiveTypes[className]; ok {
		return d
	}

	return "L" + className + ";"
}
