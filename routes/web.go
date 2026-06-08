package routes

import (
	"net/http"

	"github.com/mechneerd/gow/routing"
	"github.com/mechneerd/gow/view"
)

var router *routing.Router
var viewEngine *view.Engine

func init() {
	router = routing.NewRouter()
	viewEngine = view.NewEngine([]string{"resources/views"}, "storage/framework/views")

	router.Get("/", func(w http.ResponseWriter, r *http.Request) error {
		html, err := viewEngine.Make("welcome", map[string]any{
			"AppName": "clms_website",
			"Year":    "2026",
		})
		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1>Welcome to CLMS</h1>"))
			return nil
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return nil
	})
}

func RegisterRoutes() {
	http.Handle("/", router)
}
