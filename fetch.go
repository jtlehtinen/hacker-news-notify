package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiBaseUrl = "https://hacker-news.firebaseio.com/v0/"

type Article struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

func fetch(url string) ([]uint64, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response []uint64
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func fetchTop() ([]uint64, error) {
	return fetch(apiBaseUrl + "topstories.json")
}

func fetchNews() ([]uint64, error) {
	return fetch(apiBaseUrl + "newstories.json")
}

func fetchBest() ([]uint64, error) {
	return fetch(apiBaseUrl + "beststories.json")
}

func fetchStory(id uint64) string {
	url := fmt.Sprintf("%sitem/%d.json", apiBaseUrl, id)

	res, err := http.Get(url)
	if err != nil {
		return ""
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
