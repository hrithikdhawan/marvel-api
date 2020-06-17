package data

import (
	"encoding/json"
	"fmt"
	"io"
)

const maxCharacters = 8

type Marvel struct {
	Name      string       `json:"name" validate:"unique"`
	Character []Characters `json:"character"`
}
type Characters struct {
	Name         string `json:"name"`
	MaxPower     int    `json:"max_power"`
	CurrentPower int    `json:"current_power"`
	Count        int    `json:"count"`
}

func (m *Marvel) ToJSON(w io.Writer) error {
	encode := json.NewEncoder(w)
	return encode.Encode(m)
}

func (c *Characters) ToJSON(w io.Writer) error {
	encode := json.NewEncoder(w)
	return encode.Encode(c)
}

func (m *Marvel) FromJSON(r io.Reader) error {
	decode := json.NewDecoder(r)
	return decode.Decode(m)
}

func (c *Characters) FromJSON(r io.Reader) error {
	decode := json.NewDecoder(r)
	return decode.Decode(c)
}

func findLeastUsed(marvel Marvel) int {
	min := marvel.Character[0].Count
	index := 0
	for i, value := range marvel.Character {
		if min > value.Count {
			min = value.Count
			index = i
		}
	}
	return index
}
func findLeastPower(marvel Marvel) int {
	min := marvel.Character[0].CurrentPower
	index := 0
	for i, value := range marvel.Character {
		if min > value.CurrentPower {
			min = value.CurrentPower
			index = i
		}
	}
	return index
}

var ErrorNotFound = fmt.Errorf("Character not found")
