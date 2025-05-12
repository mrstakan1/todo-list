package handler

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todolist/internal/model"
)

func TodoRoutes(r *gin.Engine, db *gorm.DB) {
	g := r.Group("/todos")
	g.GET("", func(c *gin.Context) {
		var list []model.Todo
		db.Where("user_id = ?", c.GetUint("uid")).Find(&list)
		c.JSON(http.StatusOK, list)
	})
	g.POST("", func(c *gin.Context) {
		var t model.Todo
		if c.BindJSON(&t) != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		t.UserID = c.GetUint("uid")
		db.Create(&t)
		c.JSON(http.StatusCreated, t)
	})
	g.PUT("/:id", func(c *gin.Context) {
		var t model.Todo
		if db.Where("id = ? AND user_id = ?", c.Param("id"), c.GetUint("uid")).First(&t).Error != nil {
			c.Status(http.StatusNotFound)
			return
		}
		var in model.Todo
		if c.BindJSON(&in) != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		t.Title, t.Completed = in.Title, in.Completed
		db.Save(&t)
		c.JSON(http.StatusOK, t)
	})
	g.DELETE("/:id", func(c *gin.Context) {
		db.Where("id = ? AND user_id = ?", c.Param("id"), c.GetUint("uid")).Delete(&model.Todo{})
		c.Status(http.StatusNoContent)
	})

	g.GET("/export", func(c *gin.Context) {
		fmt := c.DefaultQuery("fmt", "json")
		var list []model.Todo
		db.Where("user_id = ?", c.GetUint("uid")).Find(&list)

		switch fmt {
		case "csv":
			c.Header("Content-Disposition", "attachment; filename=todos.csv")
			c.Header("Content-Type", "text/csv; charset=utf-8")
			w := csv.NewWriter(c.Writer)
			_ = w.Write([]string{"id", "title", "completed"})
			for _, t := range list {
				_ = w.Write([]string{
					strconv.FormatUint(uint64(t.ID), 10),
					t.Title,
					strconv.FormatBool(t.Completed),
				})
			}
			w.Flush()
		default: // json
			c.Header("Content-Disposition", "attachment; filename=todos.json")
			c.JSON(http.StatusOK, list)
		}
	})
}
