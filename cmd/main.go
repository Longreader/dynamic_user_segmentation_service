package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"

	dynamicusersegmentation "github.com/Longreader/dynamic_user_segmentation_service.git"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/handlers"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/repository"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/service"
	"github.com/spf13/viper"
)

// @title           Swagger Avito Backend Junior API
// @version         1.0
// @description     This is a swagger docs for test API
// @contact.name   Alexey Kirichek
// @contact.url    https://vk.com/luxferoanimus
// @contact.email  rokirokz@mail.ru
// @host      localhost:8080
// @BasePath  /api/v1
// @query.collection.format multi
func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error at initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error at load environment passwords: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Error at database connection: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	h := handlers.NewHandler(services)

	srv := new(dynamicusersegmentation.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), h.InitRouter()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("AvitoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("AvitoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
