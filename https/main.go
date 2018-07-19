package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type StarwarPeople struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

func main() {
	get()
}

func get() {
	url := "https://swapi.co/api/people/1"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)

	var starwarPeople StarwarPeople
	json.Unmarshal(body, &starwarPeople)
	// fmt.Println(starwarPeople)

	blob, _ := json.Marshal(starwarPeople)
	fmt.Println(string(blob))
}

func post() {
}
