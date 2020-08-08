package classfile

import (
	"bufio"
	"encoding/binary"
)

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type (
	u1 uint8
	u2 uint16
	u4 uint32
	u8 uint64
)

type ClassFile struct {
	Magic             u4
	MinorVersion      u2
	MajorVersion      u2
	ConstantPoolCount u2
	ConstantInfo      []IConstanPool
	AccessFlags       Class_Access_flag_type
	ThisClass         u2
	SuperClass        u2
	InterfacesCount   u2
	Interfaces        []InterfaceInfo
	FieldCount        u2
	FieldInfo         []Field
	MethodsCount      u2
	MethodInfo        []Method
	AttributesCount   u2
	AttributeInfo     []Attribute
}

func (cf *ClassFile) Read(reader *bufio.Reader) {
	cf.Magic = cf.readU4(reader)
	if cf.Magic != 0xCAFEBABE {
		panic("class文件格式不对，魔数不能对应")
	}
	cf.MinorVersion = cf.readU2(reader)
	cf.MajorVersion = cf.readU2(reader)
	//cf.ConstantPoolCount = readU2(reader)
	//常量池解析
	ParseConstantPool(cf, reader)
	cf.AccessFlags = Class_Access_flag_type(cf.readU2(reader))
	cf.ThisClass = cf.readU2(reader)
	cf.SuperClass = cf.readU2(reader)
	//接口索引集合
	ParseInterfaceInfo(cf, reader)
	//解析类or实列字段
	ParseFieldInfo(cf, reader)
	//解析方法
	ParseMethodInfo(cf, reader)
	cf.AttributesCount, cf.AttributeInfo = ParseAttributeStatic(cf, reader)

}

func (cf *ClassFile) readU4(reader *bufio.Reader) u4 {
	arr := make([]byte, 4)
	reader.Read(arr)
	return u4(binary.BigEndian.Uint32(arr))
}

func (cf *ClassFile) readU2(reader *bufio.Reader) u2 {
	arr := make([]byte, 2)
	reader.Read(arr)
	return u2(binary.BigEndian.Uint16(arr))
}

func (cf *ClassFile) readU1(reader *bufio.Reader) u1 {
	readByte, _ := reader.ReadByte()
	return u1(readByte)
}

func (cf *ClassFile) readU1Array(reader *bufio.Reader, length uint32) []u1 {
	arr := make([]byte, length)
	reader.Read(arr)
	res := make([]u1, length)
	for i, v := range arr {
		res[i] = u1(v)
	}
	return res
}
func (cf *ClassFile) readU2Array(reader *bufio.Reader, length uint16) []u2 {
	res := make([]u2, length)
	for i := 0; i < len(res); i++ {
		res[i] = cf.readU2(reader)
	}
	return res
}

func (cf *ClassFile) readU8(reader *bufio.Reader) u8 {
	arr := make([]byte, 8)
	reader.Read(arr)
	return u8(binary.BigEndian.Uint64(arr))
}
