package classfile

type Class_Access_flag_type u2

const (
	CLASS_ACC_PUBLIC    Class_Access_flag_type = 0x0001
	CLASS_ACC_FINAL     Class_Access_flag_type = 0x0010
	CLASS_ACC_SUPER     Class_Access_flag_type = 0x0020
	CLASS_ACC_INTERFACE Class_Access_flag_type = 0x0200
	CLASS_ACC_ABSTRACT  Class_Access_flag_type = 0x0400
	//标识这个类并非由用户代码产生的
	CLASS_ACC_SYNTHETIC Class_Access_flag_type = 0x1000
	//注解
	CLASS_ACC_ANNOTATION Class_Access_flag_type = 0x2000
	CLASS_ACC_ENUM       Class_Access_flag_type = 0x4000
	CLASS_ACC_MODULE     Class_Access_flag_type = 0x8000
)
