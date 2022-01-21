/*
 * Copyright (c) 2021 Siemens AG
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package utils

import (
	"os/exec"
	systemapi "systemservice/api/siemens_iedge_dmapi_v1"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

const (
	shell = "bash"
	//DefaultConfigPath is the absolute path for Limits and MonitoredStorage configuration
	DefaultConfigPath = "/opt/limits/default.json"
)

//Utils Interface has the wrapper of util calls
type Utils interface {
	Commander(command string) ([]byte, error)
	Uptime() (uint64, error)
	VirtualMemory() (*mem.VirtualMemoryStat, error)
	CPUPercent(interval time.Duration, percpu bool) ([]float64, error)
	CPUCounts(logical bool) (int, error)
	CPUInfo() ([]cpu.InfoStat, error)
}

//OsUtils struct for wrappers
type OsUtils struct{}

//Commander is a wrapper func
func (OsUtils) Commander(command string) ([]byte, error) {
	out, err := exec.Command(shell, "-c", command).Output()
	return out, err
}

//Uptime is a wrapper func
func (OsUtils) Uptime() (uint64, error) {
	return host.Uptime()
}

//VirtualMemory is a wrapper func
func (OsUtils) VirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

//CPUPercent is a wrapper func
func (OsUtils) CPUPercent(interval time.Duration, percpu bool) ([]float64, error) {
	return cpu.Percent(interval, percpu)
}

//CPUCounts is a wrapper func
func (OsUtils) CPUCounts(logical bool) (int, error) {
	return cpu.Counts(logical)
}

//CPUInfo is a wrapper func
func (OsUtils) CPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// DefaultConfig stores embedded objects inside default.json
type DefaultConfig struct {
	Limits           systemapi.Limits `json:"Limits"`
	MonitoredStorage struct {
		Path string `json:"Path"`
	} `json:"MonitoredStorage"`
}
