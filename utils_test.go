package lsmods

import (
	"os"

	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	expectedDescription = "Transport Layer Security Support"
	exampleTextFile     = "./test/assets/example.txt"
)

func TestUtilsGetModuleDescriptionFromElfSuccess(t *testing.T) {
	utils := &Utils{}

	resourcePath := os.Getenv("KMODULE")
	println("resourcePath:", resourcePath)
	result, _ := utils.GetModuleDescriptionFromElf(resourcePath)

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
