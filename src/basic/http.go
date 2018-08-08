package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	   client := &http.Client{}
		req, err := http.NewRequest("GET", "http://album.zhenai.com/u/1085771778", nil)
		if err != nil{
			panic(err)
		}
		req.Header.Add("User-Agent", "myClient")
		resp, err := client.Do(req)
		if err!=nil{
			panic(err)
		}
		defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s",bytes)
}
