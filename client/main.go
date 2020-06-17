package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var client = &http.Client{}

var character = []string{"Iron man", "Captain America", "Spider man", "Black Panther", "Vision", "Hawk eye"}

func main() {
	func1()
	//challenge2()

}
func challenge2() {
	c := make(chan error, 1)
	go initialiseAvengers(c)
	_ = <-c
	go addOnlywhenUsed(c)
	_ = <-c
	//	time.Sleep(time.Minute)
	go useAvengerChallenge(c)
	_ = <-c

}
func func1() {
	c := make(chan error, 1)
	go initialiseAvengers(c)
	_ = <-c
	go addAvenger(c)
	_ = <-c
	go restorePower(c)
	_ = <-c
	go deleteAvenger(c)
	_ = <-c
	//fmt.Println(character)
	go useAvengers(c)
	go updatePower(c)
	_ = <-c

}
func restorePower(c chan error) {
	i := rand.Intn(len(character))
	name := character[i]
	var p []byte
	req, err := http.NewRequest(http.MethodPatch, "http://localhost:9876/avengers/character/"+name, bytes.NewBuffer(p))
	_, err = client.Do(req)
	c <- err
}
func deleteAvenger(c chan error) {
	i := rand.Intn(len(character))
	name := character[i]
	var p []byte
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:9876/avengers/character/"+name, bytes.NewBuffer(p))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	p, err = ioutil.ReadAll(resp.Body)
	name = string(p)
	if name == character[i] {

		character = append(character[0:i], character[i+1:]...)
	}
	c <- err
}

func addOnlywhenUsed(c chan error) {
	var json = make([]string, 4)
	var err error

	json[0] = `{  "name": "Thor", "max_power": 70 }`
	json[1] = `{  "name": "Doctor Strange", "max_power": 78 }`
	json[2] = `{  "name": "Black Widow", "max_power": 40 }`
	json[3] = `{  "name": "Wanda Maximoff", "max_power": 80 }`
	for i := 0; i < 4; i++ {
		_, err = http.Post("http://localhost:9876/avengers/character/challenge", "apllication/json", bytes.NewBuffer([]byte(json[i])))

	}
	character = append(character, "Thor", "Doctor Strange", "Black Widow", "Wanda Maximoff")
	c <- err
}

func addAvenger(c chan error) {
	var json = []byte(`{  "name": "Thor", "max_power": 70 }`)
	var p []byte
	res, err := http.Post("http://localhost:9876/avengers/character", "apllication/json", bytes.NewBuffer(json))
	defer res.Body.Close()
	p, err = ioutil.ReadAll(res.Body)
	checkRemoved(string(p), "Thor")
	json = []byte(`{  "name": "Doctor Strange", "max_power": 78 }`)
	res, err = http.Post("http://localhost:9876/avengers/character", "apllication/json", bytes.NewBuffer(json))
	defer res.Body.Close()
	p, err = ioutil.ReadAll(res.Body)
	checkRemoved(string(p), "Doctor Strange")
	json = []byte(`{  "name": "Black Widow", "max_power": 40 }`)
	res, err = http.Post("http://localhost:9876/avengers/character", "apllication/json", bytes.NewBuffer(json))
	defer res.Body.Close()
	p, err = ioutil.ReadAll(res.Body)
	checkRemoved(string(p), "Black Widow")
	//character = append(character, "Thor", "Doctor Strange", "Black Widow")
	c <- err
}

func checkRemoved(res, name string) {
	if res == "" {
		character = append(character, name)
	} else {
		for i, value := range character {
			if value == res {
				character[i] = name
				break
			}
		}
	}
}

func useAvengers(c chan error) {
	for {
		<-time.After(1 * time.Second)
		i := rand.Intn(len(character))
		use := character[i]
		//fmt.Println(use)
		_, err := http.Get("http://localhost:9876/avengers/character/" + use)
		if err != nil {
			c <- err
			break
		}
	}
}

func useAvengerChallenge(c chan error) {
	for {
		<-time.After(2 * time.Second)
		i := rand.Intn(len(character))
		use := character[i]
		//fmt.Println(use)
		res, err := http.Get("http://localhost:9876/avengers/character/challenge/" + use)
		if err != nil {
			c <- err
			break
		}
		var p []byte
		defer res.Body.Close()
		p, err = ioutil.ReadAll(res.Body)
		use = string(p)
		fmt.Println(use + " " + character[i])
		for i, _ := range character {
			if use == character[i] {
				character = append(character[:i], character[i+1:]...)
				break
			}
		}

	}
}

func updatePower(c chan error) {
	for {
		<-time.After(10 * time.Second)
		for i, _ := range character {
			power := -5 + rand.Intn(10)
			_, err := http.Get("http://localhost:9876/avengers/character/" + character[i] + "/" + strconv.FormatInt(int64(power), 16))
			if err != nil {
				c <- err
				break
			}
		}
	}

}

func initialiseAvengers(c chan error) {
	var json = []byte(`{ "name": "Avengers", "character": [ { "name": "Iron man", "max_power": 60 }, { "name": "Captain America", "max_power": 68 }, { "name": "Spider man", "max_power": 58 }, { "name": "Black Panther", "max_power":68 }, { "name": "Vision", "max_power": 50 }, { "name": "Hawk eye", "max_power": 30 } ] }`)
	_, err := http.Post("http://localhost:9876/avengers", "apllication/json", bytes.NewBuffer(json))
	http.Get("http://localhost:9876/avengers")
	c <- err
}
