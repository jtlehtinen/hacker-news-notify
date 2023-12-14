package main

import (
	"time"

	"github.com/getlantern/systray"
)

type application struct {
	notifier *notifier
	config   *config
}

// newApp creates a new application instance.
func newApp(c *config) *application {
	return &application{
		config:   c,
		notifier: newNotifier(),
	}
}

// run starts the system tray application.
func (app *application) run() {
	app.refresh()
	app.notifier.clear()
	systray.Run(app.onReady, app.onExit)
}

// refresh fetches the latest stories from Hacker News API.
func (app *application) refresh() {
	if ids, err := fetchBest(); err == nil {
		app.notifier.add(storyBest, ids...)
	}
	if ids, err := fetchTop(); err == nil {
		app.notifier.add(storyTop, ids...)
	}
	if ids, err := fetchNew(); err == nil {
		app.notifier.add(storyNew, ids...)
	}
}

// onReady is called when the system tray application is ready.
func (app *application) onReady() {
	toggleMenuItem := func(item *systray.MenuItem) {
		if item.Checked() {
			item.Uncheck()
		} else {
			item.Check()
		}
	}

	systray.SetTemplateIcon(IconData, IconData)
	systray.SetTitle("Hacker News Notify")
	systray.SetTooltip("Hacker News Notify")

	top := systray.AddMenuItemCheckbox("Top", "Top", app.config.NotifyTop)
	best := systray.AddMenuItemCheckbox("Best", "Best", app.config.NotifyBest)
	new := systray.AddMenuItemCheckbox("New", "New", app.config.NotifyNew)
	refresh := systray.AddMenuItem("Refresh", "Download latest from Hacker News API")

	systray.AddSeparator()
	quit := systray.AddMenuItem("Exit", "Exit")

	hnFetchTicker := time.NewTicker(3 * time.Minute)
	hnNotifyTicker := time.NewTicker(10 * time.Second)

	go func() {
		// @IMPORTANT: This is the "main" loop or the "main" goroutine.
		for {
			select {
			case <-top.ClickedCh:
				toggleMenuItem(top)
				app.config.NotifyTop = top.Checked()

			case <-best.ClickedCh:
				toggleMenuItem(best)
				app.config.NotifyBest = best.Checked()

			case <-new.ClickedCh:
				toggleMenuItem(new)
				app.config.NotifyNew = new.Checked()

			case <-refresh.ClickedCh:
				app.refresh()

			case <-hnFetchTicker.C:
				app.refresh()

			case <-hnNotifyTicker.C:
				app.notifier.notifyOne(app.config.NotifyTop, app.config.NotifyNew, app.config.NotifyBest)

			case <-quit.ClickedCh:
				hnNotifyTicker.Stop()
				hnFetchTicker.Stop()
				systray.Quit()
				return
			}
		}
	}()
}

// onExit is called when the system tray application is exiting.
func (app *application) onExit() {
	saveConfig(app.config)
}
