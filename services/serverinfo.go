package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"../models"
	"../utils"
)

const (
	filename = "./data.json"
)

// GetServerInfo to return information about API
func GetServerInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	serverName := params["name"]

	JSONData, err := readJSONFile(filename)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	serverFiltered := gettingServerNameData(JSONData, serverName)
	if len(serverFiltered) == 0 {
		err := models.ErrNotFound{
			Message: "No info with this param.",
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	serverInfo := calculateServerInfo(serverFiltered)

	json.NewEncoder(w).Encode(serverInfo)
}

func gettingServerNameData(JSONData []models.JSONFile, serverName string) []models.JSONFile {
	serverFiltered := []models.JSONFile{}
	for _, value := range JSONData {
		if value.Hostname == serverName {
			serverFiltered = append(serverFiltered, value)
		}
	}
	return serverFiltered
}

func calculateServerInfo(serverArr []models.JSONFile) models.ServerInfo {
	occurrences := len(serverArr)
	serverName := serverArr[0].Hostname

	avarageCPU, avarageDisk, avarageUsage := calculateAvarage(serverArr)
	modeCPU, modeDisk, modeUsage := calculeteMode(serverArr)
	usageTrendCPU, usageTrendDisk, usageTrendUsage := calculeteUsageTrend(serverArr)

	avgCPU := fmt.Sprintf("%.2f", avarageCPU) + " %"
	trendCPU := fmt.Sprintf("%.2f", usageTrendCPU) + " %"

	CPUInfo := models.Info{
		Avarage:    avgCPU,
		Mode:       modeCPU,
		UsageTrend: trendCPU,
	}
	diskInfo := models.Info{
		Avarage:    fmt.Sprintf("%.2f GB", avarageDisk),
		Mode:       modeDisk,
		UsageTrend: fmt.Sprintf("%.2f GB", usageTrendDisk),
	}
	memoryInfo := models.Info{
		Avarage:    fmt.Sprintf("%.2f GB", avarageUsage),
		Mode:       modeUsage,
		UsageTrend: fmt.Sprintf("%.2f GB", usageTrendUsage),
	}

	serverInfo := models.ServerInfo{
		ServerName:  serverName,
		CPU:         CPUInfo,
		Disk:        diskInfo,
		Memory:      memoryInfo,
		Occurrences: occurrences,
	}

	return serverInfo
}

func calculateAvarage(serverArr []models.JSONFile) (avarageCPU, avarageDisk, avarageUsage float64) {
	var sumCPULoad, sumDiskUsage, sumMemoryUsage float64

	quantity := float64(len(serverArr))

	for _, value := range serverArr {
		sumCPULoad += value.CPULoad.Value
		sumDiskUsage += value.DiskUsage.Value
		sumMemoryUsage += value.MemoryUsage.Value
	}

	avarageCPU = sumCPULoad / float64(quantity)
	avarageDisk = sumDiskUsage / float64(quantity)
	avarageUsage = sumMemoryUsage / float64(quantity)

	return avarageCPU, avarageDisk, avarageUsage
}

func calculeteMode(serverArr []models.JSONFile) (modeCPU, modeDisk, modeMemory []float64) {
	countCPU, countDisk, countMemory := 0, 0, 0

	mapCPU, mapDisk, mapMemory := map[float64]int{}, map[float64]int{}, map[float64]int{}

	for _, iValue := range serverArr {
		for _, jValue := range serverArr {
			if iValue.CPULoad.Value == jValue.CPULoad.Value {
				countCPU++
				mapCPU[iValue.CPULoad.Value] = countCPU
			}

			if iValue.DiskUsage.Value == jValue.DiskUsage.Value {
				countDisk++
				mapDisk[iValue.DiskUsage.Value] = countDisk
			}

			if iValue.MemoryUsage.Value == jValue.MemoryUsage.Value {
				countMemory++
				mapMemory[iValue.MemoryUsage.Value] = countMemory
			}
		}
		countCPU, countDisk, countMemory = 0, 0, 0
	}

	indexCPU, validCPUMode := validadeMode(mapCPU)
	indexDisk, validDiskMode := validadeMode(mapDisk)
	indexMemory, validMemoryMode := validadeMode(mapMemory)

	modeCPU = gettingMode(mapCPU, indexCPU, validCPUMode)
	modeDisk = gettingMode(mapDisk, indexDisk, validDiskMode)
	modeMemory = gettingMode(mapMemory, indexMemory, validMemoryMode)

	return modeCPU, modeDisk, modeMemory
}

//func validadeMode(mapModes map[float64]int) []float64 {
func validadeMode(mapModes map[float64]int) (int, bool) {
	var check []int
	for _, v := range mapModes {
		ok, _ := utils.InArray(v, check)
		if !ok {
			check = append(check, v)
		}
	}

	if len(check) == 1 {
		return -1, false
	}

	max := utils.FindMaxValue(check)
	return max, true
}

func gettingMode(mapModes map[float64]int, index int, valid bool) (mode []float64) {
	if !valid {
		mode = append(mode, -1.0)
		return mode
	}

	for key, value := range mapModes {
		if index == value {
			mode = append(mode, key)
		}
	}

	return mode
}

func calculeteUsageTrend(serverArr []models.JSONFile) (usageTrendCPU, usageTrendDisk, usageTrendMemory float64) {
	firstValue := serverArr[0]
	var diffTrendCPU, diffTrendDisk, diffTrendMemory float64
	var count int

	for _, value := range serverArr[1:] {
		diffTrendCPU += ((value.CPULoad.Value / firstValue.CPULoad.Value) - 1)
		diffTrendDisk += ((value.DiskUsage.Value / firstValue.DiskUsage.Value) - 1)
		diffTrendMemory += ((value.MemoryUsage.Value / firstValue.MemoryUsage.Value) - 1)
		firstValue = value
		count++
	}

	usageTrendCPU = firstValue.CPULoad.Value * (diffTrendCPU/float64(count) + 1)
	usageTrendDisk = firstValue.DiskUsage.Value * (diffTrendDisk/float64(count) + 1)
	usageTrendMemory = firstValue.MemoryUsage.Value * (diffTrendMemory/float64(count) + 1)

	return usageTrendCPU, usageTrendDisk, usageTrendMemory
}
