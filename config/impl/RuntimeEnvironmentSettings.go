package impl

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type GoRuntimeSettingList map[string]string

func LoadDatabaseConfiguration(fileName string,
	databaseConfiguration *DatabaseConfiguration) error {

	configList, err := loadConfigFile(fileName)

	databaseConfiguration.Host = configList["Host"]
	atoi, _ := strconv.Atoi(configList["Port"])
	databaseConfiguration.Port = atoi
	databaseConfiguration.User = configList["User"]
	databaseConfiguration.Password = configList["Password"]
	databaseConfiguration.DBName = configList["DBName"]
	databaseConfiguration.SSLMode = configList["SSLMode"]
	databaseConfiguration.DBDriverName = configList["DBDriverName"]

	return err
}

func LoadLoggingConfiguration(fileName string,
	loggingConfiguration *LoggingConfiguration) error {
	configList, err := loadConfigFile(fileName)
	loggingConfiguration.Logfile = configList["Logfile"]
	loggingConfiguration.LogDir = configList["LogDir"]
	return err
}

func LoadApplicationConfiguration(fileName string,
	applicationConfiguration *ApplicationConfiguration) error {
	configList, err := loadConfigFile(fileName)
	applicationConfiguration.AppPort = configList["AppPort"]
	return err
}

func loadConfigFile(fileName string) (GoRuntimeSettingList, error) {
	const equalSymbol = "="
	goConfigSettingsArray := GoRuntimeSettingList{}

	configFile, err := os.Open(fileName)

	if err != nil {
		return nil, errors.New("Config file: [" + fileName + "] is missing")
	}

	defer configFile.Close()

	ioScanner := bufio.NewScanner(configFile)

	for ioScanner.Scan() {

		lineItem := ioScanner.Text()

		if isConfigValue := strings.Index(lineItem, equalSymbol); isConfigValue >= 0 {
			if configKey := lineItem[:isConfigValue]; len(configKey) > 0 {
				goConfigSettingsArray[configKey] = lineItem[isConfigValue+1:]
			}
		}
	}

	if len(goConfigSettingsArray) == 0 {
		return nil, errors.New("Config file: [" + fileName + "] is empty")
	}

	return goConfigSettingsArray, nil
}
