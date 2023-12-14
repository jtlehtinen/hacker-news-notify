package main

import (
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
				app.setNotifyNew(best.Checked())
			case <-new.ClickedCh:
				toggle(new)
				app.setNotifyNew(new.Checked())
			case <-quit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func (app *application) run() {
	systray.Run(app.onReady, app.onExit)
}

func (app *application) onExit() {

}
