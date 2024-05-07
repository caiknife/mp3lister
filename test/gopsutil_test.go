package test

import (
	"testing"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func TestGoPSUtil_Mem(t *testing.T) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(memory)
}

func TestGoPSUtil_CPU(t *testing.T) {
	info, err := cpu.Info()
	if err != nil {
		t.Error(err)
		return
	}

	for _, stat := range info {
		t.Log(stat)
	}
}

func TestGoPSUtil_Disk(t *testing.T) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		t.Error(err)
		return
	}

	for _, partition := range partitions {
		t.Log(partition)
	}
}

func TestGoPSUtil_Usage(t *testing.T) {
	usage, err := disk.Usage("/Users/caiknife/Desktop")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(usage.UsedPercent)
}
