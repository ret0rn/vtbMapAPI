package model

import "time"

type HandlingList []*Handling

type Handling struct {
	HandlingId       int64
	Title            string
	ClientType       ClientType
	HandlingDuration *time.Duration
}

type HandlingFilter struct {
	HandlingIds []int64
	ClientType  ClientType
}
