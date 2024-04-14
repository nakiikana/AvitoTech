package cache

import "errors"

var (
	ErrorPinging        = errors.New("couldn't ping")
	ErrorNoBannerFound  = errors.New("no required banner found")
	ErrorNilReturned    = errors.New("redis returned nil")
	ErrorMarshalProblem = errors.New("couldn't marshal the rawmessage before sending to redis")
	ErrorSetRedis       = errors.New("couldn't add the new value")
)
