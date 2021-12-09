package main

import (
	"chickenFarm/dataStruct"
	"chickenFarm/model"
	"chickenFarm/requests"
	"chickenFarm/serverMetric"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/robfig/cron"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func cronTask() {
	c := cron.New()
	spec := "*/10 * * * * ?"
	c.AddFunc(spec, func() {
		now := time.Now()
		fmt.Println("cron running:", now.Minute(), now.Second())
		upInfo()
	})
	c.Start()
	select {}
}

func main() {
	cronTask()
}
func upInfo() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	info := getMetrics()
	_, err = conn.Write(info)
	checkError(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	// os.Exit(0)
}

func getIPInfo(ip string) (infoMap map[string]interface{}) {
	url := "https://api.iplocation.net/?ip=" + ip
	resp := requests.Get(url)
	infoMap, _ = dataStruct.JsonToMap(resp)
	fmt.Println(resp)
	return infoMap
}

func getMetrics() []byte {
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

	ipd, _ := serverMetric.GetOutBoundIP()

	up := model.UpInfo{}
	up.IP = ipd
	up.ModelName = cpuInfo[0].ModelName
	up.Os = hInfo.OS
	up.Platform = hInfo.Platform
	up.Uptime = int(hInfo.Uptime)
	up.CPUUsed = percent[0]
	up.MemUsed = memInfo.UsedPercent
	up.UpdateTime = time.Now().Unix()
	ipInfo := getIPInfo(ipd)
	up.CCode = ipInfo["country_code2"].(string)
	up.CName = ipInfo["country_name"].(string)
	tmpdata := dataStruct.Struct2Map(up)
	str, _ := json.Marshal(tmpdata)
	return str
}
