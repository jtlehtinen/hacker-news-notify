package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiBaseUrl = "https://hacker-news.firebaseio.com/v0/"

type Article struct {
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
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
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

func fetchStory(id int64) (*Article, error) {
	url := fmt.Sprintf("%sitem/%d.json", apiBaseUrl, id)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Article{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
