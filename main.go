package main

import (
	"os"

	_ "github.com/ret0rn/vtbMapAPI/docs"
	_ "go.uber.org/automaxprocs"
)

// @title          VTB_MAP_API
// @version        1.0
// @contact.name   Alex Romantsov
// @host           0.0.0.0:8070
// @BasePath       /api/v1
func main() {
	_ = os.Setenv("TZ", "Europe/Moscow")

}
