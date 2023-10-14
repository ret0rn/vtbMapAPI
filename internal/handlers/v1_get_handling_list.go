package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ret0rn/vtbMapAPI/internal/model"
	"github.com/sirupsen/logrus"
)

type GetHandlingListByClientRequest struct {
	ClientType model.ClientType `form:"client_type" binding:"required,gt=0"`
}

type GetHandlingListByClientResponse struct {
	Data []*GetHandlingListByClientResponseData `json:"data"`
}

type GetHandlingListByClientResponseData struct {
	HandlingId       int64            `json:"handling_id"`
	Title            string           `json:"title"`
	ClientType       model.ClientType `json:"client_type"`
	HandlingDuration float64          `json:"handling_duration"` // в секундах
}

// GetHandlingListByClient     godoc
// @Summary                    Получение словаря услуг по типу клиента (физ или юр лицо)
// @Produce                    json
// @Param                      req query GetHandlingListByClientRequest true "Handling list by client request"
// @Success                    200 {object} GetHandlingListByClientResponse
// @Router                     /handling_list/by_client [get]
func (i *Implementation) GetHandlingListByClient(ctx *gin.Context) {
	var req GetHandlingListByClientRequest // получаем и валидируем запрос
	if err := ctx.BindQuery(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewError("Ошибка валидации %s", err.Error()))
		return
	}
	// проверяем что полученный тип клиента мы сможем обработать
	if !req.ClientType.IsWorkClientType() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewError("Недопустимый ClientType %d", req.ClientType))
		return
	}
	// получение данных из базы
	list, err := i.srv.GetHandlingListByFilter(ctx, &model.HandlingFilter{ClientType: req.ClientType})
	if err != nil {
		logrus.Error(ctx, "[GetHandlingListByClient] - cent's get handling list %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewError("Ошибка при получение списка услуг"))
		return
	}
	var resp = new(GetHandlingListByClientResponse)
	for _, hand := range list {
		resp.Data = append(resp.Data, &GetHandlingListByClientResponseData{
			HandlingId:       hand.HandlingId,
			Title:            hand.Title,
			ClientType:       hand.ClientType,
			HandlingDuration: hand.HandlingDuration.Seconds(),
		})
	}
	ctx.JSON(http.StatusOK, resp)
	return
}
