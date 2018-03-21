package commands

import (
	"encoding/json"
	"fmt"
	"github.com/codeclimate/hestia/internal/config"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Weather struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

type Forecast struct {
	City        string `json:"city"`
	Zip         string `json:"zip"`
	Description string `json:"description"`
}

type WeatherResponse struct {
	Main        string `json:"main"`
	Description string `json"description"`
}

type OpenWeatherMapResponse struct {
	Name    string            `json:"name"`
	Weather []WeatherResponse `json:"weather"`
}

func getWeather(zip string, country_code string) (weather OpenWeatherMapResponse) {
	weather_url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?zip=%s,%s&appid=%s",
		zip,
		country_code,
		config.Fetch("open_weather_api_key"),
	)

	resp, err := http.Get(weather_url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	weather = OpenWeatherMapResponse{}
	_ = json.Unmarshal(bodyBytes, &weather)

	return weather
}

func (c Weather) Run() {
	parts := strings.Split(c.Input.Args, " ")

	zip := "10011"
	country_code := "us"

	if len(parts) > 0 && parts[0] != "" {
		zip = parts[0]
	}

	if len(parts) > 1 && parts[1] != "" {
		country_code = parts[1]
	}

	weather := getWeather(zip, country_code)

	forecast := Forecast{
		City:        weather.Name,
		Zip:         zip,
		Description: weather.Weather[0].Description,
	}

	message := fmt.Sprintf("%s in %s (%s)", forecast.Description, forecast.City, forecast.Zip)

	c.Notifier.Log(message)
}

func (c Weather) HelpText() string {
	return "weather [zipcode=10011]"
}

func (c Weather) HelpDescription() string {
	return "Fetches the current weather forecast"
}

func (c Weather) HelpExamples() []string {
	return []string{"weather", "weather 10012"}
}
