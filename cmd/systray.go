package main

import (
	"github.com/getlantern/systray"
)

func onReady() {
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

	top := systray.AddMenuItemCheckbox("Top", "Top", true)
	best := systray.AddMenuItemCheckbox("Best", "Best", true)
	new := systray.AddMenuItemCheckbox("New", "New", true)

	systray.AddSeparator()
	quit := systray.AddMenuItem("Exit", "Exit")

	go func() {
		for {
			select {
			case <-top.ClickedCh:
				toggle(top)
			case <-best.ClickedCh:
				toggle(new)
			case <-new.ClickedCh:
				toggle(new)
			case <-quit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func run() {
	onExit := func() {}
	systray.Run(onReady, onExit)
}
