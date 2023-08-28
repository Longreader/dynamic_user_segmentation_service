package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/handlers"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/repository"
	"github.com/Longreader/dynamic_user_segmentation_service.git/service"
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

	log.SetFormatter(&log.JSONFormatter{})

	if err := initConfig(); err != nil {
		log.Fatalf("Error at initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error at load environment passwords: %s", err.Error())
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
		log.Fatalf("Error at database connection: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	h := handlers.NewHandler(services)

	r := h.InitRouter()

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(viper.GetString("port"), r))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
