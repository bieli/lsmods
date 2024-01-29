package lsmods

import (
	"github.com/zcalusic/sysinfo"
)

type MockSysInfoGetter struct{}

func (m *MockSysInfoGetter) GetSysInfo() (*sysinfo.SysInfo, error) {
	// Provide mock implementation for testing
	return &sysinfo.SysInfo{
		//Meta:    &sysinfo.Meta{},
		//Node:    &sysinfo.Node{},
		//OS:      &sysinfo.OS{},
		//Kernel:  &sysinfo.Kernel{},
		//Product: &sysinfo.Product{},
		//Board:   &sysinfo.Board{},
		//Chassis: &sysinfo.Chassis{},
		//BIOS:    &sysinfo.BIOS{},
		//CPU:     &sysinfo.CPU{},
		Memory: sysinfo.Memory{
			Type:  "RAM",
			Speed: 333,
			Size:  128000000,
		},
		Storage: []sysinfo.StorageDevice{},
		Network: []sysinfo.NetworkDevice{},
	}, nil
}

/*
func TestSomeFunction(t *testing.T) {
	mockSysInfoGetter := &MockSysInfoGetter{}

	// TODO
	result := GetSysInfoWrapper(mockSysInfoGetter)

	assert.NotNil(t, result)
}
*/
