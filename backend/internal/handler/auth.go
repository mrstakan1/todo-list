package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"todolist/internal/model"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	r.POST("/register", func(c *gin.Context) {
		var in struct{ Username, Password string }
		if c.BindJSON(&in) != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		u := model.User{Username: in.Username, Password: string(hash)}
		if err := db.Create(&u).Error; err != nil {
			c.Status(http.StatusConflict)
			return
		}
		c.Status(http.StatusCreated)
	})

	r.POST("/login", func(c *gin.Context) {
		var in struct{ Username, Password string }
		if c.BindJSON(&in) != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		var u model.User
		if db.Where("username = ?", in.Username).First(&u).Error != nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password)) != nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid": u.ID,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
		})
		token, _ := tk.SignedString(secret)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})
}
