package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kingpin"
	"github.com/bieli/lsmods/modinfo"
)

const (
	versionVarName = "LSMODS_VERSION"
)

var (
	version = "development"
	//configFile = kingpin.
	//		Flag("json", "Results in JSON format.").
	//		Short('j').
	//		Bool()
)

func main() {
	versionString := fmt.Sprintf("lsmos (%s version)", version)
	kingpin.Version(versionString)
	kingpin.Parse()

	log.Printf("Currently loaded kernel modules with descriptions:")
	modInfo, err := modinfo.NewModInfo()
	if err != nil {
		log.Fatal(err)
	}

	modules, err := modInfo.GetModInfo()
	if err != nil {
		log.Fatal(err)
	}
	for _, moduleInfo := range modules {
		fmt.Printf("%s\t\t%s\n", moduleInfo.Name, moduleInfo.Description)
	}

	//var mi modInfo.ModInfo

	//modulesList, err := si.GetModInfo()
	//if err != nil {
	//	log.Fatal(err)
	//}

	/*
	       if kingpin.json {
	   	    data, err := json.MarshalIndent(&si, "", "  ")
	   	    if err != nil {
	   		    log.Fatal(err)
	   	    }
	       	fmt.Println(string(data))
	   	} else {

	   	}
	*/
}
