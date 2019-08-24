package utils

import (
	"unsafe"
)

func ByteToString(by []byte)  string{
	str:=*(*string)(unsafe.Pointer(&by))
	return str
}

func StringToByte(str string)  []byte{
	return *(*[]byte)(unsafe.Pointer(&str))
}
