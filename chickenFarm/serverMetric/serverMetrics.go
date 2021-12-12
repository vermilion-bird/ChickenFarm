package serverMetric

import (
	"chickenFarm/requests"
	"fmt"
	"math"
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

func GetOutBoundIPByHost() (ip string) {
	ip = requests.Get("https://api.cncoder.cn/ip/v1")
	return
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

var LastSent, LastRecv uint64
var NetSpeedRecv, NetSpeedSent string
var previousTime uint64

func GetNetSpeed() {
	nInfo := GetNetInfo()
	var sentTotal, recvToal uint64
	for _, v := range nInfo {
		if v.Name != "lo" {
			sentTotal += v.BytesSent
			recvToal += v.BytesRecv
		}

	}
	sunit, runit := "bytes/s", "bytes/s"
	var sentSpeed, recvSpeed uint64
	now := uint64(time.Now().Unix())
	timeDelta := now - previousTime
	if timeDelta == 0 { //除0错误
		timeDelta = 1
	}
	sentSpeed = uint64(sentTotal-LastSent) / timeDelta
	interval := uint64(math.Pow(2, 10))
	if sentSpeed > 1024 {
		sunit = "kb/s"
		sentSpeed = sentSpeed / interval
		if sentSpeed > 1024 {
			sentSpeed = sentSpeed / interval
			sunit = "mb/s"
		}
	}
	recvSpeed = uint64(recvToal-LastRecv) / timeDelta
	if recvSpeed > 1024 {
		runit = "kb/s"
		recvSpeed = recvSpeed / interval
		if recvSpeed > 1024 {
			recvSpeed = recvSpeed / interval
			runit = "mb/s"
		}
	}
	NetSpeedRecv = fmt.Sprintf("%d %s", recvSpeed, runit)
	NetSpeedSent = fmt.Sprintf("%d %s", sentSpeed, sunit)
	LastSent = sentTotal
	LastRecv = recvToal
	previousTime = now
}
