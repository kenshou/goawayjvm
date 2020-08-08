package classfile

import "bufio"

/**
常量池
*/
type ConstantPool struct {
	Tag ConstantType `常量次的项目类型标识`
}
type CONSTANT_Utf8_info struct {
	ConstantPool
	Length u2
	Bytes  []u1
}
type CONSTANT_Integer_info struct {
	ConstantPool
	Bytes u4
}
type CONSTANT_Float_info struct {
	ConstantPool
	Bytes u4
}
type CONSTANT_Long_info struct {
	ConstantPool
	Bytes u8
}
type CONSTANT_Double_info struct {
	ConstantPool
	Bytes u8
}
type CONSTANT_Class_info struct {
	ConstantPool
	Index u2
}
type CONSTANT_String_info struct {
	ConstantPool
	Index u2
}
type CONSTANT_Fieldref_info struct {
	ConstantPool
	Index4ClassInfo   u2 //指向CONSTANT_Class_info的索引
	Index4NameAndType u2 //指向CONSTANT_NameAndType的索引
}
type CONSTANT_Methodref_info struct {
	ConstantPool
	Index4ClassInfo   u2 //指向CONSTANT_Class_info的索引
	Index4NameAndType u2 //指向CONSTANT_NameAndType的索引
}
type CONSTANT_InterfaceMethodref_info struct {
	ConstantPool
	Index4ClassInfo   u2 //指向CONSTANT_Class_info的索引
	Index4NameAndType u2 //指向CONSTANT_NameAndType的索引
}
type CONSTANT_NameAndType_info struct {
	ConstantPool
	Index4Name u2 //指向该字段名称常量项的索引
	Index4Des  u2 //指向该字段或方法描述常量项的索引
}
type CONSTANT_MethodHandle_info struct {
	ConstantPool
	ReferenceKind  u1 //值必须在1到9之间
	ReferenceIndex u2 //值必须是对常量池的有效索引
}
type CONSTANT_MethodType_info struct {
	ConstantPool
	DescriptorIndex u2
}
type CONSTANT_Dynamic_info struct {
	ConstantPool
	BootstrapMethodAttrIndex u2
	NameAndTypeIndex         u2
}
type CONSTANT_InvokeDynamic_info struct {
	ConstantPool
	BootstrapMethodAttrIndex u2
	NameAndTypeIndex         u2
}
type CONSTANT_Module_info struct {
	ConstantPool
	NameIndex u2
}
type CONSTAN_Package_info struct {
	ConstantPool
	NameIndex u2
}

/**
常量池的项目类型
*/
type ConstantType u1

const (
	CONSTANT_Utf8_info_type               ConstantType = 1  //utf-8编码的字符串
	CONSTANT_Integer_info_type            ConstantType = 3  //整型字面量
	CONSTANT_Float_info_type              ConstantType = 4  //浮点型字面量
	CONSTANT_Long_info_type               ConstantType = 5  //长整型字面量
	CONSTANT_Double_info_type             ConstantType = 6  //双精度型字面量
	CONSTANT_Class_info_type              ConstantType = 7  //类或者接口的符号引用
	CONSTANT_String_info_type             ConstantType = 8  //字符串类型字面量
	CONSTANT_Fieldref_info_type           ConstantType = 9  //字段的符号引用
	CONSTANT_Methodref_info_type          ConstantType = 10 //类中方法的符号引用
	CONSTANT_InterfaceMethodref_info_type ConstantType = 11 //接口中方法的符号引用
	CONSTANT_NameAndType_info_type        ConstantType = 12 //字段或方法的部分符号引用
	CONSTANT_MethodHandle_info_type       ConstantType = 15 //表示方法句柄
	CONSTANT_MethodType_info_type         ConstantType = 16 //表示方法类型
	CONSTANT_Dynamic_info_type            ConstantType = 17 //表示一个动态计算常量
	CONSTANT_InvokeDynamic_info_type      ConstantType = 18 //表示一个动态方法调用点
	CONSTANT_Module_info_type             ConstantType = 19 //表示一个模块
	CONSTAN_Package_info_type             ConstantType = 20 //表示一个模块中开放或者导出的包
)

func ParseConstantPool(cf *ClassFile, reader *bufio.Reader) {
	cf.ConstantPoolCount = cf.readU2(reader)

}
