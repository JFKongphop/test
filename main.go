package main

import (
	"context"
	"fmt"
	
	"net/http"
	"os"
	"time"

	"github.com/carlmjohnson/requests"
	// "github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	UserID       string    `gorm:"column:userId;primaryKey;type:varchar(50);index" json:"userId"`
	FirstName    string    `gorm:"column:firstname;type:varchar(50)" json:"firstname"`
	LastName     string    `gorm:"column:lastname;type:varchar(50)" json:"lastname"`
	Email        string    `gorm:"column:email;type:varchar(50);unique" json:"email"`
	Username     string    `gorm:"column:username;type:varchar(50)" json:"username"`
	RegisterDate time.Time `gorm:"column:registerDate;autoCreateTime" json:"registerDate"`
}

var db *gorm.DB
var err error

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"word": os.Getenv("WORD"),
		})
	})

	e.GET("/wwww", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"aaaaa": os.Getenv("WORD"),
		})
	})

	e.GET("/post/:id", func(c echo.Context) error {
		id := c.Param("id")

		dataSourceName := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("RW_USERNAME"),
			os.Getenv("RW_PASSWORD"),
			os.Getenv("RW_HOST"),
			os.Getenv("RW_PORT"),
			os.Getenv("RW_DATABASE"),
		)

		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)
    var response map[string]interface{}
    err := requests.
			URL(url).
			ToJSON(&response).
			Fetch(context.Background())
		if err != nil {
			return c.JSON(echo.ErrBadRequest.Code, err.Error())
		}

		response["data"] = dataSourceName

		return c.JSON(http.StatusOK, response)
	})

	e.GET("/test/:userId", func(c echo.Context) error {
		userId := c.Param("userId")
		fmt.Println(userId)

		var user User
		if err := db.Select("username", "firstname", "lastname").Where("userId = ?", userId).First(&user).Error; err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	})

	e.Start(":1111")
}

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	dataSourceName := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("RW_USERNAME"),
		os.Getenv("RW_PASSWORD"),
		os.Getenv("RW_HOST"),
		os.Getenv("RW_PORT"),
		os.Getenv("RW_DATABASE"),
	)

	dial := mysql.Open(dataSourceName)
	db, err = gorm.Open(dial, &gorm.Config{
		DryRun: false,
	})

	if err != nil {
		log.Fatal("connect error", err)
	}
}