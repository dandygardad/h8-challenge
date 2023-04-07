package main

import (
	"challenge08/model"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"
)

// ProcessStatus: memproses angka water dan wind
func ProcessStatus(water int, wind int) {
	var statusWater string
	var statusWind string

	// Check status water
	if water <= 5 {
		statusWater = "Aman"
	} else if water >= 6 && water <= 8 {
		statusWater = "Siaga"
	} else {
		statusWater = "Bahaya"
	}

	// Check status wind
	if wind <= 6 {
		statusWind = "Aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "Siaga"
	} else {
		statusWind = "Bahaya"
	}
	fmt.Println("Status water :", statusWater)
	fmt.Println("Status wind :", statusWind)
}

// SimulationPost: melakukan simulasi post data
func SimulationPost(urlPath string) ([]byte, error) {
	_, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	// Untuk generate angka random
	rand.Seed(time.Now().UnixNano())
	response := model.Response{
		Water: rand.Intn(16-1) + 1,
		Wind:  rand.Intn(16-1) + 1,
	}

	// Kembalikan response dalam struct menjadi JSON dengan encode
	resultJson, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return resultJson, nil
}

func main() {
	fmt.Println("Melakukan post data..")
	urlPath := "https://jsonplaceholder.typicode.com/posts"
	go func() {
		// Mengirim post data setiap 15 detik
		for range time.Tick(time.Second * 15) {
			resultJson, err := SimulationPost(urlPath)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(resultJson)) // Tampilkan hasil json

			// Decode json
			var decodeData model.Response
			err = json.Unmarshal(resultJson, &decodeData)
			if err != nil {
				log.Fatal(err)
			}
			ProcessStatus(decodeData.Water, decodeData.Wind)
		}
	}()
	time.Sleep(1 * time.Minute) // Batasi simulasi post hanya 1 menit
}
