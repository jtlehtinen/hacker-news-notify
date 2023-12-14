package main

import (
	"time"

	"github.com/getlantern/systray"
)

func (app *application) onReady() {
	toggle := func(item *systray.MenuItem) {
		if item.Checked() {
			item.Uncheck()
		} else {
			item.Check()
		}
	}

	systray.SetTemplateIcon(IconData, IconData)
	systray.SetTitle("Hacker News Notify")
	systray.SetTooltip("Hacker News Notify")

	top := systray.AddMenuItemCheckbox("Top", "Top", app.getNotifyTop())
	best := systray.AddMenuItemCheckbox("Best", "Best", app.getNotifyBest())
	new := systray.AddMenuItemCheckbox("New", "New", app.getNotifyNew())
	refresh := systray.AddMenuItem("Refresh", "Download latest from Hacker News API")

	systray.AddSeparator()
	quit := systray.AddMenuItem("Exit", "Exit")

	go func() {
		for {
			select {
			case <-top.ClickedCh:
				toggle(top)
				app.setNotifyTop(top.Checked())
			case <-best.ClickedCh:
				toggle(best)
				app.setNotifyBest(best.Checked())
			case <-new.ClickedCh:
				toggle(new)
				app.setNotifyNew(new.Checked())
			case <-refresh.ClickedCh:
				app.refresh()
			case <-quit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func (app *application) run() {
	refreshTicker := time.NewTicker(30 * time.Second)
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

func (app *application) onExit() {

}
