package model

import "time"

const (
	CacheHandlingKey = "cache-handling-%d-%d" // handling-clientType-handlingType
)

var DefaultTTl = 5 * time.Hour
