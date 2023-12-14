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
	Id          int    `json:"id"`
	Score       int    `json:"score"`
	Time        int    `json:"time"` // Unix time
	Descendants int    `json:"descendants"`
	By          string `json:"by"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}

func fetch(url string) ([]int64, error) {
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

	var response []int64
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func fetchTop() ([]int64, error) {
	return fetch(apiBaseUrl + "topstories.json")
}

func fetchNew() ([]int64, error) {
	return fetch(apiBaseUrl + "newstories.json")
}

func fetchBest() ([]int64, error) {
	return fetch(apiBaseUrl + "beststories.json")
}

func fetchStory(id int64) (*Story, error) {
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
