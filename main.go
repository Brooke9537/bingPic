package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type BingApi struct {
	Images []struct {
		Startdate     string        `json:"startdate"`
		Fullstartdate string        `json:"fullstartdate"`
		Enddate       string        `json:"enddate"`
		URL           string        `json:"url"`
		Urlbase       string        `json:"urlbase"`
		Copyright     string        `json:"copyright"`
		Copyrightlink string        `json:"copyrightlink"`
		Title         string        `json:"title"`
		Quiz          string        `json:"quiz"`
		Wp            bool          `json:"wp"`
		Hsh           string        `json:"hsh"`
		Drk           int           `json:"drk"`
		Top           int           `json:"top"`
		Bot           int           `json:"bot"`
		Hs            []interface{} `json:"hs"`
	} `json:"images"`
	Tooltips struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltips"`
}

var bingurl string

func main() {
	var err error
	route := gin.Default()
	gin.ErrorLogger()
	route.GET("/bing", bingpic)
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
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}
func bingpic(c *gin.Context) {
	var todaypic string
	todaypic = `https://s.cn.bing.net` + bingurl
	c.Redirect(302, todaypic)
}

func GetTodaypic() {
	log.Println("start get picurl")
	var b BingApi
	url := "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1"
	method := "Get"
	payload := strings.NewReader("")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	bingurl = b.Images[0].URL
	return
}
