package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://localhost:8080/ttt1")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
	
		
