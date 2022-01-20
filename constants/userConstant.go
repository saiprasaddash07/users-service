package constants

import "time"

// request related
const (
	HTTP_METHOD_GET  = "GET"
	HTTP_METHOD_POST = "POST"
)

// api status
const (
	API_FAILED_STATUS  = "Fail"
	API_SUCCESS_STATUS = "Success"
)

// DB constants
const (
	DB_READER = "reader"
	DB_WRITER = "writer"
)

// caching constants durations are in nano sec
var (
	CACHE_TTL_VERY_SHORT time.Duration = 60 * 1_000_000_000
	CACHE_TTL_SHORT      time.Duration = 300 * 1_000_000_000
	CACHE_TTL_MEDIUM     time.Duration = 1_800 * 1_000_000_000
	CACHE_TTL_LONG       time.Duration = 3_600 * 1_000_000_000
	CACHE_TTL_VERY_LONG  time.Duration = 86_400 * 1_000_000_000
)

// info messages
const (
	INFO_CACHE_DISABLED = "cache disabled"
)
