package curl

const (
	_DEBUG = 10 + (iota + 1)
	_INFO
	_WARN
	_ERROR
)

const _DEFAULT_LOG_LEVEL = _WARN

var log_level = _DEFAULT_LOG_LEVEL

func SetLogLevel(levelName string) {
	switch levelName {
	case "DEBUG":
		log_level = _DEBUG
	case "INFO":
		log_level = _INFO
	case "WARN":
		log_level = _WARN
	case "ERROR":
		log_level = _ERROR
	case "DEFAULT_LOG_LEVEL":
		log_level = _DEFAULT_LOG_LEVEL
	}
}
