package test

import (
	"fmt"
	. "github.com/jgrath/go-and-find-with-go/util"
	"time"
)

func DefaultSystemProperty(defaultSetupValue string) SystemProperty {
	property := SystemProperty{}
	property.Name = "property-name"
	property.Value = "property-value" + defaultSetupValue
	property.Enabled = true
	parsedTime, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
	property.ActiveFromDate = parsedTime
	property.DefaultValue = ""
	property.Description = "default description value"
	return property
}

func DefaultSystemPropertyList(defaultSetupValue string, listSize int) []SystemProperty {

	populatedList := make([]SystemProperty, 0)

	for counter := 0; counter < listSize; counter++ {
		property := SystemProperty{}
		property.Name = "property-name"
		property.Value = "property-value" + defaultSetupValue
		property.Enabled = true
		parsedTime, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
		property.ActiveFromDate = parsedTime
		property.DefaultValue = ""
		property.Description = "default description value"
		property.GroupName = fmt.Sprint("group-name", counter)
		property.GroupCode = fmt.Sprint("group-code", counter)
		populatedList = append(populatedList, property)
	}
	return populatedList
}

