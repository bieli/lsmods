package lsmods

import (
	"bufio"
	"debug/elf"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	modulesListPath                 = "/modules.order"
	descriptionElfSymbolNamePattern = "__UNIQUE_ID_description"
	descriptionPattern              = `(description=)(.[^=]*)(author=|srcversion=|license=|alias=|depends=|vermagic=|filename=|name=|signature=|retpoline=|intree=|sig_id=|signer=|sig_key=|sig_hashalgo=)`
	descriptionPatternMatchIdx      = 2
)

type Utils struct{}

func (u *Utils) PrepareAllKernelModulesList(libModulesPath string, kernelModulesPaths []string, kernelRelease string, modInfo *ModInfo) *ModInfo {
	for _, kernelModulePath := range kernelModulesPaths {
		_, fileName := filepath.Split(kernelModulePath)
		moduleName := strings.Replace(fileName, ".ko", "", 1)
		moduleFullPath := libModulesPath + kernelRelease + "/" + kernelModulePath
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

func (u *Utils) ReadAllKernelModules(libModulesPath string, kernelRelease string, modulesListPath string) (lines []string, err error) {
	return u.GetFirstColumnFromTextFile(libModulesPath + kernelRelease + modulesListPath)
}

func (u *Utils) GetModuleDescriptionFromElf(moduleFilePath string) (string, error) {
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
