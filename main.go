package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type WeatherData struct {
	Name string `json:"name"`
	Base string `json:"base"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
	} `json:"sys"`
}

var (
	view = template.Must(template.ParseFiles("./static/index.html"))
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server := http.Server{
		Addr:    ":9000",
		Handler: mux,
	}

	log.Println("Your app is Running...")

	log.Fatal(server.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	apikey := "5f1bcb7a619ad01056258319bcb94057"
	city := "Pakistan"

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apikey)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var data WeatherData

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

	_ = view.ExecuteTemplate(w, "index.html", data)

	fmt.Printf("Current temperature in '%s'(country is '%s', Base is '%s') is %.1f degree celcius\n", data.Name, data.Sys.Country, data.Base, data.Main.Temp)

}
