package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/super-l/machine-code/machine"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

func memStats(ctx *gin.Context) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, stat)
}

func cpuInfo(ctx *gin.Context) {
	info, err := cpu.Info()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	if len(info) == 0 {
		curd.Fail(ctx, "查询失败")
		return
	}
	curd.OK(ctx, info[0])
}

func cpuStats(ctx *gin.Context) {
	times, err := cpu.Times(false)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	if len(times) == 0 {
		curd.Fail(ctx, "查询失败")
		return
	}
	curd.OK(ctx, times[0])
}

func diskStats(ctx *gin.Context) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	usages := make([]*disk.UsageStat, 0)
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		usages = append(usages, usage)
	}
	curd.OK(ctx, usages)
}

func machineInfo(ctx *gin.Context) {
	info := machine.GetMachineData()
	curd.OK(ctx, gin.H{
		"sn":   info.SerialNumber,
		"mac":  info.Mac,
		"uuid": info.PlatformUUID,
		"cpu":  info.CpuId,
	})
}
