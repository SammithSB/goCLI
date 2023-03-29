package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	openWeatherMapURL = "https://api.openweathermap.org/data/2.5/weather"
)

type weatherData struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func main() {
	app := &cli.App{
		Name:  "Weather CLI",
		Usage: "Get the current weather for a city",
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				fmt.Println("Please provide either the city name or the ZIP code")
				return nil
			}
			location := c.Args().Get(0)
			units := "metric"
			// get api from .env file as apiKey
			apiKey := os.Getenv("OPENWEATHER_API_KEY")
			if apiKey == "" {
				fmt.Println("Please set the OPENWEATHER_API_KEY environment variable")
				return nil
			}
			url := fmt.Sprintf("%s?q=%s&units=%s&appid=%s", openWeatherMapURL, location, units, apiKey)
			response, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error fetching weather data: %s\n", err.Error())
				return nil
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("Error reading weather data: %s\n", err.Error())
				return nil
			}
			var data weatherData
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Printf("Error parsing weather data: %s\n", err.Error())
				return nil
			}
			fmt.Printf("Weather for %s:\n", location)
			fmt.Printf("Condition: %s\n", data.Weather[0].Main)
			fmt.Printf("Temperature: %.1f °C\n", data.Main.Temp)
			fmt.Printf("Humidity: %d %%\n", data.Main.Humidity)
			fmt.Printf("Wind speed: %.1f m/s\n", data.Wind.Speed)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error running app: %s\n", err.Error())
	}
}
