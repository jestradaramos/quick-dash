package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://community-open-weather-map.p.rapidapi.com/weather?q=02721&lat=0&lon=0&callback=test&id=2172797&lang=null&units=imperial&mode=xml"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "c404777dcbmshc15364aefe92996p1f6149jsn9576ffac9803")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

// from rapidapi
