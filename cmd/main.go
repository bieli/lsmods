package main

import (
	"github.com/bieli/lsmods"
	"log"
)

const (
	libModulesPath       = "/lib/modules/"
	procListModulesPath  = "/proc/modules"
	sortAllMOdulesByName = true
)

func main() {
	log.Printf("Currently loaded kernel modules with descriptions:")

	provider := &lsmods.SysInfoProvider{}
	var sysInfo = lsmods.GetSysInfoWrapper(provider)
	modInfo, err := lsmods.NewModInfo(*sysInfo, libModulesPath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = modInfo.GetModInfo(sortAllMOdulesByName, procListModulesPath)
	if err != nil {
		log.Fatal(err)
	}

	modInfo.Output()
}
