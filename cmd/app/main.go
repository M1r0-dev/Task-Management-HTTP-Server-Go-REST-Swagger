// @title My API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /
package main

import (
	"flag"
	"http_server_arch/cmd/config"
	"http_server_arch/repository/ram_storage"
	"http_server_arch/usecases/service"
	"log"

	_ "http_server_arch/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"http_server_arch/api/http"
	pkgHttp "http_server_arch/pkg/http"
)

func main() {
	flag.Parse()
	AppFlags := config.ParseFlags()
	var cfg config.HTTPconfig
	config.MustLoad(AppFlags.ConfigPath, &cfg)

	objectRepo := ram_storage.NewObject()
	objectService := service.NewObject(objectRepo)
	objectHandlers := http.NewHandler(objectService)

	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	objectHandlers.WithObjectHandlers(r)

	log.Printf("Starting server on %s", *&cfg.Adress)
	if err := pkgHttp.CreateAndRunServer(r, *&cfg.Adress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
