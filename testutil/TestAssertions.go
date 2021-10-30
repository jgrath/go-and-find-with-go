package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func AssertTestResultOutcome(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("Test Failure! Actual: [%v] expected: [%v]", actual, expected)
	}
}

func ThrowErrorWhenNull(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func Assert200(t *testing.T, recorder *httptest.ResponseRecorder) {
	if recorder.Code != http.StatusOK {
		t.Errorf("Incorrect http response code. Got: [%v] expected: [%v]",
			recorder.Code, http.StatusOK)
	}
}

func Assert422(t *testing.T, recorder *httptest.ResponseRecorder) {
	if recorder.Code != http.StatusUnprocessableEntity {
		t.Errorf("Incorrect http response code. Got: [%v] expected: [%v]", recorder.Code, http.StatusOK)
	}
}

func Assert500(t *testing.T, recorder *httptest.ResponseRecorder) {
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect http response code. Got: [%v] expected: [%v]",
			recorder.Code, http.StatusOK)
	}
}
