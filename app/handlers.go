package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ret0rn/vtbMapAPI/internal/handlers"
	"github.com/ret0rn/vtbMapAPI/internal/service"
)

type PublicHandler struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

func (a *App) InitHandlers(srv *service.Service) {
	var api = handlers.NewPublicAPI(srv)
	a.WithPublicHandler(http.MethodGet, "/handling_list/by_client", api.GetHandlingListByClient)
}

func (a *App) WithPublicHandler(method, pattern string, handlerFunc gin.HandlerFunc) {
	a.customPublicHandler = append(a.customPublicHandler, PublicHandler{
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: handlerFunc,
	})
}
