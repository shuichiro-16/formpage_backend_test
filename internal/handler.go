package internal

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFormHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form FormData
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		_, err := db.Exec("INSERT INTO forms (name, email, message) VALUES (?, ?, ?)", form.Name, form.Email, form.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
