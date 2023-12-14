package main

import (
	"path/filepath"

	"gopkg.in/toast.v1"
)

func notify(title, message, url string) error {
	iconPath, err := filepath.Abs("./assets/hnn.png")
	if err != nil {
		return err
	}

	notification := toast.Notification{
		AppID:   "Hacker News Notify",
		Title:   title,
		Message: message,
		Icon:    iconPath,
		Actions: []toast.Action{
			{Type: "protocol", Label: "Open", Arguments: url},
		},
	}

	return notification.Push()
}
