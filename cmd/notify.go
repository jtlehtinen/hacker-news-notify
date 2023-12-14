package main

import (
	"log"
	"path/filepath"

	"gopkg.in/toast.v1"
)

func notify() {
	iconPath, err := filepath.Abs("./assets/hnn.png")
	if err != nil {
		log.Fatalln(err)
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

	if err := notification.Push(); err != nil {
		log.Fatalln(err)
	}
}
