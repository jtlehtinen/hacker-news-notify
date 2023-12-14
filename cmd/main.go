package main

func main() {
	app := &application{
		notifyTop:   true,
		notifyBest:  true,
		notifyNew:   true,
		notifyQueue: make([]*NotifyQueueEntry, 0),
		seen:        make(map[int64]void),
	}

	app.run()
}
