package modinfo

import (
	"bufio"
	"debug/elf"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zcalusic/sysinfo"
)

const (
	libModulesPath      = "/lib/modules/"
	procListModulesPath = "/proc/modules"
	modulesListPath     = "/modules.order"
)

type KernelModules map[string]string

type KernelModuleInfo struct {
	Name        string
	Description string
}

type ModInfo struct {
	LoadedModules    []KernelModuleInfo
	allKernelModules KernelModules
}

func NewModInfo() (*ModInfo, error) {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	modInfo := &ModInfo{}
	modInfo.allKernelModules = make(KernelModules)

	kernelModulesPaths, err := readAllKernelModules()
	if err != nil {
		return nil, err
	}

	for _, kernelModulePath := range kernelModulesPaths {
		_, fileName := filepath.Split(kernelModulePath)
		moduleName := strings.Replace(fileName, ".ko", "", 0)
		modInfo.allKernelModules[moduleName] = kernelModulePath
	}
	return modInfo, nil
}

func (mi *ModInfo) GetModInfo() (modulesList []KernelModuleInfo, err error) {
	modules, err := readProcModules()

	for _, moduleName := range modules {
		libModulePath := mi.allKernelModules[moduleName]
		moduleData, err := readModuleDescription(moduleName, libModulePath)
		if err != nil {
			return nil, err
		}
		modulesList = append(modulesList, moduleData)
	}
	return modulesList, nil
}

func readModuleDescription(moduleName, libModulePath string) (kernelModInfo KernelModuleInfo, err error) {
	/*
		cmd := exec.Command(prepareModuleDescriptionCommand(libModulePath))
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return kernelModInfo, err
		}
		if err := cmd.Start(); err != nil {
			return kernelModInfo, err
		}
		desc, err := ioutil.ReadAll(stdout)
		if err != nil {
			return kernelModInfo, err
		}
		kernelModInfo.Description = string(desc[:])
		kernelModInfo.Name = moduleName
	*/
	desc, err := getModuleDescriptionFromElf(libModulePath)
	kernelModInfo.Description = desc
	kernelModInfo.Name = moduleName
	return kernelModInfo, nil
}

func getFirstColumnFromTextFile(filePath string) (lines []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Problem with open file %s", err)
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

func readAllKernelModules() (lines []string, err error) {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	return getFirstColumnFromTextFile(libModulesPath + si.Kernel.Release + modulesListPath)
}

/*
func prepareModuleDescriptionCommand(libModulePath string) string {
	return "cat " + libModulesPath + "$(uname -r)/" + libModulePath +
		"| strings " +
		"| grep 'description=' " +
		"| awk 'BEGIN{FS=\"description=\"} END {print $2;}'"
}
*/
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

	startDesc := 34 + len("description=")

	for _, sym := range syms {
		if strings.Contains(sym.Name, "__UNIQUE_ID_description") {
			section := _elf.Sections[sym.Section]
			data, err := section.Data()
			if err != nil {
				return "", err
			}
			data = data[startDesc:]
			i := 0
			for ; data[i] != 0x00; i++ {
			}
			return string(data[:i]), nil
		}
	}

	return "", fmt.Errorf("no description found")
}
