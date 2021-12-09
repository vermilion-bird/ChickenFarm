package serverMetric

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	psnet "github.com/shirou/gopsutil/net"
)

func GetNetInfo() []psnet.IOCountersStat {
	info, _ := psnet.IOCounters(true)
	return info
	// for index, v := range info {
	// 	fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	// }
}

// disk info
func GetDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}
func GetHostInfo() *host.InfoStat {
	hInfo, _ := host.Info()
	return hInfo
	// fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

func GetMemInfo() *mem.VirtualMemoryStat {
	memInfo, _ := mem.VirtualMemory()
	return memInfo
	// fmt.Printf("mem info:%v\n", memInfo)
}

// cpu info
func GetCpuInfo() ([]cpu.InfoStat, []float64) {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	// for _, ci := range cpuInfos {
	// 	fmt.Println(ci)
	// }
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false)
	return cpuInfos, percent
	// for {
	// 	percent, _ := cpu.Percent(time.Second, false)
	// 	fmt.Printf("cpu percent:%v\n", percent)
	// }

}

//cpu load
func GetCpuLoad() *load.AvgStat {
	info, _ := load.Avg()
	return info
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// 直接获取本地地址
func GetOutBoundIPDirect() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				fmt.Println(ipnet.IP.String())
				return
			}
		}
	}
	return
}
