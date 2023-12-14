package main

import (
	"path/filepath"

	"gopkg.in/toast.v1"
)

func notify() error {
	iconPath, err := filepath.Abs("./assets/hnn.png")
	if err != nil {
		return err
	}

	notification := toast.Notification{
		AppID:   "Hacker News Notify",
		Title:   "Title here",
		Message: "Some message here...",
		Icon:    iconPath,
		Actions: []toast.Action{
			{Type: "protocol", Label: "OK", Arguments: ""},
			{Type: "protocol", Label: "Open", Arguments: ""},
		},
	}

	return notification.Push()
}
