package main

import (
	"chickenFarm/serverMetric"
	"fmt"
)

func runMetrics() {
	// serverMetric.GetCpuInfo()
	info := serverMetric.GetCpuLoad()
	fmt.Println(info)
	cpuInfo, percent := serverMetric.GetCpuInfo()
	fmt.Println(cpuInfo, percent)

	memInfo := serverMetric.GetMemInfo()
	fmt.Printf("mem info:%v\n", memInfo)

	hInfo := serverMetric.GetHostInfo()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)

	nInfo := serverMetric.GetNetInfo()
	for index, v := range nInfo {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}

	ip, _ := serverMetric.GetOutBoundIP()
	fmt.Println(ip)

	ipd := serverMetric.GetOutBoundIPDirect()
	fmt.Println(ipd)

}
