package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Name      string      `json:"name"`
	Character []Character `json:"character"`
}

type Character struct {
	CharName string `json:"name"`
	MaxPower int    `json:"max_power"`
}

func main() {

	response, err := http.Get("http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b")
	if err != nil {
		fmt.Println("request failed with error %s \n", err)
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("request failed with error %s \n", err)
	}
	//fmt.Println(string(data))

	var responseObject Response
	json.Unmarshal(data, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Character))

	for i := 0; i < len(responseObject.Character); i++ {
		fmt.Println(responseObject.Character[i].CharName, responseObject.Character[i].MaxPower)
	}

}
