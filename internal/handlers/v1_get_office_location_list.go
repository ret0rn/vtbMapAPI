package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ret0rn/vtbMapAPI/internal/model"
	"github.com/sirupsen/logrus"
)

type GetOfficeLocationListRequest struct {
	Longitude float64                             `json:"longitude" binding:"required"`
	Latitude  float64                             `json:"latitude" binding:"required"`
	Filter    *GetOfficeLocationListRequestFilter `json:"filter" binding:"omitempty,dive"`
}

type GetOfficeLocationListRequestFilter struct {
	HandlingType int64            `json:"handling_type" binding:"required,gt=0"`
	ClientType   model.ClientType `json:"client_type" binding:"required,gt=0"`
}

type GetOfficeLocationListResponse struct {
	Data []*GetOfficeLocationListResponseData `json:"data"`
}

type GetOfficeLocationListResponseData struct {
	OfficeID            int64                                             `json:"office_id"`
	Longitude           float64                                           `json:"longitude"`
	Latitude            float64                                           `json:"latitude"`
	Rate                float64                                           `json:"rate"`
	Distance            float64                                           `json:"distance"`
	WaitTime            float64                                           `json:"wait_time"`             // В секундах
	TravelDurationHuman float64                                           `json:"travel_duration_human"` // Время на маршрут пешком
	TravelDurationCar   float64                                           `json:"travel_duration_car"`   // Время на маршрут на машине
	CountPeople         int64                                             `json:"count_people"`          // Кол-во людей в очереди
	Address             string                                            `json:"address"`
	WorkloadKoef        float64                                           `json:"workload_koef"`
	OfficeName          string                                            `json:"officeName"`
	TimetableIndividual *GetOfficeLocationListResponseDataOfficeTimeTable `json:"timetable_individual"` // Расписание работы для Физ. лиц
	TimetableEnterprise *GetOfficeLocationListResponseDataOfficeTimeTable `json:"timetable_enterprise"` // Расписание работы для Юр. лиц
	MetroStation        string                                            `json:"metro_station"`
	HasRamp             bool                                              `json:"has_ramp"`
	ClientTypes         []*model.ClientType                               `json:"client_types"`   // Обслуживание Физ. лица или Юр. лица
	HandlingTypes       []int64                                           `json:"handling_types"` // Обслуживание Физ. лица или Юр. лица
	HandlingDuration    float64                                           `json:"handling_duration"`
}

type GetOfficeLocationListResponseDataOfficeTimeTable struct {
	Days []*GetOfficeLocationListResponseDataDayTimeTable `json:"days"`
}

type GetOfficeLocationListResponseDataDayTimeTable struct {
	Day   string `json:"day"`
	Start string `json:"start"`
	Stop  string `json:"stop"`
}

// GetOfficeLocationList       godoc
// @Summary                    Получение ближайших отделений с загруженностью
// @Produce                    json
// @Param                      req query GetOfficeLocationListRequest true "office location list request"
// @Success                    200 {object} GetOfficeLocationListResponse
// @Router                     /office/location [post]
func (i *Implementation) GetOfficeLocationList(ctx *gin.Context) {
	const (
		humanSpeed = 70  // м/м
		carSpeed   = 800 // м/м
	)
	var req GetOfficeLocationListRequest // получаем и валидируем запрос
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewError("Ошибка валидации %s", err.Error()))
		return
	}

	filter := &model.OfficeLocationFilter{
		Longitude: fmt.Sprintf("%f", req.Longitude),
		Latitude:  fmt.Sprintf("%f", req.Latitude),
	}
	if req.Filter != nil {
		filter.ClientType = req.Filter.ClientType
		filter.HandlingType = req.Filter.HandlingType
	}
	list, countPeople, ratesMap, handlingDuration, err := i.srv.GetOfficeLocationList(ctx, filter)
	if err != nil {
		logrus.Error(ctx, "[GetOfficeLocationList] - cent's get handling list %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewError("Ошибка при получение списка офисов"))
		return
	}
	var resp GetOfficeLocationListResponse
	for _, office := range list {
		countPpl := countPeople[office.OfficeID]
		resp.Data = append(resp.Data, &GetOfficeLocationListResponseData{
			OfficeID:            office.OfficeID,
			Longitude:           office.Longitude,
			Latitude:            office.Latitude,
			Rate:                ratesMap[office.OfficeID],
			Distance:            office.Distance,
			WaitTime:            handlingDuration.Seconds() * float64(countPpl),
			TravelDurationHuman: (office.Distance / humanSpeed) * 60,
			TravelDurationCar:   (office.Distance / carSpeed) * 60,
			CountPeople:         countPpl,
			Address:             office.Address,
			OfficeName:          office.OfficeName,
			TimetableIndividual: converTimeTable(office.TimetableIndividual),
			TimetableEnterprise: converTimeTable(office.TimetableEnterprise),
			MetroStation:        office.MetroStation,
			HasRamp:             office.HasRamp,
			ClientTypes:         office.ClientTypes,
			HandlingTypes:       office.HandlingTypes,
			WorkloadKoef:        handlingDuration.Hours() * float64(countPpl),
			HandlingDuration:    handlingDuration.Seconds(),
		})
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func converTimeTable(timeTables *model.OfficeTimeTable) *GetOfficeLocationListResponseDataOfficeTimeTable {
	var respTimeTable GetOfficeLocationListResponseDataOfficeTimeTable
	for _, tt := range timeTables.Days {
		respTimeTable.Days = append(respTimeTable.Days, &GetOfficeLocationListResponseDataDayTimeTable{
			Day:   tt.Day,
			Start: tt.Start.Format(time.RFC3339),
			Stop:  tt.Stop.Format(time.RFC3339),
		})
	}
	return &respTimeTable
}
