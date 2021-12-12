package model

type UpInfo struct {
	CPUUsed     float64 `json:"CPUUsed"`
	IP          string  `gorm:"primaryKey"` //`gorm:"primary_key json:"ip"`
	MemUsed     float64 `json:"MemUsed"`
	ModelName   string  `json:"ModelName"`
	ISP         string  `json:"ISP"`
	CName       string  `json:"CName"`
	CCode       string  `json:"CCode"`
	Os          string  `json:"Os"`
	Platform    string  `json:"Platform"`
	UpdateTime  int64   `json:"UpdateTime"`
	Uptime      int     `json:"Uptime"`
	Flag        string  `json:"Flag"`
	SendTraffic string  `json:"SendTraffic"`
	RecvTraffic string  `json:"RecvTraffic"`
}
