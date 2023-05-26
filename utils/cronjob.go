package utils

import (
	"assignment-3/models"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	url         = "http://localhost:8080/api/data"
	delaySecond = 15
	randomMin   = 1
	randomMax   = 100
)

func CronJob() {
	for {
		time.Sleep(delaySecond * time.Second)

		data := models.Data{
			Water: randomInt(),
			Wind:  randomInt(),
		}

		jsonData, _ := json.MarshalIndent(data, "", "\t")
		jsonString := string(jsonData)
		fmt.Println(jsonString)

		err := updateData(jsonString)
		if err == nil {
			water, wind := findStatus(data.Water, data.Wind)
			fmt.Printf("status water : %s\nstatus wind : %s\n\n", water, wind)
		}
	}
}

func randomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(randomMax) + randomMin
}

func updateData(data string) error {
	client := &http.Client{}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func findStatus(water int, wind int) (string, string) {
	bahaya := "bahaya"
	aman := "aman"
	siaga := "siaga"

	waterStatus := bahaya
	if water < 5 {
		waterStatus = aman
	} else if water >= 6 && water <= 8 {
		waterStatus = siaga
	}

	windStatus := bahaya
	if wind < 6 {
		windStatus = aman
	} else if wind >= 7 && wind <= 15 {
		windStatus = siaga
	}

	return waterStatus, windStatus
}
