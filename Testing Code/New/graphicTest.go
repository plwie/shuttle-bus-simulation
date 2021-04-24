package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	// Init termui
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Bus Travelling
	g := widgets.NewGauge()
	g.Title = "Gold Experience Requiem"
	g.Percent = 0
	g.SetRect(0, 12, 50, 15)
	g.BarColor = ui.ColorRed
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorYellow

	// Rendering function
	draw := func() {
		ui.Render(g)
	}

	// Render
	event := ui.PollEvents()
	for {
		select {
		case e := <-event:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		default:
			draw()
			g.Percent += 5
			if g.Percent >= 100 {
				g.Percent = 0
			}
			time.Sleep(time.Second)

		}
	}
}
