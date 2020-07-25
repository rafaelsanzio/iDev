package models

// ServerInfo model
type ServerInfo struct {
	ServerName  string `json:"serverName"`
	CPU         Info   `json:"CPU"`
	Memory      Info   `json:"memory"`
	Disk        Info   `json:"disk"`
	Occurrences int    `json:"occurrences"`
}

//Info model
type Info struct {
	Avarage    string    `json:"avarage"`
	Mode       []float64 `json:"mode"`
	UsageTrend string    `json:"usageTrend"`
}
