package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fadhilthomas/go-slack-weather/config"
	"github.com/fadhilthomas/go-slack-weather/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
)

func weatherRequest(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func slackRequest(urlString string, profile model.SlackProfile) ([]byte, error) {
	slackToken := config.Get(config.SLACK_TOKEN)
	profileJson, err := json.Marshal(profile)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest("POST", urlString, bytes.NewBuffer(profileJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", slackToken))
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getWeather() (model.CurrentWeatherResponse, error) {
	cityId := config.Get(config.CITY)
	apiKey := config.Get(config.WEATHER_API)
	var weatherResponse model.CurrentWeatherResponse

	urlString := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%s&appid=%s&units=metric", cityId, apiKey)

	body, err := weatherRequest(urlString)
	if err != nil {
		log.Error().Str("file", "main").Msg(err.Error())
		return weatherResponse, err
	}

	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		log.Error().Str("file", "main").Msg(err.Error())
		return weatherResponse, err
	}
	log.Debug().Str("file", "main").Msg(fmt.Sprintf("%s", weatherResponse))
	return weatherResponse, nil
}

func postSlackStatus(emoji string, text string) error {
	slackProfile := model.SlackProfile{
		Profile: model.Profile{
			StatusEmoji:      emoji,
			StatusExpiration: 0,
			StatusText:       text,
		},
	}

	res, err := slackRequest("https://slack.com/api/users.profile.set", slackProfile)
	if err != nil {
		log.Error().Str("file", "main").Msg(err.Error())
		return err
	}
	log.Debug().Str("file", "main").Msg(fmt.Sprintf("%s", res))
	return nil
}

func transEmoji(icon string) string {
	var emojiList = map[string]string{
		"01": ":sunny:",
		"02": ":partly_sunny:",
		"03": ":cloud:",
		"04": ":partly_sunny:",
		"09": ":rain_cloud:",
		"10": ":rain_cloud:",
		"11": ":thunder_cloud_and_rain:",
		"13": ":snow_cloud:",
		"50": ":cloud:",
	}
	emoji := emojiList[icon]
	log.Debug().Str("file", "main").Msg(fmt.Sprintf("%s", icon))
	return emoji
}

func timeIn(t time.Time) (time.Time, error) {
	timezone := config.Get(config.TIMEZONE)
	loc, err := time.LoadLocation(timezone)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func main() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	hour, err := timeIn(time.Now())
	if err != nil{
		return
	}

	weather, err := getWeather()
	if err != nil {
		return
	}
	icon := weather.Weather[0].Icon[:2]
	temp := weather.Temp
	weatherMain := weather.Weather[0].Main

	if icon != "" {
		emoji := transEmoji(icon)
		err := postSlackStatus(emoji, fmt.Sprintf("%s - %.2f°C - %s", weatherMain, temp, hour.Format("15:04")))
		if err != nil {
			return
		}
	}
}
