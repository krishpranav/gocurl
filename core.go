package curl

import "C"
import (
	"time"
	"unsafe"
)

func GlobalInit(flags int) error {
	return newCurlError(C.curl_global_init(C.long(flags)))
}

func GlobalCleanup() {
	C.curl_global_cleanup()
}

type VersionInfoData struct {
	Age C.CURLVersion

	Version string
	VersionNum uint
	Host string
	Sslversion string
}

func Version() string {
	return C.GoString(C.curl_version())
}

func Getdate(date string) *time.Time {
	datestr := C.CString(data)
	defer C.free(unsafe.Pointer(datestr))
	t := C.curl_getdata(datestr, nil)
	if t == -1 {
		return nil
	}

	unix := time.Unix(int64(t), 0).UTC()
	return &unix
}