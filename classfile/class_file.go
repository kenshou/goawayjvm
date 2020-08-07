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
type ClassFile struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantInfo      []ConstantPool
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	InterfacesCount   uint16
	Interfaces        []InterfaceInfo
	FieldCount        uint16
	FieldInfo         []Field
	MethodsCount      uint16
	MethodInfo        []Method
	AttributesCount   uint16
	AttributeInfo     []Attribute
}

func (cf *ClassFile) Read(reader *bufio.Reader) {
	cf.Magic = readUint32(reader)
	if cf.Magic != 0xCAFEBABE {
		panic("class文件格式不对，魔数不能对应")
	}
	cf.MinorVersion = readUint16(reader)
	cf.MajorVersion = readUint16(reader)
	cf.ConstantPoolCount = readUint16(reader)
}

func readUint32(reader *bufio.Reader) uint32 {
	arr := make([]byte, 4)
	reader.Read(arr)
	return binary.BigEndian.Uint32(arr)
}

func readUint16(reader *bufio.Reader) uint16 {
	arr := make([]byte, 2)
	reader.Read(arr)
	return binary.BigEndian.Uint16(arr)
}
