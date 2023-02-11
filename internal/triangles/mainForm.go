package triangles

import (
	"os"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/softteam/framework"
)

type MainForm struct {
	window  *gtk.ApplicationWindow
	builder *framework.GtkBuilder
	da      *gtk.DrawingArea
}

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	fw := framework.NewFramework()
	builder, err := fw.Gtk.CreateBuilder("main.glade")
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.window = m.builder.GetObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.window.SetApplication(app)
	m.window.SetTitle("triangles main window")
	m.window.Maximize()

	// Hook up signals
	m.window.Connect("destroy", m.window.Close)
	m.window.Connect("key-press-event", m.onKeyPress)

	// Quit button
	button := m.builder.GetObject("main_window_quit_button").(*gtk.ToolButton)
	button.Connect("clicked", m.window.Close)

	// Status bar
	statusBar := m.builder.GetObject("main_window_status_bar").(*gtk.Statusbar)
	statusBar.Push(statusBar.GetContextId("triangles"), "triangles : version 0.1.0")

	// Menu
	m.setupMenu()

	// Show the main window
	m.window.ShowAll()

	// Drawing area
	da := m.builder.GetObject("drawingArea").(*gtk.DrawingArea)
	m.da = da
	da.Connect("draw", m.onDraw)
	da.SetCanFocus(true)
	da.GrabFocus()
}

func (m *MainForm) setupMenu() {
	menuQuit := m.builder.GetObject("menu_file_quit").(*gtk.MenuItem)
	menuQuit.Connect("activate", m.window.Close)
}

func (m *MainForm) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	w, h := float64(da.GetAllocatedWidth()), float64(da.GetAllocatedHeight())
	ctx.SetSourceRGB(1, 1, 1)
	ctx.Rectangle(0, 0, w, h)
	ctx.Fill()

	if len(triangles) == 0 {
		createInitialTriangle(h, w)
	}

	m.drawTriangles(ctx)
}

func (m *MainForm) drawTriangles(ctx *cairo.Context) {
	for _, t := range triangles {
		m.drawTriangle(ctx, t)
	}
}

func (m *MainForm) drawTriangle(ctx *cairo.Context, t triangle) {
	ctx.SetSourceRGB(0, 0, 0)
	ctx.MoveTo(t.p1.x, t.p1.y)
	ctx.LineTo(t.p2.x, t.p2.y)
	ctx.LineTo(t.p3.x, t.p3.y)
	ctx.LineTo(t.p1.x, t.p1.y)
	ctx.Stroke()
}

func (m *MainForm) onKeyPress(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)
	switch ke.KeyVal() {
	case 32, 83, 115: // space, s and S
		subDivideTriangles()
	case 67, 99: // c and C
		clearTriangles()
	case 81, 113: // q and Q
		m.window.Close()
		return
	}
	m.da.QueueDraw()
}
