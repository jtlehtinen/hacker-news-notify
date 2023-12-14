package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiBaseUrl = "https://hacker-news.firebaseio.com/v0/"

type Story struct {
	Id          int32  `json:"id"`
	Score       int32  `json:"score"`
	Time        int32  `json:"time"` // Unix time; Y2K38
	Descendants int32  `json:"descendants"`
	By          string `json:"by"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}

func fetch(url string) ([]int32, error) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	const maxBytesToRead = 10 * 1_024 * 1_024
	body, err := io.ReadAll(io.LimitReader(res.Body, maxBytesToRead))
	if err != nil {
		return nil, err
	}

	var response []int32
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func fetchTop() ([]int32, error) {
	return fetch(apiBaseUrl + "topstories.json")
}

func fetchNew() ([]int32, error) {
	return fetch(apiBaseUrl + "newstories.json")
}

func fetchBest() ([]int32, error) {
	return fetch(apiBaseUrl + "beststories.json")
}

func fetchStory(id int32) (*Story, error) {
	url := fmt.Sprintf("%sitem/%d.json", apiBaseUrl, id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	const maxBytesToRead = 10 * 1_024 * 1_024
	body, err := io.ReadAll(io.LimitReader(res.Body, maxBytesToRead))
	if err != nil {
		return nil, err
	}

	response := &Story{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
