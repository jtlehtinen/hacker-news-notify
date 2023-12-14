package main

import (
	"fmt"
	"path/filepath"

	"github.com/go-toast/toast"
)

type storyKind int

const (
	storyNew storyKind = iota
	storyTop
	storyBest
)

type entry struct {
	kind storyKind
	id   int64
}

type notifier struct {
	entries map[int64]entry
	pending []entry
}

func newNotifier() *notifier {
	return &notifier{
		entries: make(map[int64]entry),
	}
}

func (n *notifier) add(kind storyKind, ids ...int64) {
	for _, id := range ids {
		if _, ok := n.entries[id]; !ok {
			e := entry{kind: kind, id: id}
			n.entries[id] = e
			n.pending = append(n.pending, e)
		}
	}
}

func (n *notifier) clearPending() {
	n.pending = nil
}

var kindToTitle = map[storyKind]string{
	storyTop:  "Top Story",
	storyNew:  "New Story",
	storyBest: "Best Story",
}

func (n *notifier) notifyOne(notifyTop, notifyNew, notifyBest bool) {
	accept := func(kind storyKind) bool {
		switch kind {
		case storyTop:
			return notifyTop
		case storyNew:
			return notifyNew
		case storyBest:
			return notifyBest
		}
		return false
	}

	for len(n.pending) > 0 {
		e := n.pending[0]
		n.pending = n.pending[1:]

		if accept(e.kind) {
			go func() {
				if story, err := fetchStory(e.id); err == nil {
					showToast(kindToTitle[e.kind], story.Title, story.Url, getHackerNewsUrl(story))
				}
			}()
			break
		}
	}
}

func getHackerNewsUrl(story *Story) string {
	return fmt.Sprintf("https://news.ycombinator.com/item?id=%d", story.Id)
}

func showToast(title, message, url, urlHackerNews string) error {
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
