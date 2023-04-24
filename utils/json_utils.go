package utils

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"os"
)

func GetConfigJsonFile(pathConfigFile, pathOfData string) (result string) {
	jsonFile, _ := os.Open(pathConfigFile)
	byteValue, _ := io.ReadAll(jsonFile)
	return gjson.Get(string(byteValue), pathOfData).String()
}

func GetJsonRead(pathOfFile string) (result []byte) {
	jsonFile, err := os.Open(pathOfFile)
	if err != nil {
		_ = fmt.Errorf("File not found in " + pathOfFile)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	return byteValue
}
