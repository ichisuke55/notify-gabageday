package main

import (
	"flag"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/ichisuke55/notify-gabageday/conf"
	"log"
	"math"
	"reflect"
	"strconv"
	"time"
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
	// command line arguments
	var (
		argDay  = flag.String("d", "today", "no args execute today information. you can select today(default) or tomorrow")
		testDay = flag.Int("n", 0, "to check specific day at this month")
	)
	// *argDay is contain today or tomorrow
	// parse args
	flag.Parse()

	t := time.Now()
	wd := t.Weekday().String()
	dc := 0
	if t.Day() == 1 {
		dc = 1
	} else {
		dc = int(math.Floor(float64((t.Day()-1)/7))) + 1
	}

	message := "<!channel> "

	switch {
	case *argDay == "today":
		message += "今日は"
		fmt.Println("today!")
	case *argDay == "tomorrow":
		message += "明日は"
		t = time.Now().Add(time.Duration(24) * time.Hour)
		wd = t.Weekday().String()
		dc = int(math.Floor(float64((t.Day()-1)/7))) + 1
		fmt.Println("tomorrow!")
	default:
		panic(*argDay + ": args is exception!!")
	}

	// if args -n is not default:0
	switch {
	case *testDay != 0:
		dd := int(math.Floor(float64((*testDay-1)/7))) + 1
		fmt.Println("after(specific day): " + strconv.Itoa(dd))
	default:
		break // do nothing
	}

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
			message += cv.Field(i).Field(2).Interface().(string)
			message += " "
			PostSlack(message, webhookUrl)
		}
	}
}
