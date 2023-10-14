package model

import "time"

type Handling struct {
	HandlingId       int64
	Title            string
	ClientType       ClientType
	HandlingDuration *time.Duration
	OfficesIDs       []int64
}

type HandlingFilter struct {
	HandlingIds []int64
	ClientType  ClientType
}
