package info

import (
	"beautyProject/internal/pkg/dto"
	"fmt"
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

type HardwareSpecsInfo struct {
}

func (r *HardwareSpecsInfo) Query() dto.Table {
	var cpuSpec, hostSpec, memorySpec, diskSpec string
	// 查詢 CPU 資訊
	cpuInfos, err := cpu.Info()
	if err != nil {
		log.Fatalf("取得 CPU 資訊失敗: %v", err)
	}
	log.Info("CPU 資訊:")
	for _, info := range cpuInfos {
		cpuSpec = fmt.Sprint("Model: ", info.ModelName, ", Cores: ", info.Cores)
	}
	// 查詢記憶體資訊
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("取得記憶體資訊失敗: %v", err)
	}
	totalGB := float64(vmStat.Total) / (1024 * 1024 * 1024)
	availableGB := float64(vmStat.Available) / (1024 * 1024 * 1024)
	memorySpec = fmt.Sprintf("\n記憶體資訊:\nTotal: %.2f GB, Available: %.2f GB, UsedPercent: %.2f%%\n",
		totalGB, availableGB, vmStat.UsedPercent)
	// 查詢硬碟資訊
	block, err := ghw.Block()
	if err != nil {
		log.Fatalf("無法取得磁碟資訊: %v", err)
	}

	var totalSize uint64 = 0
	for _, disk := range block.Disks {
		totalSize += disk.SizeBytes
	}
	totalDiskGB := float64(totalSize) / (1024 * 1024 * 1024)
	diskSpec = fmt.Sprintf("總硬碟容量: %.2f GB\n", totalDiskGB)

	// 查詢主機資訊
	hostStat, err := host.Info()
	if err != nil {
		log.Fatalf("取得主機資訊失敗: %v", err)
	}
	hostSpec = fmt.Sprintf("\n主機資訊:\nHostname: %s, OS: %s, Platform: %s",
		hostStat.Hostname, hostStat.OS, hostStat.Platform)
	dataArray := []map[string]any{
		{"cpu_info": cpuSpec, "host_info": hostSpec, "memory_info": memorySpec, "disk_info": diskSpec},
	}
	return dto.Table{Success: true, Message: "查詢成功", DataArray: dataArray, Total: len(dataArray)}
}
