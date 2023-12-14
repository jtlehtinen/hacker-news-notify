package main

import (
	"path/filepath"

	"github.com/go-toast/toast"
)

func notify(title, message, url, urlHackerNews string) error {
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
			{Type: "protocol", Label: "Open HN", Arguments: urlHackerNews},
		},
		Duration: toast.Long,
	}

	return notification.Push()
}
