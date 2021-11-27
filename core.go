package curl

import "C"

func GlobalInit(flags int) error {
	return newCurlError(C.curl_global_init(c.long(flags)))
}
