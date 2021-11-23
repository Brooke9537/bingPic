package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var bingurl string

func main() {
	route := gin.Default()
	gin.ErrorLogger()
	route.GET("/bing", func(c *gin.Context) {
		c.Redirect(302, `https://s.cn.bing.net`+bingurl)
	})
	srv := &http.Server{
		Addr:    "127.0.0.1:1000",
		Handler: route,
	}

	go func() {
		GetTodaypic()
		var spec = fmt.Sprintf("00 00 */1 * * ?")
		crontab := cron.New(cron.WithSeconds())
		crontab.AddFunc(spec, GetTodaypic)
		crontab.Start()
		log.Println("start cron")
	}()

	log.Println("start server")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}
