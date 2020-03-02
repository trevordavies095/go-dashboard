package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Config holds data imported from config.json also passed into template
type Config struct {
	DarkMode         bool   `json:"dark-mode"`
	Bookmark1Name    string `json:"bookmark1name"`
	Bookmark1URL     string `json:"bookmark1url"`
	Bookmark2Name    string `json:"bookmark2name"`
	Bookmark2URL     string `json:"bookmark2url"`
	Bookmark3Name    string `json:"bookmark3name"`
	Bookmark3URL     string `json:"bookmark3url"`
	Name             string `json:"name"`
	WelcomeMsg       string
	DarkSkySecretKey string `json:"darksky-secretkey"`
	Lat              string `json:"lat"`
	Lon              string `json:"lon"`
	WeatherData      struct {
		TodaySummary    string
		TodayHigh       float64
		TodayLow        float64
		TomorrowSummary string
		TomorrowHigh    float64
		TomorrowLow     float64
	}
}

// WeatherData holds data from DarkSky call
type WeatherAPI struct {
	Daily struct {
		Icon string                   `json:"icon"`
		Days []map[string]interface{} `json:"data"`
	} `json:"daily"`
}

// Global constants

// Global variables
var tpl *template.Template
var config Config

func init() {
	// Local constants

	// Local variables

	/****** start init() ******/

	tpl = template.Must(template.ParseGlob("templates/*"))
	config = openConfig()

}

func main() {
	// Local constants

	// Local variables

	/****** start main() ******/

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	// Local constants

	// Local variables

	/****** start index() ******/

	genWelcomeMsg()
	getWeather()

	tpl.ExecuteTemplate(w, "index.html", config)
}

func genWelcomeMsg() {
	// Local constants

	// Local variables
	time := time.Now()
	year, month, day := time.Date()
	weekday := time.Weekday().String()
	hour := time.Hour()

	/****** start genWelcomeMsg() ******/

	if hour >= 5 && hour < 12 {
		config.WelcomeMsg = "Good morning "
	} else if hour >= 12 && hour < 17 {
		config.WelcomeMsg = "Good afternoon "
	} else {
		config.WelcomeMsg = "Good evening "
	}
	config.WelcomeMsg += config.Name + ". Today is " + weekday + ", " + month.String() + " " + strconv.Itoa(day) + ", " + strconv.Itoa(year) + "."
}

func getWeather() {
	// Local constants
	const excluded = "?exclude=currently,minutely,hourly,alerts,flags"

	// Local variables
	apiCall := "https://api.darksky.net/forecast/" + config.DarkSkySecretKey + "/" + config.Lat + "," + config.Lon + excluded
	resp, err := http.Get(apiCall)
	var weatherAPI WeatherAPI

	/****** start getWeather() ******/

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	byteValue, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(byteValue), &weatherAPI)

	today := weatherAPI.Daily.Days[0]
	tomorrow := weatherAPI.Daily.Days[1]

	config.WeatherData.TodaySummary = today["summary"].(string)
	config.WeatherData.TodayHigh = today["temperatureHigh"].(float64)
	config.WeatherData.TodayLow = today["temperatureLow"].(float64)
	config.WeatherData.TomorrowSummary = tomorrow["summary"].(string)
	config.WeatherData.TomorrowHigh = tomorrow["temperatureHigh"].(float64)
	config.WeatherData.TomorrowLow = tomorrow["temperatureLow"].(float64)
}

func openConfig() Config {
	// Local constants

	// Local variables
	configFile, err := os.Open("config.json")
	var config Config

	/****** start open_config() ******/

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, &config)

	configFile.Close()
	return config
}
