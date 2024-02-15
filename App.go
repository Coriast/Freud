package Freud

import "runtime"

type App struct {
	Title   string
	Width   int32
	Height  int32
	Startup func() // Load components of the project
	Update  func() // Updates logic, inputs, per time stuff
	Render  func() // Calls the rendering target to OpenGL with Data

	running bool
}

var platform Platform

func (app *App) Run() {
	runtime.LockOSThread()
	app.running = true

	platform.Startup(app)
	app.Startup()

	for app.running {
		app.step()
	}
}

func (app *App) step() {
	platform.Update(app)
	platform.Present()

	app.Update()
	app.Render()
}
