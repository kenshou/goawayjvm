package classfile

import "bufio"

/**
java接口
*/
type InterfaceInfo u2

/**
java接口的解析，包含个数，和数组
*/
func ParseInterfaceInfo(cf *ClassFile, reader *bufio.Reader) {
	interfaceLength := cf.readU2(reader)
	if interfaceLength > 0 {
		cf.Interfaces = make([]InterfaceInfo, interfaceLength)
		var i u2 = 0
		for ; i < interfaceLength; i++ {
			cf.Interfaces[i] = InterfaceInfo(cf.readU2(reader))
		}
	}
}
