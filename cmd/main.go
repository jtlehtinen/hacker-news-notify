package main

import "sync"

type application struct {
	notifyTop  bool
	notifyBest bool
	notifyNew  bool
	mu         sync.Mutex // Guards everything declared before

	top  []int64
	best []int64
	new  []int64
}

func (app *application) getNotifyTop() bool {
	notify()
	app.mu.Lock()
	defer app.mu.Unlock()
	return app.notifyTop
}

func (app *application) setNotifyTop(value bool) {
	app.mu.Lock()
	app.notifyTop = value
	app.mu.Unlock()
}

func (app *application) getNotifyBest() bool {
	app.mu.Lock()
	defer app.mu.Unlock()
	return app.notifyBest
}

func (app *application) setNotifyBest(value bool) {
	app.mu.Lock()
	app.notifyBest = value
	app.mu.Unlock()
}

func (app *application) getNotifyNew() bool {
	app.mu.Lock()
	defer app.mu.Unlock()
	return app.notifyNew
}

func (app *application) setNotifyNew(value bool) {
	app.mu.Lock()
	app.notifyNew = value
	app.mu.Unlock()
}

func main() {
	app := &application{
		notifyTop:  true,
		notifyBest: true,
		notifyNew:  true,
	}
	app.run()
}
