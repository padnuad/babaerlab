package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"baberlab/domain"
	userDelivery "baberlab/features/user/delivery"
	userRepo "baberlab/features/user/repository/postgres"
	userUsecase "baberlab/features/user/usecase"
)

var postgresDB *gorm.DB

func init() {
	// os.Setenv("RUN_ENV", "prod")
	// runEnv = os.Getenv("RUN_ENV")
	// if runEnv == "" {
	// 	runEnv = "dev"
	// 	firebaseProjectName = "driveman-authentication"
	// } else if runEnv == "prod" {
	// 	firebaseProjectName = "driveman-authentication-prod"
	// }

	// log.Println("Running in env :: ", runEnv)

	var err error
	postgresDB, err = newPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	autoMigrate()
}
func main() {

	e := echo.New()
	e.HideBanner = true
	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

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
	connectionString := fmt.Sprintf("host=localhost user=postgres password=1234 dbname=fillgoods-lab port=5433 sslmode=disable TimeZone=Asia/Shanghai")
	// if runEnv == "prod" {
	// 	connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 		utils.Config.PostgresDatabaseProd.Server,
	// 		utils.Config.PostgresDatabaseProd.User,
	// 		utils.Config.PostgresDatabaseProd.Password,
	// 		utils.Config.PostgresDatabaseProd.Database,
	// 		utils.Config.PostgresDatabaseProd.Port,
	// 	)
	// }

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	return db, err
}

func autoMigrate() {
	err := postgresDB.AutoMigrate(
		&domain.User{},
	)

	if err != nil {
		panic(err)
	}
}
