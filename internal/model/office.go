package model

type Office struct {
	OfficeID            int64
	Longitude           float64
	Latitude            float64
	Address             string
	OfficeName          string
	IsActive            bool
	TimetableIndividual *OfficeTimeTable
	TimetableEnterprise *OfficeTimeTable
	ClientTypes         []*ClientType
	HandlingTypes       []int64
	MetroStation        string
	HasRamp             bool
}

type OfficeLocationList []*OfficeLocation

type OfficeLocation struct {
	OfficeID            int64
	Longitude           float64
	Latitude            float64
	Distance            float64
	Address             string
	OfficeName          string
	TimetableIndividual *OfficeTimeTable
	TimetableEnterprise *OfficeTimeTable
	MetroStation        string
	HasRamp             bool
	ClientTypes         []*ClientType
	HandlingTypes       []int64
}

type OfficeTimeTable struct {
	Days []*DayTimeTable `json:"days"`
}

type DayTimeTable struct {
	Day   string `json:"day"`
	Start string `json:"start"`
	Stop  string `json:"stop"`
}

type OfficeLocationFilter struct {
	Longitude    string
	Latitude     string
	HandlingType int64
	ClientType   ClientType
}

func (oll OfficeLocationList) GetOfficeIds() (offices []int64) {
	offices = make([]int64, len(oll))
	for _, office := range oll {
		offices = append(offices, office.OfficeID)
	}
	return
}
