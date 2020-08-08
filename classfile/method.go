package classfile

import "bufio"

/**
方法
*/
type Method struct {
	Access_flags     Method_access_flag_type
	Name_index       u2
	Descriptor_index u2
	Attributes_count u2
	Attribute_info   []IAttribute
}

type Method_access_flag_type u2

const (
	METHOD_ACC_PUBLIC       Method_access_flag_type = 0x0001
	METHOD_ACC_PRIVATE      Method_access_flag_type = 0x0002
	METHOD_ACC_PROTECTED    Method_access_flag_type = 0x0004
	METHOD_ACC_STATIC       Method_access_flag_type = 0x0008
	METHOD_ACC_FINAL        Method_access_flag_type = 0x0010
	METHOD_ACC_SYNCHRONIZED Method_access_flag_type = 0x0020
	//方法是不是由编辑器产生的桥接方法
	METHOD_ACC_BRIDGE Method_access_flag_type = 0x0040
	//方法是否接受不定参数
	METHOD_ACC_VARARGS  Method_access_flag_type = 0x0080
	METHOD_ACC_NATIVE   Method_access_flag_type = 0x0100
	METHOD_ACC_ABSTRACT Method_access_flag_type = 0x0400
	//方法是否为strictfp
	METHOD_ACC_STRICT Method_access_flag_type = 0x0800
	//字段是否由编辑器自动产生
	METHOD_ACC_SYNTHETIC Method_access_flag_type = 0x1000
)

/**
解析方法，包含方法个数和方法详情数组
*/
func ParseMethodInfo(cf *ClassFile, reader *bufio.Reader) {
	cf.MethodsCount = cf.readU2(reader)
	if cf.MethodsCount > 0 {
		cf.MethodInfo = make([]Method, cf.MethodsCount)
		var i u2 = 0
		for ; i < cf.MethodsCount; i++ {
			method := &(cf.MethodInfo[i])
			method.Access_flags = Method_access_flag_type(cf.readU2(reader))
			method.Name_index = cf.readU2(reader)
			method.Descriptor_index = cf.readU2(reader)
			method.Attributes_count, method.Attribute_info = ParseAttributeStatic(cf, reader)
		}
	}
}
