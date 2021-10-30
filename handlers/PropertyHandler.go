package handlers

import (
	"encoding/json"
	"fmt"
	. "github.com/jgrath/go-and-find-with-go/store"
	. "github.com/jgrath/go-and-find-with-go/util"
	"net/http"
	. "strings"
)

const pathSeparator = "/"

func FindSystemProperties(w http.ResponseWriter, r *http.Request) {

	listOfSystemProperties, err := MainPropertyStore.FindProperties()

	jsonBytes, err := json.Marshal(listOfSystemProperties)

	if err != nil {
		LogError.Println(fmt.Errorf("finding data - backend error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

func FindSystemPropertiesByCriteria(w http.ResponseWriter, r *http.Request) {

	splitString := Split(r.URL.Path, pathSeparator)

	searchTypeItem := splitString[2]
	searchKeyItem := splitString[3]

	listOfSystemProperties, err := MainPropertyStore.FindProperties()

	filteredElements := FilterSystemProperties(listOfSystemProperties, searchKeyItem, Contains(searchTypeItem, "group-code"))

	jsonBytes, err := json.Marshal(filteredElements)

	if err != nil {
		LogError.Println(fmt.Errorf("error finding data by criteria: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonBytes)
}

func AddSystemProperties(w http.ResponseWriter, r *http.Request) {

	systemPropertyToAdd := &SystemProperty{}

	json.NewDecoder(r.Body).Decode(systemPropertyToAdd)

	validationResult := SystemPropertyValidation(systemPropertyToAdd, DefaultSystemValidation())

	if validationResult {
		LogError.Println("adding property - validation failure")
		http.Error(w, "validation failure", http.StatusUnprocessableEntity)
		return
	}

	err := MainPropertyStore.AddProperty(systemPropertyToAdd)
	LogInfo.Println("property item added successfully")

	if err != nil {
		LogError.Println(fmt.Errorf("error persisting data: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DefaultSystemValidation() []func(property *SystemProperty) (bool, string) {
	return []func(property *SystemProperty) (bool, string){
		IsGroupCodeValid,
		IsGroupDescriptionValid,
		IsNameNameValid}
}

func FilterSystemProperties(listOfSystemProperties []*SystemProperty, searchKey string, includeCodesInFilter bool) []SystemProperty {

	searchByGroupSlice := FilterByGroup(listOfSystemProperties, searchKey)

	if includeCodesInFilter {
		searchByCodeSlice := FilterByCode(listOfSystemProperties, searchKey)
		for index, _ := range searchByCodeSlice {
			searchByGroupSlice = append(searchByGroupSlice, searchByCodeSlice[index])
		}
	}

	return searchByGroupSlice
}

func FilterByGroup(listOfSystemProperties []*SystemProperty, groupName string) []SystemProperty {
	filteredElements := []SystemProperty{}
	for index := range listOfSystemProperties {
		if listOfSystemProperties[index].GroupName == groupName {
			filteredElements = append(filteredElements, *listOfSystemProperties[index])
		}
	}
	return filteredElements
}

func FilterByCode(listOfSystemProperties []*SystemProperty, groupCode string) []SystemProperty {
	filteredElements := []SystemProperty{}
	for index := range listOfSystemProperties {
		if listOfSystemProperties[index].GroupCode == groupCode {
			filteredElements = append(filteredElements, *listOfSystemProperties[index])
		}
	}
	return filteredElements
}

func IsGroupCodeValid(property *SystemProperty) (bool, string) {
	if isPropertyStandardLength(property.GroupCode) {
		return false, "invalid code length"
	} else {
		return true, ""
	}
}

func IsNameNameValid(property *SystemProperty) (bool, string) {
	if isPropertyStandardLength(property.GroupCode) {
		return false, "invalid code length"
	} else {
		return true, ""
	}
}

func isPropertyStandardLength(value string) bool {
	return len(value) < 5 || len(value) > 10
}

func IsGroupDescriptionValid(property *SystemProperty) (bool, string) {
	if len(property.Description) < 10 || len(property.Description) > 20 {
		return false, "invalid description length"
	} else {
		return true, ""
	}
}

func SystemPropertyValidation(property *SystemProperty,
	validations []func(systemProperty *SystemProperty) (bool, string)) bool {

	containsValidationErrors := false

	for _, exec := range validations {
		res, _ := exec(property)
		if res == false {
			containsValidationErrors = true
			break
		}
	}
	return containsValidationErrors
}
