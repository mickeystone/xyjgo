// dll for vba
package main

import "C"
import (
	"fmt"
	"unsafe"

	"github.com/xiongyejun/xyjgo/vbatd"
)

func main() {
}

//export RetVarPtr
func RetVarPtr() (VarPtr int32) {
	var str = "测试直接返回VBA String的VarPtr"

	VarPtr = int32(uintptr(str2VarPtr(str)))

	return

}

//export Sprintf
func Sprintf(pformat, pParamArray, nCount int32) (ptr unsafe.Pointer, lenth int) {
	var err error
	var str string

	var it interface{}
	if it, err = vbatd.Variant2interface(uintptr(pformat)); err != nil {
		str = err.Error()
		return str2ptr(str)
	}

	var strformat string
	var ok bool
	if strformat, ok = it.(string); !ok {
		return str2ptr("pformat 指向的不是VBA Variant的String.")
	}

	var its []interface{}
	if nCount > 0 {
		if its, err = vbatd.Variants2interfaces(uintptr(pParamArray), int(nCount)); err != nil {
			str = err.Error()
			return str2ptr(str)
		}
	}

	str = fmt.Sprintf(strformat, its...)

	return str2ptr(str)
}

//export Sum
func Sum(a, b int) int {
	return a + b
}
