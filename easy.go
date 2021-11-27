package curl

import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

type CurlInfo C.CURLINFO
type CurlError C.CURLcode

type CurlString *C.char

func NewCurlString(s string) CurlString {
	return CurlString(unsafe.Pointer(C.CString(s)))
}

func FreCurlString(s CurlString) {
	C.free(unsafe.Pointer(s))
}

func (e CurlError) Error() string {
	ret := C.curl_easy_strerror(C.CURLcode(e))
	return fmt.Sprintf("curl: %s", C.GoString(ret))
}

type CURL struct {
	handle unsafe.Pointer
}

type contextMap struct {
	items map[uintptr]*CURL
	sync.RWMutex
}

func (c *contextMap) Set(k uintptr, v *CURL) {
	c.Lock()
	defer c.Unlock()


	c.items[k] = v
}

func (c *contextMap) Get(k uintptr) *CURL {
	c.RLock()
	defer c.RUnlock()

	return c.items[k]
}
