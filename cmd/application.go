package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/getlantern/systray"
)

// @TODO: Clean up this mess.

type void struct{}

type NotifyQueueEntry struct {
	title string
	id    int64
}

type application struct {
	config      *config
	notifyQueue []*NotifyQueueEntry
	seen        map[int64]void
	mu          sync.Mutex // Guards everything

	notifyTicker     *time.Ticker
	notifyTickerDone chan bool
}

func newApp(config *config) *application {
	return &application{
		config:      config,
		notifyQueue: make([]*NotifyQueueEntry, 0),
		seen:        make(map[int64]void),
	}
}

func filterNew(ids []int64, seen map[int64]void) []int64 {
	result := make([]int64, 0)
	for _, id := range ids {
		_, ok := seen[id]
		if !ok {
			result = append(result, id)
		}
	}
	return result
}

func getHackerNewsUrl(story *Story) string {
	return fmt.Sprintf("https://news.ycombinator.com/item?id=%d", story.Id)
}

func (app *application) startNotifyRoutine() {
	app.notifyTicker = time.NewTicker(15 * time.Second)
	app.notifyTickerDone = make(chan bool)
	go func() {
		for {
			select {
			case <-app.notifyTickerDone:
				return
			case <-app.notifyTicker.C:
				app.mu.Lock()
				var entry *NotifyQueueEntry
				if len(app.notifyQueue) > 0 {
					entry = app.notifyQueue[0]
					app.notifyQueue = app.notifyQueue[1:]
				}
				app.mu.Unlock()

				if entry != nil {
					story, _ := fetchStory(entry.id)
					notify(entry.title, story.Title, story.Url, getHackerNewsUrl(story))
				}
			}
		}
	}()
}

func (app *application) stopNotifyRoutine() {
	app.notifyTicker.Stop()
	app.notifyTickerDone <- true
}

func (app *application) getNotifyTop() bool {
	app.mu.Lock()
	defer app.mu.Unlock()
	return app.config.NotifyTop
}

func (app *application) setNotifyTop(value bool) {
	app.mu.Lock()
	app.config.NotifyTop = value
	app.mu.Unlock()
}

func (app *application) getNotifyBest() bool {
	app.mu.Lock()
	defer app.mu.Unlock()
	return app.config.NotifyBest
}

func (app *application) setNotifyBest(value bool) {
	app.mu.Lock()
	app.config.NotifyBest = value
	app.mu.Unlock()
}

func (app *application) getNotifyNew() bool {
	app.mu.Lock()
	defer app.mu.Unlock()
	return app.config.NotifyNew
}

func (app *application) setNotifyNew(value bool) {
	app.mu.Lock()
	app.config.NotifyNew = value
	app.mu.Unlock()
}

func (app *application) addToNotifyQueue(notifyTitle string, stories []int64) {
	if len(stories) == 0 {
		return
	}

	app.mu.Lock()
	defer app.mu.Unlock()
	for _, id := range stories {
		fmt.Println(notifyTitle, id)
		app.notifyQueue = append(app.notifyQueue, &NotifyQueueEntry{title: notifyTitle, id: id})
	}
}

func (app *application) refresh() {
	// @TODO: Separate go routine. Don't block tray app.
	// @TODO: Process story only once.
	top, _ := fetchTop()
	new, _ := fetchNew()
	best, _ := fetchBest()

	newTop := filterNew(top, app.seen)
	newNew := filterNew(new, app.seen)
	newBest := filterNew(best, app.seen)

	firstRefresh := len(app.seen) == 0
	if !firstRefresh {
		if app.config.NotifyTop {
			app.addToNotifyQueue("Top Story", newTop)
		}
		if app.config.NotifyBest {
			app.addToNotifyQueue("Best Story", newBest)
		}
		if app.config.NotifyNew {
			app.addToNotifyQueue("New Story", newNew)
		}
	}

	app.mu.Lock()
	defer app.mu.Unlock()

	for _, id := range newTop {
		app.seen[id] = void{}
	}
	for _, id := range newBest {
		app.seen[id] = void{}
	}
	for _, id := range newNew {
		app.seen[id] = void{}
	}
}

func (app *application) run() {
	refreshTicker := time.NewTicker(3 * time.Minute)
	refreshTickerDone := make(chan bool)

	go func() {
		for {
			select {
			case <-refreshTickerDone:
				return
			case <-refreshTicker.C:
				app.refresh()
			}
		}
	}()
	app.refresh()

	app.startNotifyRoutine()
	systray.Run(app.onReady, app.onExit)
	app.stopNotifyRoutine()

	refreshTicker.Stop()
	refreshTickerDone <- true
}
