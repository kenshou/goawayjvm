package classfile

import "bufio"

type Attribute struct {
	Attribute_name_index u2
	Attribute_length     u4
	Info                 []u1
}

/**
解析属性（用于字段和方法上的解析），不会更新ClassFile中的值
*/
func ParseAttributeStatic(cf *ClassFile, reader *bufio.Reader) (u2, []Attribute) {
	attributeLength := cf.readU2(reader)
	var attrArray []Attribute = nil
	if attributeLength > 0 {
		attrArray = make([]Attribute, attributeLength)
		for i := 0; i < len(attrArray); i++ {
			nameIndex := cf.readU2(reader)
			infoLength := cf.readU4(reader)
			attrArray[i] = Attribute{
				Attribute_name_index: nameIndex,
				Attribute_length:     infoLength,
				Info:                 cf.readU1Array(reader, uint32(infoLength)),
			}
		}
	}
	return attributeLength, attrArray
}
