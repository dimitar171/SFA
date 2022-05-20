package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handleGetWeather(infos WeatherInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resultWeather := infos
		json.NewEncoder(w).Encode(resultWeather)
	}
}

func TestGetWeather(t *testing.T) {
	router := http.NewServeMux()

	WantedWeather := formatedWeather{
		// FormatedTemp:"25",
		// Description:"DescTest",
		City: "CityTest",
	}
	infos := WeatherInfo{

		Name: "CityTest",
	}

	urlTest := "https://api.openweathermap.org/data/2.5/weather?lat=5&lon=5&appid=ae63cfc04efb2375a879a9a9587a7589"
	router.Handle(urlTest, handleGetWeather(infos))
	mockServer := httptest.NewServer(router)
	ww := NewWeather(mockServer.URL)
	fmt.Println(mockServer.URL)

	res := ww.GetWeather("5", "5")

	if !reflect.DeepEqual(res, WantedWeather) {
		t.Fatalf("got: %v, want %v", res, WantedWeather)
	}
}
