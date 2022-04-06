package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Drink struct {
	Instructions string `json:"strInstructions"`
}
type DrinksResponcePayload struct {
	Drinks []Drink
}
type Bartender struct {
	url             string
	strInstructions DrinksResponcePayload
	recepy          string
}

func NewBartender(url string) Bartender {
	return Bartender{url: url}
}

func (ba *Bartender) Start() { //Start is a method on barthender
	for {
		//1. Ask the user what the user wants to drink
		var UserAsk string
		fmt.Println("What do you want to drink: ")
		fmt.Scanln(&UserAsk)
		//4.repeat until user enters nothing
		if UserAsk == "nothing" {
			return
		}
		//constructing the url from user
		u, err := url.Parse(ba.url)
		if err != nil {
			log.Fatal(err)
		}
		q := u.Query()
		q.Set("s", UserAsk)
		u.RawQuery = q.Encode()
		ba.url = u.String()
		//2. pass to user input to API, get recepy and store the first result it in field strInstructions
		req, err := http.NewRequest("GET", ba.url, nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			return
		}
		payload := DrinksResponcePayload{}
		json.NewDecoder(resp.Body).Decode(&payload)
		ba.recepy = payload.Drinks[0].Instructions
		time.Sleep(time.Second)
		resp.Body.Close()

		//3. Split the recipe into different sentences using strings.Split and print each one back to the user on a new line.
		temp := strings.Split(ba.recepy, ".")
		for _, te := range temp {
			fmt.Println(te)
		}
	}
}

func main() {
	ba := NewBartender("http://www.thecocktaildb.com/api/json/v1/1/search.php?s=margarita")
	ba.Start()
}
