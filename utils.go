package lsmods

import (
	"bufio"
	"debug/elf"
	"fmt"
	"github.com/zcalusic/sysinfo"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Utils struct{}

func (u *Utils) PrepareAllKernelModulesList(libModulesPath string, kernelModulesPaths []string, si sysinfo.SysInfo, modInfo *ModInfo) *ModInfo {
	for _, kernelModulePath := range kernelModulesPaths {
		_, fileName := filepath.Split(kernelModulePath)
		moduleName := strings.Replace(fileName, ".ko", "", 1)
		moduleFullPath := libModulesPath + si.Kernel.Release + "/" + kernelModulePath
		if len(moduleName) > 0 && len(moduleFullPath) > 0 {
			modInfo.allKernelModules[moduleName] = moduleFullPath
		}
	}
	return modInfo
}

func (u *Utils) GetFirstColumnFromTextFile(filePath string) (lines []string, err error) {
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

func (u *Utils) ReadProcModules(procListModulesPath string) (lines []string, err error) {
	return u.GetFirstColumnFromTextFile(procListModulesPath)
}

func (u *Utils) ReadAllKernelModules(si sysinfo.SysInfo, libModulesPath string, modulesListPath string) (lines []string, err error) {
	return u.GetFirstColumnFromTextFile(libModulesPath + si.Kernel.Release + modulesListPath)
}

func (u *Utils) GetModuleDescriptionFromElf(moduleFilePath string, descriptionElfSymbolNamePattern string, descriptionPattern string, descriptionPatternMatchIdx int) (string, error) {
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
			r, _ := regexp.Compile(descriptionPattern)
			for index, match := range r.FindStringSubmatch(string(data[:])) {
				if index == descriptionPatternMatchIdx {
					return match[:len(match)-1], nil
				}
			}
		}
	}
	return "", nil
}
