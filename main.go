package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type Specs struct {
	Cores      uint16
	Threads    uint16
	BaseClock  string
	BoostClock string
	TDP        string
}

var CPUs = map[string]map[string][]string{
	"Xeon": {
		"Sandy Bridge": {"Xeon E3", "Xeon E5", "Xeon E7"},
	},
}

type CPUSpecs map[string]Specs

var SandyBridge = map[string]CPUSpecs{
	"Xeon E3": {
		"E3-1220L": {2, 4, "2.20 GHz", "3.40 Ghz", "20W"},
		"E3-1220": {4, 4, "3.10 GHz", "3.40 Ghz", "80W"},
		"E3-1225": {4, 4, "3.10 GHz", "3.40 Ghz", "95W"},
		"E3-1230": {4, 8, "3.20 GHz", "3.60 Ghz", "80W"},
		"E3-1235": {4, 8, "3.20 GHz", "3.60 Ghz", "95W"},
		"E3-1240": {4, 8, "3.30 GHz", "3.70 Ghz", "80W"},
		"E3-1245": {4, 8, "3.30 GHz", "3.70 Ghz", "95W"},
		"E3-1270": {4, 8, "3.30 GHz", "3.80 Ghz", "80W"},
		"E3-1275": {4, 8, "3.30 GHz", "3.80 Ghz", "95W"},
		"E3-1280": {4, 8, "3.50 GHz", "3.90 Ghz", "95W"},
		"E3-1290": {4, 8, "3.60 GHz", "4.00 Ghz", "95W"},
		"E3-1260L": {4, 8, "2.40 GHz", "3.30 Ghz", "45W"},
		"E3-1265L": {4, 8, "2.40 GHz", "3.30 Ghz", "45W"},
	},
}

func ReturnCPU(model string, family string, data map[string]CPUSpecs) string {
	if cpuFamily, ok := data[family]; ok {
		if spec, ok := cpuFamily[model]; ok {
			return fmt.Sprintf(
				"Model Name: %s\nCores: %d\nThreads: %d\nBase Clock: %s\nBoost Clock: %s\nTDP: %s\n",
				model, spec.Cores, spec.Threads, spec.BaseClock, spec.BoostClock, spec.TDP,
			)
		}
	}
	return "CPU Specs not found"
}

func main() {
	var ChoosenArch, ChoosenCPU, ChoosenModel string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Sandy Bridge", "Ivy Bridge", "Haswell", "Broadwell", "Skylake")...).
				Title("Choose your microarchitecture").
				Value(&ChoosenArch),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Xeon E3", "Xeon E5", "Xeon E7")...).
				Title("Choose a CPU Family").
				Value(&ChoosenCPU),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions(
					"E3-1220", "E3-1220L", "E3-1225",
					"E3-1230", "E3-1235", "E3-1240", "E3-1245", "E3-1260L",
					"E3-1265L", "E3-1270", "E3-1275", "E3-1280", "E3-1290",
				)...).
				Title("Choose a CPU Model").
				Value(&ChoosenModel),
		),
	)
	form.Run()

	if ChoosenArch == "Sandy Bridge" {
		fmt.Println(ReturnCPU(ChoosenModel, ChoosenCPU, SandyBridge))
	} else {
		fmt.Println("Architecture not supported yet.")
	}
}
