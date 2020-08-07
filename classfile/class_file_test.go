package classfile

import (
	"bufio"
	"os"
	"testing"
)

func TestClassFile_Read(t *testing.T) {
	fi, err := os.Open("C:\\project\\luban\\jvm\\luban-jvm-research\\target\\classes\\com\\luban\\ziya\\jvm\\Test.class")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	cf := ClassFile{}
	cf.Read(r)
}
