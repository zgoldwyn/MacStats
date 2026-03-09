package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func startUI(ch <-chan Stats) {
	if err := ui.Init(); err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// CPU gauge
	cpuGauge := widgets.NewGauge()
	cpuGauge.Title = "CPU Usage"
	cpuGauge.SetRect(0, 0, 50, 5)
	cpuGauge.BarColor = ui.ColorGreen

	memGauge := widgets.NewGauge()
	memGauge.Title = "Memory Usage"
	memGauge.SetRect(0, 5, 50, 10)
	memGauge.BarColor = ui.ColorGreen

	batGauge := widgets.NewGauge()
	batGauge.Title = "Battery Percentage"
	batGauge.SetRect(0, 10, 50, 15)
	batGauge.BarColor = ui.ColorGreen

	netParagraph := widgets.NewParagraph()
	netParagraph.Title = "Network I/O"
	netParagraph.SetRect(50, 0, 100, 15)

	// render loop
	uiEvents := ui.PollEvents()
	for {
		select {
		case s := <-ch:
			cpuGauge.Percent = int(s.CPU)
			memGauge.Percent = int(s.Memory)
			batGauge.Percent = int(s.Battery)
			netParagraph.Text = fmt.Sprintf("↑ %.3f MB/s\n↓ %.3f MB/s", s.NetSent, s.NetRecv)
			ui.Render(cpuGauge, memGauge, batGauge, netParagraph)
		case e := <-uiEvents:
			if e.ID == "q" || e.ID == "<C-c>" || e.ID == "<Escape>" {
				return
			}
		}
	}
}
