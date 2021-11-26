package curl

import "C"
import "unsafe"

type CurlMultiError C.CURLMcode
type CurlMultiMsg C.CURLMSG

func (e CurlMultiError) Error() string {
	ret := C.curl_multi_strerror(C.CURLMcode(e))
	return c.GoString(ret)
}

func newCurlMultiError(errno C.CURLMcode) error {
	if errno == 0 {
		return nil
	}

	return CurlMultiError(errno)
}

type CURLM struct {
	handle unsafe.Pointer
}
