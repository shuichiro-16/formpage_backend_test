package main

import (
	"formpage_backend_test/internal"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := internal.InitDB()
	if err != nil {
		log.Fatal("DB接続失敗: ", err)
	}
	defer db.Close()

	r := gin.Default()
	r.POST("/form", internal.PostFormHandler(db))

	if err := r.Run(":8080"); err != nil {
		log.Fatal("サーバ起動失敗: ", err)
	}
}
