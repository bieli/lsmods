package lsmods

import (
	"github.com/zcalusic/sysinfo"
)

type SysInfoGetter interface {
	GetSysInfo() (*sysinfo.SysInfo, error)
}

type SysInfoProvider struct{}

func (s *SysInfoProvider) GetSysInfo() *sysinfo.SysInfo {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	return &si
}
