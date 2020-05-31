package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
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
	links := []string{
		"http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b",  //avenger
		"http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e",  //mutant
		" http://www.mocky.io/v2/5ed38822340000810001f377", //anti hero
	}

	checkUrls(links)
}

func checkUrls(urls []string) {
	c := make(chan string)
	var wg sync.WaitGroup

	for _, link := range urls {
		wg.Add(1)
		go checkUrl(link, c, &wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for msg := range c {
		fmt.Println(msg)
	}
}

func checkUrl(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	response, err := http.Get(url)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c <- "We could not reach:" + url

	} else {
		//		fmt.Println(string(data))

		var responseObject Response
		json.Unmarshal(data, &responseObject)

		fmt.Println("Name", responseObject.Name)
		fmt.Println("NUmber of char", len(responseObject.Character))

		fmt.Println("All the", responseObject.Name, " name and their Mac power level ")
		for i := 0; i < len(responseObject.Character); i++ {
			fmt.Println(responseObject.Character[i].CharName, responseObject.Character[i].MaxPower)
			// if responseObject.Character[i].CharName == "Iron man" {
			// 	fmt.Println("Requested power of char is ", responseObject.Character[i].MaxPower)
			// }
		}
		fmt.Println()
	}
}
