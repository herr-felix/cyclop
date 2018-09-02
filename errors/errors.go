package errors

import "errors"

const (
	CacheNotFound    = errors.New("Cache item not found")
	CacheUnavailable = errors.New("Cache is unvailable")
)
