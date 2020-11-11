package main

import (
	"bufio"
	"debug/elf"
	"fmt"
	"github.com/zcalusic/sysinfo"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
)

func main() {
	log.Printf("Currently loaded kernel modules with descriptions:")
	var si sysinfo.SysInfo
	si.GetSysInfo()

	modInfo, err := NewModInfo(si)
	if err != nil {
		log.Fatal(err)
	}
	modules, err := modInfo.GetModInfo()
	if err != nil {
		log.Fatal(err)
	}

  sort.Sort(NameSorter(modules))

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

	for _, moduleInfo := range modules {
		fmt.Fprintf(writer, "%s\t%s\n", moduleInfo.Name, moduleInfo.Description)
	}
	writer.Flush()
}

const (
	libModulesPath                  = "/lib/modules/"
	procListModulesPath             = "/proc/modules"
	modulesListPath                 = "/modules.order"
	descriptionElfSymbolNamePattern = "__UNIQUE_ID_description"
	descriptionPattern              = `(description=)(.[^=]*)(author=|srcversion=|license=|alias=|depends=|vermagic=|filename=|name=|signature=|retpoline=|intree=|sig_id=|signer=|sig_key=|sig_hashalgo=)`
	descriptionPatternMatchIdx      = 2
)

// NameSorter sorts KernelModules struct by Name field.
type NameSorter []KernelModuleInfo

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }


type KernelModules map[string]string

type KernelModuleInfo struct {
	Name        string
	Description string
}

type ModInfo struct {
	LoadedModules    []KernelModuleInfo
	allKernelModules KernelModules
}

func NewModInfo(si sysinfo.SysInfo) (*ModInfo, error) {
	modInfo := &ModInfo{
		allKernelModules: make(KernelModules),
	}

	kernelModulesPaths, err := readAllKernelModules(si)
	if err != nil {
		return nil, err
	}

	for _, kernelModulePath := range kernelModulesPaths {
		_, fileName := filepath.Split(kernelModulePath)
		moduleName := strings.Replace(fileName, ".ko", "", 1)
		moduleFullPath := libModulesPath + si.Kernel.Release + "/" + kernelModulePath
		if len(moduleName) > 0 && len(moduleFullPath) > 0 {
			modInfo.allKernelModules[moduleName] = moduleFullPath
		}
	}
	return modInfo, nil
}

func (mi *ModInfo) GetModInfo() (modulesList []KernelModuleInfo, err error) {
	modules, err := readProcModules()

	for _, moduleName := range modules {
		// it's possible to find modules without description i.e. hid, parport, lp
		if libModulePath, ok := mi.allKernelModules[moduleName]; ok {
			moduleData, err := readModuleDescription(moduleName, libModulePath)
			if err != nil {
				return nil, err
			}
			modulesList = append(modulesList, moduleData)
		}
	}
	return modulesList, nil
}

func readModuleDescription(moduleName, libModulePath string) (kernelModInfo KernelModuleInfo, err error) {
	desc, err := getModuleDescriptionFromElf(libModulePath)
	if err != nil {
		return kernelModInfo, fmt.Errorf("[ERROR] Problem with get module '%s' description: %s", libModulePath, err)
	}

	kernelModInfo.Description = desc
	kernelModInfo.Name = moduleName
	return kernelModInfo, nil
}

func getFirstColumnFromTextFile(filePath string) (lines []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Problem with open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		splited := strings.Split(scanner.Text(), " ")
		if len(splited[0]) > 0 {
			lines = append(lines, splited[0])
		}
	}
	return lines, nil
}

func readProcModules() (lines []string, err error) {
	return getFirstColumnFromTextFile(procListModulesPath)
}

func readAllKernelModules(si sysinfo.SysInfo) (lines []string, err error) {
	return getFirstColumnFromTextFile(libModulesPath + si.Kernel.Release + modulesListPath)
}

func getModuleDescriptionFromElf(moduleFilePath string) (string, error) {
	fh, err := os.Open(moduleFilePath)
	if err != nil {
		return "", err
	}

	_elf, err := elf.NewFile(fh)
	if err != nil {
		return "", err
	}

	syms, err := _elf.Symbols()
	if err != nil {
		return "", err
	}

	for _, sym := range syms {
		if strings.Contains(sym.Name, descriptionElfSymbolNamePattern) {
			section := _elf.Sections[sym.Section]
			data, err := section.Data()
			if err != nil {
				return "", err
			}
			data = data[:]
			i := 0
			for ; data[i] != 0x00; i++ {
			}

			r, _ := regexp.Compile(descriptionPattern)
			for index, match := range r.FindStringSubmatch(string(data[:])) {
				if index == descriptionPatternMatchIdx {
					return match, nil
				}
			}
		}
	}
	return "", nil
}
