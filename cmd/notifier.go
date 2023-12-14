package main

import (
	"fmt"
	"path/filepath"

	"github.com/go-toast/toast"
)

type void struct{}
type storyKind int32

const (
	storyNew storyKind = iota
	storyTop
	storyBest
)

// @IMPORTANT: The max id can be fetched from
// https://hacker-news.firebaseio.com/v0/maxitem.json
// 32-bit integer more than enough.
type entry struct {
	id   int32
	kind storyKind
}

type notifier struct {
	seen  map[entry]void // Entries (id + kind pair) we have received.
	toast map[int32]void // Story ids for which toast has been shown.
	queue []entry        // Entries waiting to be toasted.
}

func newNotifier() *notifier {
	return &notifier{
		seen:  make(map[entry]void),
		toast: make(map[int32]void),
	}
}

func (n *notifier) add(kind storyKind, ids ...int32) {
	for _, id := range ids {
		e := entry{id: id, kind: kind}
		if _, ok := n.seen[e]; !ok {
			n.seen[e] = void{}
			n.queue = append(n.queue, e)
		}
	}
}

func (n *notifier) clear() {
	n.queue = nil
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

	for len(n.queue) > 0 {
		e := n.queue[0]
		n.queue = n.queue[1:]

		if _, ok := n.toast[e.id]; ok {
			continue
		}

		if accept(e.kind) {
			n.toast[e.id] = void{}
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
