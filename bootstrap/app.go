package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"clms_website/config"
	"clms_website/routes"
	"github.com/joho/godotenv"
)

type Application struct {
	Config config.AppConfig
}

func NewApplication() *Application {
	godotenv.Load()
	cfg := config.Load()
	return &Application{Config: cfg}
}

func (a *Application) Serve() {
	routes.RegisterRoutes()
	routes.RegisterAPIRoutes()

	addr := ":" + a.Config.Port
	fmt.Printf("%s is running on http://localhost%s\n", a.Config.AppName, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
