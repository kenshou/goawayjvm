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
	ConstantInfo      []ConstantPool
	AccessFlags       u2
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
	ParseConstantPool(cf, reader)
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
