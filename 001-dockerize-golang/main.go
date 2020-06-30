package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hello")
	resp, err := http.Get("http://169.24.2.82:8000/")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(len(body))
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
