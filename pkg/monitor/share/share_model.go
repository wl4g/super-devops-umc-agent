/**
 * Copyright 2017 ~ 2025 the original author or authors[983708408@qq.com].
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package share

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// PhysicalId
var PhysicalId string = "UNKNOWN"

type Total struct {
	Id          string                 `json:"physicalId"`
	Type        string                 `json:"type"`
	Mem         *mem.VirtualMemoryStat `json:"memInfo"`
	Cpu         []float64              `json:"cpu"`
	DiskInfos   []DiskInfo             `json:"diskInfos"`
	NetInfos    []NetInfo              `json:"netInfos"`
	DockerInfos []DockerInfo           `json:"dockerInfos"`
}

type DiskInfo struct {
	PartitionStat disk.PartitionStat `json:"partitionStat"`
	Usage         disk.UsageStat     `json:"usage"`
}

type NetInfo struct {
	Port      int `json:"port"`
	Up        int `json:"up"`
	Down      int `json:"down"`
	Count     int `json:"count"`
	Estab     int `json:"estab"`
	CloseWait int `json:"closeWait"`
	TimeWait  int `json:"timeWait"`
	Close     int `json:"close"`
	Listen    int `json:"listen"`
	Closing   int `json:"closing"`
}

type DockerInfo struct {
	ContainerId string `json:"containerId"`
	Name        string `json:"name"`
	CpuPerc     string `json:"cpuPerc"`
	MemUsage    string `json:"memUsage"`
	MemPerc     string `json:"memPerc"`
	NetIO       string `json:"netIO"`
	BlockIO     string `json:"blockIO"`
	PIDs        string `json:"PIDs"`
}
