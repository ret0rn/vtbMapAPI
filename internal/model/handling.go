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
	HandlingId int64
	ClientType ClientType
}

func (hl HandlingList) GetHandlingIds() []int64 {
	var handlingIds = make([]int64, len(hl))
	for _, h := range hl {
		handlingIds = append(handlingIds, h.HandlingId)
	}
	return handlingIds
}
