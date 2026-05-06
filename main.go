package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	adminKey := os.Getenv("ADMIN_KEY")
	if adminKey == "" {
		log.Fatal("ADMIN_KEY env required")
	}

	err := InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-type", "Authorization"},
		AllowCredentials: false,
	}))

	router.Static("/static", "./static")
	router.Static("/assets", "./static/dist/assets")
	router.GET("/health", HealthCheckHandler)
	router.GET("/crud", CrudPageHandler)

	router.NoRoute(func(c *gin.Context) {
		p := c.Request.URL.Path
		if strings.HasPrefix(p, "/cars") || strings.HasPrefix(p, "/admin") || strings.HasPrefix(p, "/health") {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.File("./static/dist/index.html")
	})

	vehicles := router.Group("/cars")
	{
		vehicles.GET("", GetAllVehiclesHandler)
		vehicles.GET("/:id", GetVehicleByIdHandler)
	}
	adminGroup := router.Group("/admin")
	adminGroup.Use(AuthMiddleware(adminKey))
	{
		adminGroup.POST("", CreateVehicleHandler)
		adminGroup.PUT("/:id", UpdateVehicleByIdHandler)
		adminGroup.DELETE("/:id", DeleteVehicleByIdHandler)
		adminGroup.POST("/:id/photos", UploadPhotosByIdHandler)
	}

	server := http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
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
