package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type config struct {
	NotifyTop  bool `json:"notifyTop"`
	NotifyBest bool `json:"notifyBest"`
	NotifyNew  bool `json:"notifyNew"`
}

var defaultConfig = &config{NotifyTop: true, NotifyBest: true, NotifyNew: true}

func getConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, "hacker-news-notify", "conf.json")

	return path, nil
}

func loadConfig() *config {
	path, err := getConfigPath()
	if err != nil {
		return defaultConfig
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		return defaultConfig
	}

	c := &config{}
	if err := json.Unmarshal(buf, c); err != nil {
		return defaultConfig
	}

	return c
}

func saveConfig(c *config) {
	path, err := getConfigPath()
	if err != nil {
		return
	}

	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return
	}

	os.MkdirAll(filepath.Dir(path), 0644)
	os.WriteFile(path, b, 0644)
}
