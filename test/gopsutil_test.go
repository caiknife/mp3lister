package test

import (
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
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
	partitions, err := disk.Partitions(false)
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

func TestGoPSUtil_Host(t *testing.T) {
	ti, err := host.BootTime()
	if err != nil {
		t.Error(err)
		return
	}
	unixTime := time.Unix(int64(ti), 0)
	t.Log(unixTime)
	t.Log(time.Since(unixTime))
}

func TestGoPSUtil_Host_Uptime(t *testing.T) {
	uptime, err := host.Uptime()
	if err != nil {
		t.Error()
		return
	}
	t.Log(time.Duration(uptime) * time.Second)
}
