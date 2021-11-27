package curl

import "C"
import "unsafe"

func goCallHeaderFunction(ptr *C.char, size C.size_t, ctx unsafe.Pointer) uintptr {
	curl := context_map.Get(uintptr(ctx))
	buf := C.GoBytes(unsafe.Pointer(ptr), C.int(size))
	if (*curl.headerFunction)(buf, curl.headerData) {
		return uintptr(size)
	}
	return C.CURL_WRITEFUNC_PAUSE
}