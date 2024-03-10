package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func getUrls(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer f.Close()

	fs := bufio.NewScanner(f)

	var lines []string
	for fs.Scan() {
		lines = append(lines, fs.Text())
	}
	if err := fs.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!!!")
		c <- link
		return
	}
	fmt.Println(link + " is Ok.")
	c <- link
}

func checkLinks(links []string) {
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		//go checkLink(l, c)
		go func(link string, waitInSeconds int32) {
			time.Sleep(time.Duration(waitInSeconds) * time.Second)
			checkLink(link, c)
		}(l, 2)
	}

	/*
		for i := 0; i < len(links); i++ {
			fmt.Println(<-c)
		}
	*/

	/*
		for {
			go checkLink(<-c, c)
		}
	*/

}
