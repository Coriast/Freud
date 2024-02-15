package Freud

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Platform struct {
	version sdl.Version
	window  sdl.Window
}

func (plt *Platform) Startup(app *App) {
	sdl.GetVersion(&plt.version)
	fmt.Printf("SDL %v, %v, %v", plt.version.Major, plt.version.Minor, plt.version.Patch)

	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER | sdl.INIT_EVENTS); err != nil {
		panic(err)
	}
	//defer sdl.Quit()

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_FLAGS, sdl.GL_CONTEXT_FORWARD_COMPATIBLE_FLAG)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24)
	sdl.GLSetAttribute(sdl.GL_STENCIL_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_MULTISAMPLEBUFFERS, 1)
	sdl.GLSetAttribute(sdl.GL_MULTISAMPLESAMPLES, 4)

	window, err := sdl.CreateWindow(app.Title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, app.Width, app.Height, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	//defer window.Destroy()
	plt.window = *window

}

func (plt *Platform) Update(app *App) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			app.running = false
		}
	}
}

func (plt *Platform) Present() {
	plt.window.GLSwap()
}
