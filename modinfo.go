package lsmods

import (
	"fmt"
	"github.com/zcalusic/sysinfo"
	"os"
	"sort"
	"text/tabwriter"
)

const (
	modulesListPath                 = "/modules.order"
	descriptionElfSymbolNamePattern = "__UNIQUE_ID_description"
	descriptionPattern              = `(description=)(.[^=]*)(author=|srcversion=|license=|alias=|depends=|vermagic=|filename=|name=|signature=|retpoline=|intree=|sig_id=|signer=|sig_key=|sig_hashalgo=)`
	descriptionPatternMatchIdx      = 2
)

type KernelModules map[string]string

type KernelModuleInfo struct {
	Name        string
	Description string
}

type ModInfo struct {
	LoadedModules    []KernelModuleInfo
	allKernelModules KernelModules
	utils            *Utils
}

func NewModInfo(si sysinfo.SysInfo, libModulesPath string) (*ModInfo, error) {
	modInfo := &ModInfo{
		allKernelModules: make(KernelModules),
		utils:            &Utils{},
	}

	kernelModulesPaths, err := modInfo.utils.ReadAllKernelModules(si, libModulesPath, modulesListPath)
	if err != nil {
		return nil, err
	}
	return modInfo.utils.PrepareAllKernelModulesList(libModulesPath, kernelModulesPaths, si, modInfo), nil
}

func (mi *ModInfo) GetModInfo(sortByName bool, procListModulesPath string) (modulesList []KernelModuleInfo, err error) {
	modules, err := mi.utils.ReadProcModules(procListModulesPath)

	for _, moduleName := range modules {
		// it's possible to find modules without description i.e. hid, parport, lp
		if libModulePath, ok := mi.allKernelModules[moduleName]; ok {
			moduleData, err := mi.readModuleDescription(moduleName, libModulePath)
			if err != nil {
				return nil, err
			}
			modulesList = append(modulesList, moduleData)
		}
	}

	if sortByName {
		sort.Sort(NameSorter(modulesList))
	}
	mi.LoadedModules = modulesList
	return modulesList, nil
}

func (mi *ModInfo) readModuleDescription(moduleName, libModulePath string) (kernelModInfo KernelModuleInfo, err error) {
	desc, err := mi.utils.GetModuleDescriptionFromElf(libModulePath, descriptionElfSymbolNamePattern, descriptionPattern, descriptionPatternMatchIdx)
	if err != nil {
		return kernelModInfo, fmt.Errorf("[ERROR] Problem with get module '%s' description: %s", libModulePath, err)
	}

	kernelModInfo.Description = desc
	kernelModInfo.Name = moduleName
	return kernelModInfo, nil
}

func (mi *ModInfo) Output() {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	for _, moduleInfo := range mi.LoadedModules {
		fmt.Fprintf(writer, "%s\t%s\t\n", moduleInfo.Name, moduleInfo.Description)
	}
	writer.Flush()
}
