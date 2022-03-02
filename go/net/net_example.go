package net

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Run() {
	//Make Get call
	resp, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", body)
}
