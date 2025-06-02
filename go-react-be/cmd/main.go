package main

import (
	"bhachauk.github.io/go-react-sample/go-react-be/config"
	"bhachauk.github.io/go-react-sample/go-react-be/dao"
	"bhachauk.github.io/go-react-sample/go-react-be/models"
	"bhachauk.github.io/go-react-sample/go-react-be/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Connect to database
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}
	log.Println("Database migration completed!")

	// Initialize Gin router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Add any custom headers your frontend sends
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           86400, // How long the preflight request can be cached (24 hours)
	}))

	userDAO := dao.NewUserDAO(config.DB)
	routes.UserRoutes(router, userDAO)
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	err = router.Run(port)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
