package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatFile(fileName string, value float64) {
	valueTxt := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueTxt), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to read the file.")
	}

	valueTxt := string(data)
	value, err := strconv.ParseFloat(valueTxt, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse stored value.")
	}

	return value, nil
}
