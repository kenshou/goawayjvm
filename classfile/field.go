package classfile

import "bufio"

/**
属性
*/
type Field struct {
	Access_flags Field_access_flags_type
	Name_index   u2
	//字段类型描述符
	Descriptor_index u2
	Attributes_count u2
	Attribute_info   []IAttribute
}

//字段修饰符类型
type Field_access_flags_type u2

const (
	FIELD_ACC_PUBLIC    Field_access_flags_type = 0x0001
	FIELD_ACC_PRIVATE   Field_access_flags_type = 0x0002
	FIELD_ACC_PROTECTED Field_access_flags_type = 0x0004
	FIELD_ACC_STATIC    Field_access_flags_type = 0x0008
	FIELD_ACC_FINAL     Field_access_flags_type = 0x0010
	FIELD_ACC_VOLATILE  Field_access_flags_type = 0x0040
	FIELD_ACC_TRANSIENT Field_access_flags_type = 0x0080
	//字段是否由编辑器自动产生
	FIELD_ACC_SYNTHETIC Field_access_flags_type = 0x1000
	FIELD_ACC_ENUM      Field_access_flags_type = 0x4000
)

/**
解析字段，包含字段个数和字段详情数组
*/
func ParseFieldInfo(cf *ClassFile, reader *bufio.Reader) {
	cf.FieldCount = cf.readU2(reader)
	if cf.FieldCount > 0 {
		var i u2 = 0
		cf.FieldInfo = make([]Field, cf.FieldCount)
		for ; i < cf.FieldCount; i++ {
			field := &(cf.FieldInfo[i])
			field.Access_flags = Field_access_flags_type(cf.readU2(reader))
			field.Name_index = cf.readU2(reader)
			field.Descriptor_index = cf.readU2(reader)
			field.Attributes_count, field.Attribute_info = ParseAttributeStatic(cf, reader)
		}
	}
}
