package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

var API_KEY string = "07bc3b6ddad9e6e1a9882a0b2509007f"
var API_URL string = "http://api.positionstack.com/v1/forward"

func main() {
	raw, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var data []map[string]interface{}
	err = json.Unmarshal(raw, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range data {
		delete(v, "Other_Details")
	}

	for _, v := range data {
		location := v["Name"].(string) + " " + v["Park_Name"].(string)
		lon, lat, err := GetLatAndLong(location)
		if err != nil {
			fmt.Println(err.Error())
		}
		v["lon"] = lon
		v["lat"] = lat
	}

	fmt.Println(data)

}

func GetLatAndLong(location string) (float64, float64, error) {
	query := "?access_key=" + url.QueryEscape(API_KEY) + "&query=" + url.QueryEscape(location)
	fmt.Println(query)
	url := API_URL + query
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return 0, 0, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	body = []byte(body)

	if err != nil {
		log.Fatal(err)
		return 0, 0, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
		return 0, 0, err
	}

	fmt.Println(data)

	if reflect.ValueOf(data).IsNil() {
		return 0, 0, nil
	}

	var check []interface{}

	if reflect.TypeOf(data) == reflect.TypeOf(check) {
		return 0, 0, nil
	}

	checkinn := data["data"]

	if reflect.ValueOf(checkinn).Len() == 0 {
		return 0, 0, nil
	}

	fmt.Println(checkinn)

	results := data["data"].([]interface{})[0]

	if reflect.TypeOf(results) == reflect.TypeOf(check) {
		return 0, 0, nil
	}

	inndata := results.(map[string]interface{})

	longitude := inndata["longitude"].(float64)
	latitude := inndata["latitude"].(float64)

	return longitude, latitude, nil
}
