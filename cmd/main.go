package main

import (
	"github.com/bieli/lsmods"
	"log"
)

/*
func GetSysInfoWrapper(sysInfoGetter SysInfoGetter) (*sysinfo.SysInfo, error) {
	sysInfo, err := sysInfoGetter.GetSysInfo()
	if err != nil {
		// handle error
		return &sysinfo.SysInfo{}, errors.New("problem with getting information from SySInfo module")
	}
	return sysInfo, nil
}*/

const (
	libModulesPath                  = "/lib/modules/"
	procListModulesPath             = "/proc/modules"
	modulesListPath                 = "/modules.order"
	descriptionElfSymbolNamePattern = "__UNIQUE_ID_description"
	descriptionPattern              = `(description=)(.[^=]*)(author=|srcversion=|license=|alias=|depends=|vermagic=|filename=|name=|signature=|retpoline=|intree=|sig_id=|signer=|sig_key=|sig_hashalgo=)`
	descriptionPatternMatchIdx      = 2
)

func main() {
	log.Printf("Currently loaded kernel modules with descriptions:")
	sip := &lsmods.SysInfoProvider{}
	si := sip.GetSysInfo()

	modInfo, err := lsmods.NewModInfo(*si, libModulesPath, modulesListPath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = modInfo.GetModInfo(true, procListModulesPath, libModulesPath, descriptionElfSymbolNamePattern, descriptionPattern, descriptionPatternMatchIdx)
	if err != nil {
		log.Fatal(err)
	}

	modInfo.Output()
}
