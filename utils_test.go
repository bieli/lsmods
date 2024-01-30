package lsmods

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	exampleKernelModulElfFIlePath = "test/assets/i2c-smbus.ko"
	exampleTextFile               = "test/assets/example.txt"
	expectedDescription           = "SMBus protocol extensions support"
)

func TestUtilsGetModuleDescriptionFromElfSuccess(t *testing.T) {
	utils := &Utils{}

	result, _ := utils.GetModuleDescriptionFromElf(exampleKernelModulElfFIlePath)

	assert.Equal(t, result, expectedDescription)
}

func TestUtilsGetModuleDescriptionFromElfFailWhenReadingNonExistsFile(t *testing.T) {
	utils := &Utils{}

	result, err := utils.GetModuleDescriptionFromElf("/exampleKernelModulElfFIlePath_NOT_EXISTS")

	assert.Equal(t, result, "")
	assert.NotNil(t, err)
}

func TestUtilsGetModuleDescriptionFromElfFailWhenNotElfFile(t *testing.T) {
	utils := &Utils{}

	result, err := utils.GetModuleDescriptionFromElf("Makefile")

	assert.Equal(t, result, "")
	assert.NotNil(t, err)
}

func TestGetFirstColumnFromTextFileSuccess(t *testing.T) {
	expectedResult := []string([]string{"first_column_first_line", "first_column_second_line"})
	utils := &Utils{}

	result, _ := utils.GetFirstColumnFromTextFile(exampleTextFile)

	assert.Equal(t, result, expectedResult)
}

func TestGetFirstColumnFromTextFileFailNotExistsFile(t *testing.T) {
	utils := &Utils{}

	_, err := utils.GetFirstColumnFromTextFile("/exampleTextFile_NOT_EXISTS")

	assert.NotNil(t, err)
}