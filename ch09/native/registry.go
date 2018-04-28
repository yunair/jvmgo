package native

import "github.com/yunair/jvmgo/ch09/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func emptyNativeMethod(frame *rtda.Frame){
	// do nothing
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}

	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}

	return nil
}
