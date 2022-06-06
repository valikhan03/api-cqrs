package main

import(
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"user_service/handler"
	"user_service/usecase"
	"user_service/repository"
)

// @Title User Service
// @Version 1.0
// @Description

// @host localhost:8085
// @BasePath /api/v1
func main() {
	db, err := repository.InitDB()
	if err != nil{
		log.Println(err)
	}
	repository := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repository)
	handler := handler.NewHandler(usecase)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.LogIn)
	}

	server := &http.Server{
		Addr: "localhost:8085",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}