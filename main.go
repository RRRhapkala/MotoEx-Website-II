package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	err := InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/health", HealthCheckHandler)
	router.GET("/", MainPageHandler)
	router.GET("/catalog", CatalogPageHandler)
	router.GET("/about/:id", AboutPageHandler)
	router.GET("/reviews", ReviewsPageHandler)
	router.GET("/crud", CrudPageHandler)

	for _, lang := range []string{"ru", "en"} {
		router.GET("/"+lang+"/", MainPageHandler)
		router.GET("/"+lang+"/catalog", CatalogPageHandler)
		router.GET("/"+lang+"/about/:id", AboutPageHandler)
		router.GET("/"+lang+"/reviews", ReviewsPageHandler)
		router.GET("/"+lang+"/crud", CrudPageHandler)

	}

	vehicles := router.Group("/cars")
	{
		vehicles.GET("", GetAllVehiclesHandler)
		vehicles.GET("/:id", GetVehicleByIdHandler)
	}
	adminGroup := router.Group("/admin")
	adminGroup.Use(AuthMiddleware())
	{
		adminGroup.POST("", CreateVehicleHandler)
		adminGroup.PUT("/:id", UpdateVehicleByIdHandler)
		adminGroup.DELETE("/:id", DeleteVehicleByIdHandler)
		adminGroup.POST("/:id/photos", UploadPhotosByIdHandler)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to run server", err)
		}
	}()
	<-quit

	fmt.Println("\nShutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to shut down server", err)
	}

	fmt.Println("Server shut down")
}
