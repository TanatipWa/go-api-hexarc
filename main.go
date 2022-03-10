package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TanatipWa/product-api/config"
	"github.com/TanatipWa/product-api/crud/entity"
	"github.com/TanatipWa/product-api/crud/repository"
	"github.com/TanatipWa/product-api/crud/usecase"
	"github.com/TanatipWa/product-api/db"
	"github.com/TanatipWa/product-api/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.SetConfigFile("config", "./config", "json")
}

func main() {
	envConfig := getConfig()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Mongo
	mongo, err := db.Connect(envConfig.Mongo)
	if err != nil {
		log.Println(err)
		return
	}

	crudRepo := repository.NewCrudRepository(mongo)
	crudUseCase := usecase.NewCrudUseCase(&envConfig, crudRepo)

	// Router
	router.NewRouter(e, crudUseCase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s%v", envConfig.Host, ":", envConfig.Port)))
}

func getConfig() entity.EnvConfig {

	return entity.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Mongo: db.MongoConfig{
			Timeout:  config.GetInt("database.mongodb.timeout"),
			MongoUri: config.GetString("database.mongodb.mongouri"),
			DBname:   config.GetString("database.mongodb.dbname"),
		},
	}
}
