package impl

import (
	"path"
	"runtime"
)

type DatabaseConfiguration struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	SSLMode      string
	DBDriverName string
}

type LoggingConfiguration struct {
	Logfile string
	LogDir  string
}

type ApplicationConfiguration struct {
	AppPort string
}

type Configuration struct {
	DatabaseConfiguration
	LoggingConfiguration
	ApplicationConfiguration
}

const databaseFileName = "/database_config.properties"
const loggingFileName = "/logging_config.properties"
const applicationFileName = "/application_config.properties"
const defaultLogConfigurationDirectory = "../"

func GetConfigurationWithDirectory(directory string) *Configuration {

	databaseConfiguration := DatabaseConfiguration{}
	loggingConfiguration := LoggingConfiguration{}
	applicationConfiguration := ApplicationConfiguration{}
	_, filename, _, _ := runtime.Caller(1)

	dbPath := path.Join(path.Dir(filename), directory+databaseFileName)
	dbError := LoadDatabaseConfiguration(dbPath, &databaseConfiguration)
	if dbError != nil {
		panic(dbError)
	}

	logPath := path.Join(path.Dir(filename), directory+loggingFileName)
	logFileError := LoadLoggingConfiguration(logPath, &loggingConfiguration)
	if logFileError != nil {
		panic(logFileError)
	}

	applicationPath := path.Join(path.Dir(filename), directory+applicationFileName)
	applicationError := LoadApplicationConfiguration(applicationPath, &applicationConfiguration)
	if applicationError != nil {
		panic(logFileError)
	}

	configuration := new(Configuration)
	configuration.DatabaseConfiguration = databaseConfiguration
	configuration.LoggingConfiguration = loggingConfiguration
	configuration.ApplicationConfiguration = applicationConfiguration

	return configuration
}

func GetConfiguration() *Configuration {
	return GetConfigurationWithDirectory(defaultLogConfigurationDirectory)
}
