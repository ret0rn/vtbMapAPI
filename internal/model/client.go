package model

type ClientType int64

const (
	IndividualClientType ClientType = 1
	EnterpriseClientType ClientType = 2
)

var workClientType = []ClientType{IndividualClientType, EnterpriseClientType}

func (ct ClientType) IsWorkClientType() bool {
	for _, clientType := range workClientType {
		if clientType == ct {
			return true
		}
	}
	return false
}

func (ct ClientType) GetInt() int64 {
	return int64(ct)
}
