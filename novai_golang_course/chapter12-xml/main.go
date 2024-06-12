package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type RSS struct {
	Channel Channel `xml:"channel"`
}

func main() {
	url := "http://feeds.bbci.co.uk/news/rss.xml"

	data, err := getDataFromURL(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, item := range data.Channel.Items {
		fmt.Println("Title:", item.Title)
		fmt.Println("Description:", item.Description)
		fmt.Println("Link:", item.Link)
		fmt.Println("-----------------------")
	}
}

func getDataFromURL(url string) (RSS, error) {
	var rss RSS

	response, err := http.Get(url)
	if err != nil {
		return rss, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return rss, err
	}

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return rss, err
	}

	return rss, nil
}
