package main

import (
	"github.com/devzzk/memrizr/core"
	"github.com/devzzk/memrizr/global"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")

	core.InitConf()
	// 链接数据库
	global.DB = core.InitGorm()

	gin.SetMode(gin.ReleaseMode)

	//router := gin.Default() // warning: Creating an Engine instance with the Logger and Recovery middleware already attached.
	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	quit := make(chan os.Signal)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}
