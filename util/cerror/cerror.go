package cerror

import "errors"

var (
	ErrGoMissing             = errors.New("go (golang) is missing")
	ErrGoVersionNotSupported = errors.New("go (golang) version is not supported")
)
