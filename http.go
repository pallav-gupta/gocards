package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func httpGet() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := make([]byte, 99999) // 9999 -> initial number of elements
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
