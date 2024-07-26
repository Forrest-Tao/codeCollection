package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GeoIP struct {
	Country   string  `json:"country"`
	Region    string  `json:"region"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func getGeoIP(ip string) (*GeoIP, error) {
	url := fmt.Sprintf("https://freegeoip.app/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geoIP GeoIP
	err = json.Unmarshal(body, &geoIP)
	if err != nil {
		return nil, err
	}
	return &geoIP, nil
}

func main() {
	ip := "8.8.8.8" // 示例IP地址
	geoIP, err := getGeoIP(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("IP地址: %s\n国家: %s\n地区: %s\n城市: %s\n纬度: %.4f\n经度: %.4f\n",
		ip, geoIP.Country, geoIP.Region, geoIP.City, geoIP.Latitude, geoIP.Longitude)
}
