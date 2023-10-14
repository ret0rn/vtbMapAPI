package main

import (
	"context"
	"os"

	"github.com/ret0rn/vtbMapAPI/app"
	_ "github.com/ret0rn/vtbMapAPI/docs"
	_ "go.uber.org/automaxprocs"
)

// @title          VTB_MAP_API
// @version        0.1
// @contact.name   Alex Romantsov
// @host           0.0.0.0:8070
// @BasePath       /api/v1
func main() {
	_ = os.Setenv("TZ", "Europe/Moscow")
	var ctx = context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	err = a.Run()
	if err != nil {
		panic(err)
	}
}
