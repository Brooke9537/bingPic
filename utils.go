package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetTodaypic() {
	log.Println("start get picurl")
	var b BingApi
	url := "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1"
	method := "Get"
	payload := strings.NewReader("")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &b)
	bingurl = b.Images[0].URL
	return
}
