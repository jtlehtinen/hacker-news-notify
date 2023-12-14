package main

func main() {
	c := loadConfig()
	app := newApp(c)
	app.run()
	saveConfig(app.config)
}
