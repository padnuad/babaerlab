package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	userDelivery "baberlab/features/user/delivery"
	userRepo "baberlab/features/user/repository/postgres"
	userUsecase "baberlab/features/user/usecase"
)

var postgresDB *gorm.DB

func init() {
	var err error
	postgresDB, err = newPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	// if err := config.InitialDatabase(); err != nil {
	// 	panic(err)
	// }
}
func main() {

	e := echo.New()
	e.HideBanner = true
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")

	userDelivery.NewHandler(
		v1, userUsecase.NewUserUsecase(
			userRepo.NewUserRepository(postgresDB),
		),
	)

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "hello saharat"})
}

func newPostgresDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=localhost port=5433 user=postgres dbname=fillgoods-lab password=1234 sslmode=disable")

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	return db, err
}
