package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social-todos-rest-api/module/item/transport"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to db: ", db)

	db = db.Debug()

	router := gin.Default()
	v1Routes := router.Group("/api/v1")
	{
		items := v1Routes.Group("items")
		{
			items.POST("", transport.CreateNewItem(db))
			items.GET("", transport.GetItems(db))
			items.GET("/:id", transport.GetItemById(db))
			items.PUT("/:id", transport.UpdateItemById(db))
			items.DELETE("/:id", transport.DeleteItemById(db))
		}
	}
	router.GET("/api/v1", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to golang rest api",
		})
	})

	if err2 := router.Run(":3000"); err2 != nil {
		log.Fatal(err2)
	}
}
