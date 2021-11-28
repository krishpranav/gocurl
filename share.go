package curl

import "C"
import "unsafe"

type CurlShareError C.CURLMcode

func (e CurlShareError) Error() string {
	ret := C.curl_share_strerror(C.CURLSHcode(e))
	return C.GoString(ret)
}

func newCurlShareError(errno C.CURLSHcode) error {
	if errno == 0 {
		return nil
	}

	return CurlShareError(errno)
}

type CURLSH struct {
	handle unsafe.Pointer
}

func ShareInit() *CURLSH {
	p := C.curl_share_init()
	return &CURLSH{p}
}
