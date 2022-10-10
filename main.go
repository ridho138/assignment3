package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Detail struct {
	Water, Wind int
}

type Result struct {
	Status Detail
}

func main() {

	uptimeTicker := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-uptimeTicker.C:
			fmt.Println(writeData())
		}
	}
}

func writeData() string {
	data := Detail{
		Water: randomNumber(),
		Wind:  randomNumber(),
	}

	dataResult := Result{
		Status: data,
	}

	file, _ := json.MarshalIndent(dataResult, "", " ")

	err := ioutil.WriteFile("data.json", file, 0644)
	if err != nil {
		return err.Error()
	}

	return "Data Updated"
}

func randomNumber() int {
	min := 1
	max := 100

	random := rand.Intn(max-min) + min

	return random
}
