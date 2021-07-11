package test

import (
	"bytes"
	"encoding/json"
	"github.com/jgrath/go-and-find-with-go/handlers"
	"github.com/jgrath/go-and-find-with-go/test/mocks"
	. "github.com/jgrath/go-and-find-with-go/testutil"
	. "github.com/jgrath/go-and-find-with-go/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindPropertyHandlerWithDataFound(t *testing.T) {
	mockStore := mocks.InitPropertyMockStore()

	property := DefaultSystemProperty("")

	mockReturnValue := []*SystemProperty{&property}

	mockStore.On("FindProperties").Return(mockReturnValue, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	ThrowErrorWhenNull(t, err)
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.FindSystemProperties)

	handler.ServeHTTP(recorder, req)

	Assert200(t, recorder)

	returnRef := []SystemProperty{}
	err = json.NewDecoder(recorder.Body).Decode(&returnRef)

	ThrowErrorWhenNull(t, err)

	valueFromService := returnRef[0]

	AssertTestResultOutcome(t, valueFromService, property)

	mockStore.AssertExpectations(t)
	mockStore.AssertNumberOfCalls(t, "FindProperties", 1)
}

func TestFindPropertyHandlerWithDataFoundByGroupName(t *testing.T) {
	mockStore := mocks.InitPropertyMockStore()

	propertyList := DefaultSystemPropertyList("", 3)

	mockReturnValue := []*SystemProperty{&propertyList[0], &propertyList[1], &propertyList[2]}

	mockStore.On("FindProperties").Return(mockReturnValue, nil).Once()

	req, err := http.NewRequest("GET", "/system-settings/group-name/group-name1", nil)
	ThrowErrorWhenNull(t, err)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.FindSystemPropertiesByCriteria)

	handler.ServeHTTP(recorder, req)

	Assert200(t, recorder)

	returnRef := []SystemProperty{propertyList[1]}
	err = json.NewDecoder(recorder.Body).Decode(&returnRef)

	ThrowErrorWhenNull(t, err)

	valueFromService := returnRef[0]

	AssertTestResultOutcome(t, valueFromService, propertyList[1])
}

func TestFindPropertyHandlerWithDataFoundByGroupCode(t *testing.T) {
	mockStore := mocks.InitPropertyMockStore()

	propertyList := DefaultSystemPropertyList("", 3)

	mockReturnValue := []*SystemProperty{&propertyList[0], &propertyList[1], &propertyList[2]}

	mockStore.On("FindProperties").Return(mockReturnValue, nil).Once()

	req, err := http.NewRequest("GET", "/system-settings/group-code/group-code1", nil)
	ThrowErrorWhenNull(t, err)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.FindSystemPropertiesByCriteria)

	handler.ServeHTTP(recorder, req)

	Assert200(t, recorder)

	returnRef := []SystemProperty{propertyList[1]}
	err = json.NewDecoder(recorder.Body).Decode(&returnRef)

	ThrowErrorWhenNull(t, err)

	valueFromService := returnRef[0]

	AssertTestResultOutcome(t, valueFromService, propertyList[1])
}

func TestAddSystemPropertyHandler(t *testing.T) {

	mockStore := mocks.InitPropertyMockStore()

	systemPropertyToAdd := DefaultSystemProperty("")

	mockStore.On("AddProperty").Return(nil).Once()

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(systemPropertyToAdd)

	req, err := http.NewRequest("POST", "", buffer)
	ThrowErrorWhenNull(t, err)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.AddSystemProperties)

	handler.ServeHTTP(recorder, req)
}

func TestAddInvalidSystemPropertyHandler(t *testing.T) {

	mockStore := mocks.InitPropertyMockStore()

	systemPropertyToAdd := DefaultSystemProperty("")

	systemPropertyToAdd.GroupCode = "x"

	mockStore.On("AddProperty").Return(nil).Once()

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(systemPropertyToAdd)

	req, err := http.NewRequest("POST", "", buffer)
	ThrowErrorWhenNull(t, err)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.AddSystemProperties)

	handler.ServeHTTP(recorder, req)

	Assert422(t, recorder)
}
