package vbatd

import (
	"errors"
	"unsafe"
)

func getLong(vv VBAVariant) (ret interface{}, err error) {
	return int32(vv.getPtr()), nil
}

func getLongFromPtr(vv VBAVariant) (ret interface{}, err error) {
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	return *((*int32)(unsafe.Pointer(vv.getPtr()))), nil
}

func getLongArrFromPtr(vv VBAVariant) (ret interface{}, err error) {
	return getLongArr(vv.getPtr(), sizeVBADataType[vv.Flags[0]])
}

func getLongArrFromPtrPtr(vv VBAVariant) (ret interface{}, err error) {
	pv := *((*int32)(unsafe.Pointer(vv.getPtr())))

	return getLongArr(uintptr(pv), sizeVBADataType[vv.Flags[0]])
}

func getLongArr(pv uintptr, typeSize uint32) (ret interface{}, err error) {
	safearr := *((*SafeArray)(unsafe.Pointer(pv)))

	if safearr.cbElemets != typeSize {
		err = errors.New("SafeArray.cbElemets元素的字节大小与VBA Variant Type不一致.")
		return
	}
	if safearr.cDims == 1 {
		b := make([]int32, safearr.rgsabound[0].cElemets)
		for i := range b {
			b[i] = (*(*int32)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*4)))))
		}
		ret = b
	} else {
		err = errors.New("SafeArray.cDims != 1，未处理的情况.")
		return
	}

	return
}
