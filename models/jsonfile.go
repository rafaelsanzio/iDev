package models

//JSONFile model
type JSONFile struct {
	Hostname    string `json:"hostname"`
	CPULoad     Values `json:"cpu_load"`
	MemorySize  Values `json:"memory_size"`
	MemoryUsage Values `json:"memory_usage"`
	DiskSize    Values `json:"disk_size"`
	DiskUsage   Values `json:"disk_usage"`
}

//Values model
type Values struct {
	Value float64 `json:"Value"`
	Unit  string  `json:"Unit"`
}
