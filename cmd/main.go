package main

import (
	"./conf"
	"github.com/ashwanthkumar/slack-go-webhook"
	"log"
	"math"
	"reflect"
	"time"
	//"fmt"
)

func containStr(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func containInt(n []int, num int) bool {
	for _, v := range n {
		if v == num {
			return true
		}
	}
	return false
}

func PostSlack(msg string, url string) {
	field1 := slack.Field{Title: "ごみの日タイマー", Value: msg}
	attachment := slack.Attachment{}
	attachment.AddField(field1)
	color := "good"
	attachment.Color = &color
	payload := slack.Payload{
		Username:    "notify-gabage",
		Channel:     "to-do",
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(url, "", payload)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func main() {
	t := time.Now()
	wd := t.Weekday().String()
	dc := int(math.Ceil(float64(t.Day() / 7)))

	//Read WEBHOOKURL
	confJson, err := conf.ReadJson()
        if err != nil {
                log.Printf("error: %v\n", err)
        }
	webhookUrl := confJson.WEBHOOKURL

	//Read Config from config.toml
	config, err := conf.ReadConfig()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	cv := reflect.ValueOf(*config)
	ct := cv.Type()
	for i := 0; i < ct.NumField(); i++ {
		weeksIntSlice := cv.Field(i).Field(1).Interface().([]int)
		weekdayStrSlice := cv.Field(i).Field(0).Interface().([]string)
		if containStr(weekdayStrSlice, wd) && containInt(weeksIntSlice, dc) {
			message := cv.Field(i).Field(2).Interface().(string)
			PostSlack(message, webhookUrl)
		}
	}

}
