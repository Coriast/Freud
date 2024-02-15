package main

import "Freud"

func startup() {

}

func update() {

}

func render() {

}

func main() {
	app := Freud.App{
		Title:   "Test",
		Width:   800,
		Height:  600,
		Startup: startup,
		Update:  update,
		Render:  render,
	}

	app.Run()
}
