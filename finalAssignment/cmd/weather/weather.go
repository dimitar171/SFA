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

func (api WeatherInfo) GetWeather(lat, long string) formatedWeather {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=ae63cfc04efb2375a879a9a9587a7589", lat, long)
	fmt.Println(url)
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
	res := formatedWeather{
		FormatedTemp: temp,
		Description:  st.Weather[0].Description,
		City:         st.Name,
	}
	return res
}
