package errors

import "errors"

const (
	CacheNotFound    error = errors.New("Cache item not found")
	CacheUnavailable error = errors.New("Cache is unvailable")
)
