package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"todolist/internal/handler"
	"todolist/internal/middleware"
	"todolist/internal/model"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.User{}, &model.Todo{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://77.221.135.36:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))
	r.GET("/", func(c *gin.Context) { c.String(200, "API up") })

	handler.AuthRoutes(r, db)
	r.Use(middleware.Auth())
	handler.TodoRoutes(r, db)
	r.Run("77.221.135.36:8080")
}
