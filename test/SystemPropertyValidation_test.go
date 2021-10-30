package test

import (
	"github.com/jgrath/go-and-find-with-go/handlers"
	. "github.com/jgrath/go-and-find-with-go/util"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestInvalidGroupValidationsLower(t *testing.T) {
	var validations = []func(property *SystemProperty) (bool, string){handlers.IsGroupCodeValid}
	propertyUnderTest := DefaultSystemProperty("")
	propertyUnderTest.GroupCode = generateStringFromLength(4)
	validationResult := handlers.SystemPropertyValidation(&propertyUnderTest, validations)
	assert.Equal(t, validationResult, true)
}

func TestInvalidGroupValidationsUpper(t *testing.T) {
	var validations = []func(property *SystemProperty) (bool, string){handlers.IsGroupCodeValid}
	propertyUnderTest := DefaultSystemProperty("")
	propertyUnderTest.GroupCode = generateStringFromLength(11)
	validationResult := handlers.SystemPropertyValidation(&propertyUnderTest, validations)
	assert.Equal(t, validationResult, true)
}

func TestValidGroupValidationsLower(t *testing.T) {
	var validations = []func(property *SystemProperty) (bool, string){handlers.IsGroupCodeValid}
	propertyUnderTest := DefaultSystemProperty("")
	propertyUnderTest.GroupCode = generateStringFromLength(6)
	validationResult := handlers.SystemPropertyValidation(&propertyUnderTest, validations)
	assert.Equal(t, validationResult, false)
}

func TestValidGroupValidationsUpper(t *testing.T) {
	var validations = []func(property *SystemProperty) (bool, string){handlers.IsGroupCodeValid}
	propertyUnderTest := DefaultSystemProperty("")
	propertyUnderTest.GroupCode = generateStringFromLength(10)
	validationResult := handlers.SystemPropertyValidation(&propertyUnderTest, validations)
	assert.Equal(t, validationResult, false)
}

func TestInvalidDescriptionValidationsLower(t *testing.T) {
	var validations = []func(property *SystemProperty) (bool, string){handlers.IsGroupDescriptionValid}
	propertyUnderTest := DefaultSystemProperty("")
	propertyUnderTest.Description = generateStringFromLength(9)
	validationResult := handlers.SystemPropertyValidation(&propertyUnderTest, validations)
	assert.Equal(t, validationResult, true)
}

func TestValidDescriptionValidationsLower(t *testing.T) {
	var validations = []func(property *SystemProperty) (bool, string){handlers.IsGroupDescriptionValid}
	propertyUnderTest := DefaultSystemProperty("")
	propertyUnderTest.Description = generateStringFromLength(10)
	validationResult := handlers.SystemPropertyValidation(&propertyUnderTest, validations)
	assert.Equal(t, validationResult, false)
}

func TestValidDescriptionValidationsUpper(t *testing.T) {
	var validations = []func(property *SystemProperty) (bool, string){handlers.IsGroupDescriptionValid}
	propertyUnderTest := DefaultSystemProperty("")
	propertyUnderTest.Description = generateStringFromLength(20)
	validationResult := handlers.SystemPropertyValidation(&propertyUnderTest, validations)
	assert.Equal(t, validationResult, false)
}

func TestInValidDescriptionValidationsUpper(t *testing.T) {
	var validations = []func(property *SystemProperty) (bool, string){handlers.IsGroupDescriptionValid}
	propertyUnderTest := DefaultSystemProperty("")
	propertyUnderTest.Description = generateStringFromLength(21)
	validationResult := handlers.SystemPropertyValidation(&propertyUnderTest, validations)
	assert.Equal(t, validationResult, true)
}

func generateStringFromLength(length int) string {
	var seedChars = []rune("abcdefghijklmnopqrstuvwxyz!@££$%&*()")

	seed := make([]rune, length)

	for index := range seed {
		seed[index] = seedChars[rand.Intn(len(seed))]
	}
	return string(seed)
}
