package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/zgwit/iot-master/protocols"
)

func systemRoutes(app *gin.RouterGroup) {
	app.GET("version")
	//app.GET("cpu/info", cpuInfo)
	app.GET("cpu", cpuTimes)
	app.GET("memory", memoryInfo)
	app.GET("disk", diskInfo)
	app.GET("cron")
	app.GET("protocols", protocolList)

}

func memoryInfo(ctx *gin.Context)  {
	stat, err := mem.VirtualMemory()
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, stat)
}

func cpuInfo(ctx *gin.Context)  {
	info, err := cpu.Info()
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, info)
}

func cpuTimes(ctx *gin.Context)  {
	times, err := cpu.Times(false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if len(times) == 0 {
		replyFail(ctx, "查询失败")
		return
	}
	replyOk(ctx, times[0])
}

func diskInfo(ctx *gin.Context)  {
	partitions, err := disk.Partitions(false)
	if err != nil {
		replyError(ctx, err)
		return
	}

	usages := make([]*disk.UsageStat, 0)
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			replyError(ctx, err)
			return
		}
		usages = append(usages, usage)
	}
	replyOk(ctx, usages)
}


func protocolList(ctx *gin.Context) {
	ps := protocols.Protocols()
	replyOk(ctx, ps)
}