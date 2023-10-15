package main

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/ret0rn/vtbMapAPI/internal/config"
	"github.com/ret0rn/vtbMapAPI/internal/model"
	"github.com/ret0rn/vtbMapAPI/internal/service"
	"github.com/sirupsen/logrus"
)

type JsonFileData struct {
	SalePointName       string         `json:"salePointName"`
	Address             string         `json:"address"`
	HasRamp             string         `json:"hasRamp"`
	Latitude            float64        `json:"latitude"`
	Longitude           float64        `json:"longitude"`
	MetroStation        *string        `json:"metroStation"`
	OpenHours           []timeTableDay `json:"openHours"`
	OpenHoursIndividual []timeTableDay `json:"openHoursIndividual"`
}

type timeTableDay struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}

func main() {
	var ctx = context.Background()
	config.Load()
	srv, err := service.NewService(ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	jsonFile, err := os.Open("../../data/offices.json")
	if err != nil {
		logrus.Fatal(err)
	}
	defer jsonFile.Close()

	byteFile, err := io.ReadAll(jsonFile)
	if err != nil {
		logrus.Fatal(err)
	}

	var rawOffices []JsonFileData
	err = json.Unmarshal(byteFile, &rawOffices)

	clientTypeIndiv := model.IndividualClientType
	clientTypeEntr := model.EnterpriseClientType
	i := 0
	for _, office := range rawOffices {
		var clOffice = &model.Office{
			Longitude:  office.Longitude,
			Latitude:   office.Latitude,
			Address:    office.Address,
			OfficeName: office.SalePointName,
			HasRamp:    office.HasRamp == "Y",
		}
		if office.MetroStation != nil {
			clOffice.MetroStation = *office.MetroStation
		}
		var timeTableEnt model.OfficeTimeTable
		for _, ttble := range office.OpenHours {
			hours := strings.Split(ttble.Hours, "-")
			if len(hours) < 2 {
				continue
			}
			timeTableEnt.Days = append(timeTableEnt.Days, &model.DayTimeTable{
				Day:   ttble.Days,
				Start: hours[0],
				Stop:  hours[1],
			})
		}
		clOffice.TimetableEnterprise = &timeTableEnt

		var timeTableIndv model.OfficeTimeTable
		for _, ttble := range office.OpenHoursIndividual {
			hours := strings.Split(ttble.Hours, "-")
			if len(hours) < 2 {
				continue
			}
			timeTableIndv.Days = append(timeTableIndv.Days, &model.DayTimeTable{
				Day:   ttble.Days,
				Start: hours[0],
				Stop:  hours[1],
			})
		}
		clOffice.TimetableIndividual = &timeTableIndv

		i++
		if i%5 == 0 {
			clOffice.ClientTypes = []*model.ClientType{&clientTypeIndiv, &clientTypeEntr}
		} else {
			clOffice.ClientTypes = []*model.ClientType{&clientTypeIndiv}
		}
		handlList, err := srv.GetHandlingListByFilter(ctx, nil)
		if err != nil {
			logrus.Fatal(err)
		}

		clOffice.HandlingTypes = handlList.GetHandlingIds()
		err = srv.NewOffice(ctx, clOffice)
		if err != nil {
			logrus.Fatal(err)
		}
	}

}
