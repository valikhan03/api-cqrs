package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"resource_admin/docs" // docs is generated by Swag CLI, you have to import it.
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"resource_admin/handlers"
	"resource_admin/repository"
	"resource_admin/usecase"
)

// @title 			Resource Admin Service API
// @version	 		1.0
// @Description 	This service is created for resource admins so they could create their content

// @host 		localhost:8077
// @BasePath	/api/v1/

func main() {
	docs.SwaggerInfo.Title = "Resource Admin Service API"
	docs.SwaggerInfo.Description = "This is an API for Resource Admin Service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/api/v1/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	
	router := gin.Default()



	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	repository := repository.NewRepository(InitDB())
	usecase := usecase.NewUseCase(repository, NewKafkaProducer())
	handers := handlers.NewHandler(usecase)

	api := router.Group("/api/v1/")
	{
		api.POST("/", handers.CreateResourceHandler)
		api.PUT("/", handers.UpdateResourceHandler)
		api.DELETE("/:id", handers.DeleteResourceHandler)
	}
	

	godotenv.Load()
	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() *gorm.DB {
	dbconn, err := ReadDBConfigs()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	db, err := gorm.Open(postgres.Open(dbconn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func ReadDBConfigs() (string, error) {

	viper.SetConfigFile("config/database.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	err = godotenv.Load()
	if err != nil {
		return "", err
	}

	dbconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		viper.Get("host"), viper.Get("user"), os.Getenv("DB_PASSWORD"), viper.Get("dbname"), viper.Get("port"), viper.Get("sslmode"))

	return dbconn, nil
}

func NewKafkaProducer() sarama.SyncProducer {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	addrs := []string{os.Getenv("KAFKA_CLUSTER_ADDR")}
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Timeout = 5 * time.Second
	prod, err := sarama.NewSyncProducer(addrs, configs)
	if err != nil {
		log.Fatal(err)
	}
	return prod
}
