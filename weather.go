package main

import "net/http"
import "log"
import "encoding/json"
import "strings"

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main2() {
	http.HandleFunc("/weather/", weatherDataHandler)
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func weatherDataHandler(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]

	data, err := query(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Dude !!"))
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=9aa7f88f84d94f46966fb852e9f2e9e1&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}
