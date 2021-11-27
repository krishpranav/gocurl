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
	Age C.CURLversion
	Version       string
	VersionNum    uint
	Host          string
	Features      int
	SslVersion    string
	SslVersionNum int
	LibzVersion   string
	Protocols     []string
	Ares    string
	AresNum int
	Libidn string
	IconvVerNum   int
	LibsshVersion string
}

func Version() string {
	return C.GoString(C.curl_version())
}

func VersionInfo(ver C.CURLversion) *VersionInfoData {
	data := C.curl_version_info(ver)
	ret := new(VersionInfoData)
	ret.Age = data.age
	switch age := ret.Age; {
	case age >= 0:
		ret.Version = string(C.GoString(data.version))
		ret.VersionNum = uint(data.version_num)
		ret.Host = C.GoString(data.host)
		ret.Features = int(data.features)
		ret.SslVersion = C.GoString(data.ssl_version)
		ret.SslVersionNum = int(data.ssl_version_num)
		ret.LibzVersion = C.GoString(data.libz_version)
		ret.Protocols = []string{}
		for i := C.int(0); C.string_array_index(data.protocols, i) != nil; i++ {
			p := C.string_array_index(data.protocols, i)
			ret.Protocols = append(ret.Protocols, C.GoString(p))
		}
		fallthrough
	case age >= 1:
		ret.Ares = C.GoString(data.ares)
		ret.AresNum = int(data.ares_num)
		fallthrough
	case age >= 2:
		ret.Libidn = C.GoString(data.libidn)
		fallthrough
	case age >= 3:
		ret.IconvVerNum = int(data.iconv_ver_num)
		ret.LibsshVersion = C.GoString(data.libssh_version)
	}
	return ret
}


func Getdate(date string) *time.Time {
	datestr := C.CString(date)
	defer C.free(unsafe.Pointer(datestr))
	t := C.curl_getdate(datestr, nil)
	if t == -1 {
		return nil
	}
	unix := time.Unix(int64(t), 0).UTC()
	return &unix

}