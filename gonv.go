package gonv

import (
	"reflect"
	"unsafe"
)

func Atob(s string) (b []byte) {
	if len(s) == 0 {
		return b
	}
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}))
}
