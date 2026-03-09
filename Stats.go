package main

import (
	"time"

	"github.com/distatus/battery"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type Stats struct {
	CPU         float64
	CoreCPU     []float64
	Memory      float64
	Battery     float64
	NetSent     float64
	NetRecv     float64
	PrevNetSent uint64
	PrevNetRecv uint64
}

func (s *Stats) UpdateCPU() error {
	all, err := cpu.Percent(time.Second, true)
	if err != nil {
		return err
	}
	// average
	var sum float64
	for _, v := range all {
		sum += v
	}
	s.CPU = sum / float64(len(all))
	s.CoreCPU = all
	return nil
}
func (s *Stats) UpdateMemory() error {
	percentMemory, err := mem.VirtualMemory()
	if err != nil {
		return err
	}
	s.Memory = percentMemory.UsedPercent
	return nil
}
func (s *Stats) UpdateBattery() error {
	bat, err := battery.GetAll()
	if err != nil {
		return err
	}
	percentBat := bat[0].Current / bat[0].Full * 100
	s.Battery = percentBat
	return nil
}

func (s *Stats) UpdateNetwork() error {
	netStats, err := net.IOCounters(false)
	if err != nil {
		return err
	}
	sent := netStats[0].BytesSent
	recv := netStats[0].BytesRecv
	s.NetSent = (float64(sent) - float64(s.PrevNetSent)) / 1_000_000.0 // MB/s
	s.NetRecv = (float64(recv) - float64(s.PrevNetRecv)) / 1_000_000.0
	s.PrevNetSent = sent
	s.PrevNetRecv = recv
	return nil
}
func pollStats(ch chan<- Stats) {
	s := Stats{}
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		s.UpdateCPU()
		s.UpdateMemory()
		s.UpdateBattery()
		s.UpdateNetwork()
		ch <- s // send updated stats
	}
}
