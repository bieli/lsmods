package lsmods

import (
	"github.com/zcalusic/sysinfo"
)

type SysInfoGetter interface {
	GetSysInfo() *sysinfo.SysInfo
}

type SysInfoProvider struct{}

func GetSysInfoWrapper(sysInfoGetter SysInfoGetter) *sysinfo.SysInfo {
	sysInfo := sysInfoGetter.GetSysInfo()
	return sysInfo
}
func (s *SysInfoProvider) GetSysInfo() *sysinfo.SysInfo {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	return &si
}
