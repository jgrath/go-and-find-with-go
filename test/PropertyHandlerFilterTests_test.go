package test

import (
	"github.com/jgrath/go-and-find-with-go/config"
	"github.com/jgrath/go-and-find-with-go/handlers"
	. "github.com/jgrath/go-and-find-with-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterSystemPropertiesByCode(t *testing.T) {

	systemProperty := SystemProperty{GroupCode: "code-x"}
	systemProperty1 := SystemProperty{GroupCode: "code-y"}
	systemProperty2 := SystemProperty{GroupCode: "code-z"}

	listOfSystemProperties := []*SystemProperty{&systemProperty, &systemProperty1, &systemProperty2}

	filteredList := handlers.FilterByCode(listOfSystemProperties, "code-y")

	assert.Equal(t, len(filteredList), 1)
	assert.Equal(t, filteredList[0], systemProperty1)
}

func TestFilterSystemPropertiesByCodeNoMatch(t *testing.T) {

	systemProperty := SystemProperty{GroupCode: "code-x"}
	systemProperty1 := SystemProperty{GroupCode: "code-y"}
	systemProperty2 := SystemProperty{GroupCode: "code-z"}

	listOfSystemProperties := []*SystemProperty{&systemProperty, &systemProperty1, &systemProperty2}

	filteredList := handlers.FilterByCode(listOfSystemProperties, "XXX")

	assert.Equal(t, len(filteredList), 0)
}

func TestFilterSystemPropertiesByGroup(t *testing.T) {

	systemProperty := SystemProperty{GroupName: "name-x"}
	systemProperty1 := SystemProperty{GroupName: "name-y"}
	systemProperty2 := SystemProperty{GroupName: "name-z"}

	listOfSystemProperties := []*SystemProperty{&systemProperty, &systemProperty1, &systemProperty2}

	filteredList := handlers.FilterByGroup(listOfSystemProperties, "name-y")

	assert.Equal(t, len(filteredList), 1)
	assert.Equal(t, filteredList[0], systemProperty1)
}

func TestFilterSystemPropertiesByGroupNoMatch(t *testing.T) {

	systemProperty := SystemProperty{GroupName: "name-x"}
	systemProperty1 := SystemProperty{GroupName: "name-y"}
	systemProperty2 := SystemProperty{GroupName: "name-z"}

	listOfSystemProperties := []*SystemProperty{&systemProperty, &systemProperty1, &systemProperty2}

	filteredList := handlers.FilterByGroup(listOfSystemProperties, "x")

	assert.Equal(t, len(filteredList), 0)
}

func TestFilterSystemPropertiesCombinded(t *testing.T) {

	systemProperty := SystemProperty{GroupName: "x"}
	systemProperty1 := SystemProperty{GroupName: "y"}
	systemProperty2 := SystemProperty{GroupCode: "x"}
	systemProperty3 := SystemProperty{GroupCode: "y"}

	listOfSystemProperties := []*SystemProperty{&systemProperty, &systemProperty1,
		&systemProperty2, &systemProperty3}

	filteredList := handlers.FilterSystemProperties(listOfSystemProperties, "y", true)

	assert.Equal(t, len(filteredList), 2)
	assert.Equal(t, filteredList[0], systemProperty1)
	assert.Equal(t, filteredList[1], systemProperty3)

	filteredListEmpty := handlers.FilterSystemProperties(listOfSystemProperties, "XXX", true)

	assert.Equal(t, len(filteredListEmpty), 0)
}

func TestFilterSystemPropertiesCombinedOnlyGroupInFilter(t *testing.T) {

	systemProperty := SystemProperty{GroupName: "x"}
	systemProperty1 := SystemProperty{GroupName: "y"}
	systemProperty2 := SystemProperty{GroupCode: "x"}
	systemProperty3 := SystemProperty{GroupCode: "y"}

	listOfSystemProperties := []*SystemProperty{&systemProperty, &systemProperty1,
		&systemProperty2, &systemProperty3}

	filteredList := handlers.FilterSystemProperties(listOfSystemProperties, "y", false)

	assert.Equal(t, len(filteredList), 1)
	assert.Equal(t, filteredList[0], systemProperty1)
}

func TestConfiguration(t *testing.T) {
	databaseConfiguration := config.GetConfiguration().DatabaseConfiguration
	assert.Equal(t, databaseConfiguration.Host, "localhost")
}
