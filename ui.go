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

	// Left side
	memGauge := widgets.NewGauge()
	memGauge.Title = "Memory Usage"
	memGauge.SetRect(0, 0, 50, 5)
	memGauge.BarColor = ui.ColorGreen

	batGauge := widgets.NewGauge()
	batGauge.Title = "Battery"
	batGauge.SetRect(0, 5, 50, 10)
	batGauge.BarColor = ui.ColorGreen

	netParagraph := widgets.NewParagraph()
	netParagraph.Title = "Network I/O"
	netParagraph.SetRect(0, 10, 50, 16)

	// Right side - average CPU at top
	cpuAvgGauge := widgets.NewGauge()
	cpuAvgGauge.Title = "CPU Average"
	cpuAvgGauge.SetRect(50, 0, 100, 5)
	cpuAvgGauge.BarColor = ui.ColorCyan

	// Per-core gauges
	// Per-core gauges - 2 columns of 6
	numCores := 12
	coreGauges := make([]*widgets.Gauge, numCores)
	for i := 0; i < numCores; i++ {
		g := widgets.NewGauge()
		g.Label = fmt.Sprintf("Core %d", i+1)
		//g.Border = false

		col := i / 6 // 0 for cores 1-6, 1 for cores 7-12
		row := i % 6 // 0-5 within each column
		x1 := 50 + col*25
		x2 := x1 + 25
		y1 := 5 + row*3
		y2 := y1 + 3

		g.SetRect(x1, y1, x2, y2)
		g.BarColor = ui.ColorBlue
		coreGauges[i] = g
	}

	// Build render list
	renderItems := []ui.Drawable{memGauge, batGauge, netParagraph, cpuAvgGauge}
	for _, g := range coreGauges {
		renderItems = append(renderItems, g)
	}

	uiEvents := ui.PollEvents()
	for {
		select {
		case s := <-ch:
			memGauge.Percent = int(s.Memory)
			batGauge.Percent = int(s.Battery)
			cpuAvgGauge.Percent = int(s.CPU)
			netParagraph.Text = fmt.Sprintf("↑ %.3f MB/s\n↓ %.3f MB/s", s.NetSent, s.NetRecv)
			for i := 0; i < numCores && i < len(s.CoreCPU); i++ {
				coreGauges[i].Percent = int(s.CoreCPU[i])
				coreGauges[i].Label = fmt.Sprintf("Core %d: %d%%", i+1, int(s.CoreCPU[i]))
			}
			ui.Render(renderItems...)
		case e := <-uiEvents:
			if e.ID == "q" || e.ID == "<C-c>" || e.ID == "<Escape>" {
				return
			}
		}
	}
}
