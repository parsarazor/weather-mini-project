// what does servieces?
package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
)

type apiResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}
func GetWeather(city string) (float64, error) {
	// what the actual fuck sprinf doing man
	const apiKey = "b511e68a6762ae4780840786b558d776"
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		neturl.QueryEscape(city),
		apiKey,
	)
	resp, err := http.Get(url)
	
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	var data apiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}
	
	return data.Main.Temp, nil
}