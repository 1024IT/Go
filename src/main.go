package main

import (
	//configをインポート
	"go_sample/config"
	// ルーターをインポート
	"go_sample/router"
	// Ginをインポート
	"github.com/gin-gonic/gin"
	//DBをインポート
	"go_sample/database"
)

func main() {
	//configの初期化
	config.Init()

	// DBの初期化
	database.Init()
	defer database.Close()

	//routerの初期化
	router.Init()
	router := gin.Default()

	// JSON形式で「"message": "Hello World"」を出力するAPI
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!",
		})
	})

	router.Run(":3001")
}
