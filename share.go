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

func (shcurl *CURLSH) Cleanup() error {
	p := shcurl.handle
	return newCurlShareError(C.curl_share_cleanup(p))
}

func (shcurl *CURLSH) Setopt(opt int, param interface{}) error {
	p := shcurl.handle
	if param == nil {
		return newCurlShareError(C.curl_share_setopt_pointer(p, C.CURLSHoption(opt), nil))
	}
	switch opt {

	case SHOPT_SHARE, SHOPT_UNSHARE:
		if val, ok := param.(int); ok {
			return newCurlShareError(C.curl_share_setopt_long(p, C.CURLSHoption(opt), C.long(val)))
		}
	}
	panic("not supported CURLSH.Setopt opt or param")
	return nil
}
