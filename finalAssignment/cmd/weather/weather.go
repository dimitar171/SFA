package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WeatherInfo struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}
type formatedWeather struct {
	FormatedTemp string `json:"formatedTemp"`
	Description  string `json:"description"`
	City         string `json:"city"`
}
type WeatherUrl struct{ urlBase string }

func NewWeather(url string) *WeatherUrl {
	return &WeatherUrl{urlBase: url}
}

func (wp *WeatherUrl) GetWeather(lat, long string) formatedWeather {
	wetURL := wp.urlBase
	url := fmt.Sprintf(wetURL, lat, long)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var st WeatherInfo
	json.NewDecoder(resp.Body).Decode(&st)
	temp := fmt.Sprintf("%v", st.Main.Temp)
	result := formatedWeather{
		FormatedTemp: temp,
		Description:  st.Weather[0].Description,
		City:         st.Name,
	}
	return result
}
